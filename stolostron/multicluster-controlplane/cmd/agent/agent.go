// Copyright Contributors to the Open Cluster Management project

package main

import (
	"context"
	"os"
	"path"

	"github.com/spf13/cobra"

	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/component-base/cli"
	"k8s.io/klog/v2"

	ocmagent "open-cluster-management.io/multicluster-controlplane/pkg/cmd/agent"
	"open-cluster-management.io/multicluster-controlplane/pkg/features"

	"open-cluster-management.io/multicluster-controlplane/stolostron/multicluster-controlplane/pkg/agent"
	"open-cluster-management.io/multicluster-controlplane/stolostron/multicluster-controlplane/pkg/feature"
)

func init() {
	utilruntime.Must(features.DefaultAgentMutableFeatureGate.Add(feature.DefaultControlPlaneFeatureGates))
}

func main() {
	agentOptions := agent.NewAgentOptions()
	cmd := &cobra.Command{
		Use:   "multicluster-agent",
		Short: "Start a multicluster agent",
		RunE: ocmagent.RunAgentHandler(func(ctx context.Context) error {
			// starting agent firstly to request the hub kubeconfig
			go func() {
				klog.Info("starting the controlplane agent")
				if err := agentOptions.RunAgent(ctx); err != nil {
					klog.Fatalf("failed to run agent, %v", err)
				}
			}()

			// wait for the agent is registered
			hubKubeConfig := path.Join(agentOptions.RegistrationAgent.HubKubeconfigDir, "kubeconfig")
			if err := agentOptions.WaitForValidHubKubeConfig(ctx, hubKubeConfig); err != nil {
				return err
			}

			if err := agentOptions.RunAddOns(ctx); err != nil {
				return err
			}
			return nil
		}),
	}

	flags := cmd.Flags()

	flags.UintVar(
		&agentOptions.Frequency,
		"update-frequency",
		60,
		"The status update frequency (in seconds) of a mutation policy",
	)

	flags.Uint8Var(
		&agentOptions.DecryptionConcurrency,
		"decryption-concurrency",
		5,
		"The max number of concurrent policy template decryptions",
	)

	flags.Uint8Var(
		&agentOptions.EvaluationConcurrency,
		"evaluation-concurrency",
		// Set a low default to not add too much load to the Kubernetes API server in resource constrained deployments.
		2,
		"The max number of concurrent configuration policy evaluations",
	)

	flags.BoolVar(
		&agentOptions.EnableMetrics,
		"enable-metrics",
		false,
		"Disable custom metrics collection",
	)

	agentOptions.AddFlags(flags)

	os.Exit(cli.Run(cmd))
}
