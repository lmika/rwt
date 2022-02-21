package builderservice

import (
	"context"

	"github.com/lmika/rwt/projects"
	"github.com/lmika/rwt/utils/logger"
	"github.com/pkg/errors"
)

type BuilderService struct {
	targetBuilder TargetBuilder
}

func New(targetBuilder TargetBuilder) *BuilderService {
	return &BuilderService{}
}

func (bs *BuilderService) Build(ctx context.Context, proj *projects.Project) error {
	log := logger.FromContext(ctx)

	for _, target := range proj.Targets {
		log.Debugf("building '%v' from '%v' (type %v)", target.Target, target.Source, target.Type)
		if err := bs.targetBuilder.BuildTarget(ctx, target); err != nil {
			return errors.Wrapf(err, "unable to build target '%v' (type %v) from '%v'", target.Target, target.Type, target.Source)
		}
	}
	return nil
}

type TargetBuilder interface {
	BuildTarget(ctx context.Context, target projects.Target) error
}
