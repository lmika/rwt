package builder

import (
	"context"
	"sync"

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

func (s *Service) Watch(ctx context.Context, cfg *projects.Project) {
	waitGroup := new(sync.WaitGroup)

	for _, target := range cfg.Targets {
		termout.FromCtx(ctx).Verbosef("watching '%v' from '%v' (type %v)", target.Target, target.Source, target.Type)

		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()

			if err := s.esbuildable.WatchTarget(ctx, target); err != nil {
				termout.FromCtx(ctx).Verbosef("error: %v")
			}
		}()
	}

	<-ctx.Done()
	waitGroup.Wait()
}

type ESBuildable interface {
	BuildTarget(ctx context.Context, target projects.Target) error
	WatchTarget(ctx context.Context, target projects.Target) error
}
