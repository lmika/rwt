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
		if err := s.esbuildable.Run(ctx, target, "--bundle", "--outfile="+source); err != nil {
			return err
		}
	}
	return nil
}

type ESBuildable interface {
	Run(ctx context.Context, args ...string) error
}
