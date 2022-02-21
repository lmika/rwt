package esbuild

import (
	"context"
	"os"
	"os/exec"

	"github.com/lmika/rwt/projects"
	"github.com/pkg/errors"
)

type ESBuild struct{}

func New() *ESBuild {
	return &ESBuild{}
}

func (eb *ESBuild) BuildTarget(ctx context.Context, target projects.Target) error {
	return eb.execCommand(ctx, target.Source, "--bundle", "--outfile="+target.Target)
}

func (eb *ESBuild) execCommand(ctx context.Context, args ...string) error {
	e := exec.CommandContext(ctx, "node_modules/.bin/esbuild", args...)
	e.Stdout = os.Stdout
	e.Stderr = os.Stderr
	if err := e.Run(); err != nil {
		return errors.Wrapf(err, "could not execute esbuild: %v", args)
	}

	return nil
}
