package main

import (
	"flag"
	"os"

	genericapiserver "k8s.io/apiserver/pkg/server"

	"github.com/marcos30004347/kubeapi/pkg/cmd/server"
	"k8s.io/component-base/logs"
	"k8s.io/klog"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	stopCh := genericapiserver.SetupSignalHandler()
	options := server.NewCustomServerOptions(os.Stdout, os.Stderr)
	cmd := server.NewCommandStartCustomServer(options, stopCh)
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	if err := cmd.Execute(); err != nil {
		klog.Fatal(err)
	}
}
