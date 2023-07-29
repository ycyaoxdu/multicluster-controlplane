// Copyright Contributors to the Open Cluster Management project

package main

import (
	"os"

	"github.com/spf13/cobra"

	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/component-base/cli"
	logsapi "k8s.io/component-base/logs/api/v1"

	cmdcontroller "open-cluster-management.io/multicluster-controlplane/pkg/cmd/controller"
	controllers "open-cluster-management.io/multicluster-controlplane/pkg/controllers"
	"open-cluster-management.io/multicluster-controlplane/pkg/features"
	"open-cluster-management.io/multicluster-controlplane/pkg/servers/options"

	controller "open-cluster-management.io/multicluster-controlplane/stolostron/multicluster-controlplane/pkg/controllers"
	"open-cluster-management.io/multicluster-controlplane/stolostron/multicluster-controlplane/pkg/controllers/selfmanagement"
	"open-cluster-management.io/multicluster-controlplane/stolostron/multicluster-controlplane/pkg/feature"
)

func init() {
	// register log to featuregate
	utilruntime.Must(logsapi.AddFeatureGates(utilfeature.DefaultMutableFeatureGate))
	// init feature gates
	utilruntime.Must(features.DefaultControlplaneMutableFeatureGate.Add(feature.DefaultControlPlaneFeatureGates))
}

func main() {
	options := options.NewServerRunOptions()
	cmd := &cobra.Command{
		Use:   "multicluster-controlplane",
		Short: "Start a multicluster controlplane",

		RunE: cmdcontroller.ServerRunHandler(options,
			controllers.ControllerInstaller{
				Name:       "next-gen-controlplane-controllers",
				Controller: controller.InstallControllers,
			},
			controllers.ControllerInstaller{
				Name:       "next-gen-controlplane-self-management",
				Controller: selfmanagement.InstallControllers(options),
			}),
	}

	options.AddFlags(cmd.Flags())

	os.Exit(cli.Run(cmd))
}
