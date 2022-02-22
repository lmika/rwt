package esbuild

import (
	"context"
	"fmt"
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

func (eb *ESBuild) BuildTarget(ctx context.Context, project *projects.Project, target projects.Target) error {
	args := append(eb.projectBuildArgs(project), target.Source, "--bundle", "--outfile="+target.Target)
	return eb.run(ctx, args)
}

func (eb *ESBuild) WatchTarget(ctx context.Context, project *projects.Project, target projects.Target) error {
	args := append(eb.projectBuildArgs(project), "--watch", target.Source, "--bundle", "--outfile="+target.Target)

	termout.FromCtx(ctx).Verbosef("esbuild %v", args)

	e := exec.Command("node_modules/.bin/esbuild", args...)
	e.Stdout = os.Stdout
	e.Stderr = os.Stderr

	w, err := e.StdinPipe()
	if err != nil {
		return errors.Wrap(err, "cound not setup stdin pipe")
	}

	go func() {
		<-ctx.Done()
		w.Close()
	}()

	if err := e.Run(); err != nil {
		return errors.Wrap(err, "could not execute esbuild")
	}
	return nil
}

func (eb *ESBuild) run(ctx context.Context, args []string) error {
	termout.FromCtx(ctx).Verbosef("esbuild %v", args)

	e := exec.Command("node_modules/.bin/esbuild", args...)
	e.Stdout = os.Stdout
	e.Stderr = os.Stderr
	if err := e.Run(); err != nil {
		return errors.Wrap(err, "could not execute esbuild")
	}
	return nil
}

func (eb *ESBuild) projectBuildArgs(project *projects.Project) []string {
	projectArgs := make([]string, 0)

	// Add any loaders
	for _, loader := range project.Loaders {
		projectArgs = append(projectArgs, fmt.Sprintf("--loader:%v=%v", loader.Pattern, loader.Type))
	}

	return projectArgs
}
