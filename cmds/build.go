package cmds

import (
	"github.com/lmika/gopkgs/cli"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func Build() *cobra.Command {
	command := &cobra.Command{
		Use: "build",
		Short: "Build the web assets",
		Long: `A tool for dealing with web assets which really should not exist.`,
		Run: func(cmd *cobra.Command, args []string) {
			e := exec.Command("node_modules/.bin/esbuild", "assets/css/main.css", "--bundle", "--outfile=build/assets/css/main.css")
			e.Stdout = os.Stdout
			e.Stderr = os.Stderr
			if err := e.Run(); err != nil {
				cli.Fatalf("could not execute esbuild: %v", err)
			}
		},
	}

	return command
}
