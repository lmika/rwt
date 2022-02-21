package projects

import (
	"context"
	"encoding/json"
	"io"
	"strings"

	"github.com/lmika/rwt/utils/logger"
	"github.com/pkg/errors"
)

func ReadFromPackageJson(ctx context.Context, r io.Reader) (*Project, error) {
	var pkg packageJson
	if err := json.NewDecoder(r).Decode(&pkg); err != nil {
		return nil, errors.Wrap(err, "cannot unmarshal package.json")
	}

	return buildProjectFromPackage(ctx, pkg)
}

func buildProjectFromPackage(ctx context.Context, pkg packageJson) (*Project, error) {
	proj := new(Project)

	for targetFile, srcFile := range pkg.Targets {

		var targetType TargetType
		if strings.HasSuffix(targetFile, ".js") {
			targetType = JSTargetType
		} else if strings.HasSuffix(targetFile, ".css") {
			targetType = CSSTargetType
		} else {
			logger.FromContext(ctx).Warnf("unrecognised target type '%v', ignoring", targetFile)
			continue
		}

		proj.Targets = append(proj.Targets, Target{
			Source: srcFile,
			Target: targetFile,
			Type:   targetType,
		})
	}

	return proj, nil
}

type packageJson struct {
	Targets map[string]string `json:"rwt:targets"`
}
