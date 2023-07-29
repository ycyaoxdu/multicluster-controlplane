package agent

import (
	"context"

	"github.com/spf13/pflag"
)

type Interface interface {
	AddFlags(fs *pflag.FlagSet)
	WithClusterName(clusterName string) *AgentOptions
	WithKubeconfig(kubeConfig string) *AgentOptions
	WithSpokeKubeconfig(spokeKubeConfig string) *AgentOptions
	WithBootstrapKubeconfig(bootstrapKubeconfig string) *AgentOptions
	WithHubKubeconfigDir(hubKubeconfigDir string) *AgentOptions
	WithHubKubeconfigSecreName(hubKubeconfigSecreName string) *AgentOptions
	RunAgent(ctx context.Context) error
}
