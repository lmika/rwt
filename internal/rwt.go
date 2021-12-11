package internal

import (
	"context"
	"github.com/lmika/rwt/internal/models"
	"github.com/lmika/rwt/internal/providers/cfgreader"
	"github.com/lmika/rwt/internal/providers/esbuild"
	"github.com/lmika/rwt/internal/providers/termout"
	"github.com/lmika/rwt/internal/services/builder"
)

type RWT struct {
	config  *models.Config
	builder *builder.Service
}

func New() (*RWT, error) {
	eb := esbuild.New()
	bld := builder.New(eb)

	cfg, err := cfgreader.ReadConfig("rwt.json")
	if err != nil {
		return nil, err
	}

	return &RWT{
		config:  cfg,
		builder: bld,
	}, nil
}

func (r *RWT) Build(ctx context.Context) error {
	ctx = termout.WithCtx(ctx, termout.New())
	return r.builder.Build(ctx, r.config)
}
