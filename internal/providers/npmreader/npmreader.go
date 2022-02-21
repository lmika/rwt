package npmreader

import (
	"context"
	"encoding/json"
	"io"
	"strings"

	"github.com/lmika/rwt/internal/models/projects"
	"github.com/lmika/rwt/internal/providers/termout"
	"github.com/pkg/errors"
)

func ReadFromPackageJson(ctx context.Context, r io.Reader) (*projects.Project, error) {
	var pkg packageJson
	if err := json.NewDecoder(r).Decode(&pkg); err != nil {
		return nil, errors.Wrap(err, "cannot unmarshal package.json")
	}

	return buildProjectFromPackage(ctx, pkg)
}

func buildProjectFromPackage(ctx context.Context, pkg packageJson) (*projects.Project, error) {
	proj := new(projects.Project)

	for targetFile, srcFile := range pkg.Targets {

		var targetType projects.TargetType
		if strings.HasSuffix(targetFile, ".js") {
			targetType = projects.JSTargetType
		} else if strings.HasSuffix(targetFile, ".css") {
			targetType = projects.CSSTargetType
		} else {
			termout.FromCtx(ctx).Warnf("unrecognised target type '%v', ignoring", targetFile)
			continue
		}

		proj.Targets = append(proj.Targets, projects.Target{
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
