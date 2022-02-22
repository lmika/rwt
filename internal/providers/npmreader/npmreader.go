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

	for targetFile, srcFile := range pkg.Project.Targets {
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

	for loader, patterns := range pkg.Project.Loaders {
		for _, pattern := range patterns {
			var loaderType projects.LoaderType
			switch loader {
			case "file":
				loaderType = projects.FileLoader
			case "text":
				loaderType = projects.TextLoader
			default:
				termout.FromCtx(ctx).Warnf("unrecognised loader type '%v', ignoring", loader)
				continue
			}

			proj.Loaders = append(proj.Loaders, projects.Loader{
				Type:    loaderType,
				Pattern: pattern,
			})
		}
	}

	return proj, nil
}

type packageJson struct {
	Project packageProject `json:"rwt:project"`
}

type packageProject struct {
	Targets map[string]string   `json:"targets"`
	Loaders map[string][]string `json:"loaders"`
}
