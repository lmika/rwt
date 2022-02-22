package main

import (
	"github.com/lmika/rwt/cmds"
	"github.com/spf13/cobra"
)

var rootCmt = &cobra.Command{
	Use:   "rwt",
	Short: "Redundant Web Toolkit",
	Long:  `A tool for dealing with web assets which really should not exist.`,
}

func main() {
	rootCmt.AddCommand(cmds.Build())
	rootCmt.AddCommand(cmds.Watch())
	rootCmt.Execute()
}
