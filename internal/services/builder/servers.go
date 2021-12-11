package builder

import (
	"context"
	"github.com/lmika/rwt/internal/models"
)

type Service struct {
	esbuildable ESBuildable
}

func New(esbuildable ESBuildable) *Service {
	return &Service{
		esbuildable: esbuildable,
	}
}

func (s *Service) Build(ctx context.Context, cfg *models.Config) error {
	for target, source := range cfg.Targets {
		if err := s.esbuildable.Build(ctx, source, target); err != nil {
			return err
		}
	}
	return nil
}

type ESBuildable interface {
	Build(ctx context.Context, source, target string) error
}
