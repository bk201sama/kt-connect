package main

import (
	"os"

	"github.com/alibaba/kt-connect/pkg/kt/cmd"
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func main() {
	flags := pflag.NewFlagSet("kubectl-connect", pflag.ExitOnError)
	pflag.CommandLine = flags

	root := cmd.NewConnectCommand(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}