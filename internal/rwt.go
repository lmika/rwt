package internal

import (
	"bytes"
	"context"
	"os"

	"github.com/lmika/rwt/internal/models/projects"
	"github.com/lmika/rwt/internal/providers/esbuild"
	"github.com/lmika/rwt/internal/providers/npmreader"
	"github.com/lmika/rwt/internal/providers/termout"
	"github.com/lmika/rwt/internal/services/builder"
	"github.com/pkg/errors"
)

type RWT struct {
	project *projects.Project
	builder *builder.Service
}

func New() (*RWT, error) {
	eb := esbuild.New()
	bld := builder.New(eb)

	packageJson, err := os.ReadFile("package.json")
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read 'package.json'")
	}

	project, err := npmreader.ReadFromPackageJson(context.Background(), bytes.NewReader(packageJson))
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read project from package.json")
	}

	return &RWT{
		project: project,
		builder: bld,
	}, nil
}

func (r *RWT) Build(ctx context.Context) error {
	ctx = termout.WithCtx(ctx, termout.New())
	return r.builder.Build(ctx, r.project)
}
