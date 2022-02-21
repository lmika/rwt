package builder

import (
	"context"

	"github.com/lmika/rwt/internal/models/projects"
	"github.com/lmika/rwt/internal/providers/termout"
)

type Service struct {
	esbuildable ESBuildable
}

func New(esbuildable ESBuildable) *Service {
	return &Service{
		esbuildable: esbuildable,
	}
}

func (s *Service) Build(ctx context.Context, cfg *projects.Project) error {
	for _, target := range cfg.Targets {
		termout.FromCtx(ctx).Verbosef("building '%v' from '%v' (type %v)", target.Target, target.Source, target.Type)
		if err := s.esbuildable.BuildTarget(ctx, target); err != nil {
			return err
		}
	}
	return nil
}

type ESBuildable interface {
	BuildTarget(ctx context.Context, target projects.Target) error
}
