package esbuild

import (
	"context"
	"github.com/lmika/rwt/internal/providers/termout"
	"github.com/pkg/errors"
	"os"
	"os/exec"
)

type ESBuild struct {
}

func New() *ESBuild {
	return &ESBuild{}
}

func (eb *ESBuild) Build(ctx context.Context, source, target string) error {
	return eb.run(ctx, source, "--bundle", "--outfile="+target)
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
