package cmds

import (
	"bytes"
	"context"
	"os"

	"github.com/lmika/gopkgs/cli"
	"github.com/lmika/rwt/projects"
	"github.com/lmika/rwt/providers/esbuild"
	"github.com/lmika/rwt/services/builderservice"
	"github.com/spf13/cobra"
)

func Build() *cobra.Command {
	command := &cobra.Command{
		Use:   "build",
		Short: "Build the web assets",
		Long:  `A tool for dealing with web assets which really should not exist.`,
		Run: func(cmd *cobra.Command, args []string) {
			packageJson, err := os.ReadFile("package.json")
			if err != nil {
				cli.Fatalf("could not read package.json: %v", err)
			}

			ctx := context.Background()
			project, err := projects.ReadFromPackageJson(ctx, bytes.NewReader(packageJson))
			if err != nil {
				cli.Fatalf("could not read project from package.json: %v", err)
			}

			builderService := builderservice.New(esbuild.New())
			if err := builderService.Build(ctx, project); err != nil {
				cli.Fatalf("could not build project: %v", err)
			}
		},
	}

	return command
}
