package cfgreader

import (
	"encoding/json"
	"github.com/lmika/rwt/internal/models"
	"github.com/pkg/errors"
	"os"
)

func ReadConfig(filename string) (*models.Config, error) {
	bts, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read config %v", filename)
	}

	m := models.Config{}
	if err := json.Unmarshal(bts, &m); err != nil {
		return nil, errors.Wrapf(err, "cannot read config %v", filename)
	}

	return &m, nil
}
