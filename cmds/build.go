package cmds

import (
	"context"
	"github.com/lmika/gopkgs/cli"
	"github.com/lmika/rwt/internal"
	"github.com/spf13/cobra"
)

func Build() *cobra.Command {
	command := &cobra.Command{
		Use:   "build",
		Short: "Build the web assets",
		Long:  `A tool for dealing with web assets which really should not exist.`,
		Run: func(cmd *cobra.Command, args []string) {
			rwt, err := internal.New()
			if err != nil {
				cli.Fatal(err)
			}

			if err := rwt.Build(context.Background()); err != nil {
				cli.Fatal(err)
			}
		},
	}

	return command
}
