package cmds

import (
	"context"
	"os"
	"os/signal"

	"github.com/lmika/gopkgs/cli"
	"github.com/lmika/rwt/internal"
	"github.com/spf13/cobra"
)

func Watch() *cobra.Command {
	command := &cobra.Command{
		Use:   "build",
		Short: "Start watching for changes",
		Long:  `Start watching for changes in the web assets and rebuild them when they change.`,
		Run: func(cmd *cobra.Command, args []string) {
			rwt, err := internal.New()
			if err != nil {
				cli.Fatal(err)
			}

			ctx, cancelFn := context.WithCancel(context.Background())
			defer cancelFn()

			go func() {
				c := make(chan os.Signal, 1)
				signal.Notify(c, os.Interrupt)

				// Block until a signal is received.
				<-c
				cancelFn()
			}()

			rwt.Watch(ctx)
		},
	}

	return command
}
