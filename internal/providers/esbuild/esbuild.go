package esbuild

import (
	"context"
	"os"
	"os/exec"

	"github.com/lmika/rwt/internal/models/projects"
	"github.com/lmika/rwt/internal/providers/termout"
	"github.com/pkg/errors"
)

type ESBuild struct {
}

func New() *ESBuild {
	return &ESBuild{}
}

func (eb *ESBuild) BuildTarget(ctx context.Context, target projects.Target) error {
	return eb.run(ctx, target.Source, "--bundle", "--outfile="+target.Target)
}

func (eb *ESBuild) WatchTarget(ctx context.Context, target projects.Target) error {
	return eb.run(ctx, "--watch", target.Source, "--bundle", "--outfile="+target.Target)
}

func (eb *ESBuild) run(ctx context.Context, args ...string) error {
	termout.FromCtx(ctx).Verbosef("esbuild %v", args)

	e := exec.Command("node_modules/.bin/esbuild", args...)
	e.Stdout = os.Stdout
	e.Stderr = os.Stderr
	if err := e.Run(); err != nil {
		return errors.Wrap(err, "could not execute esbuild")
	}
	return nil
}
