package npmreader_test

import (
	"context"
	"strings"
	"testing"

	"github.com/lmika/rwt/internal/models/projects"
	"github.com/lmika/rwt/internal/providers/npmreader"
	"github.com/stretchr/testify/assert"
)

func TestReadFromPackageJson(t *testing.T) {
	t.Run("read project from package.json", func(t *testing.T) {
		proj, err := npmreader.ReadFromPackageJson(context.Background(), strings.NewReader(samplePackage))
		assert.NoError(t, err)

		assert.Contains(t, proj.Targets, projects.Target{
			Type:   projects.JSTargetType,
			Source: "assets/js/main.js",
			Target: "build/assets/js/main.js",
		})
		assert.Contains(t, proj.Targets, projects.Target{
			Type:   projects.CSSTargetType,
			Source: "assets/css/main.css",
			Target: "build/assets/css/main.css",
		})
		assert.Contains(t, proj.Loaders, projects.Loader{
			Type:    projects.FileLoader,
			Pattern: ".csv",
		})
		assert.Contains(t, proj.Loaders, projects.Loader{
			Type:    projects.FileLoader,
			Pattern: ".tsv",
		})
	})
}

var samplePackage = `
{
	"name": "broadtail",
	"version": "1.0.0",
	"description": "",
	"main": "index.js",
	"scripts": {
	  "build-js": "esbuild assets/js/main.js --bundle --outfile=build/assets/js/main.js",
	  "build-css": "esbuild assets/css/main.css --bundle --outfile=build/assets/css/main.css",
	  "watch-js": "esbuild assets/js/main.js --bundle --outfile=build/assets/js/main.js --watch",
	  "watch-css": "esbuild assets/css/main.css --bundle --loader:.eot=file --loader:.svg=file --loader:.woff=file --loader:.woff2=file --loader:.ttf=file --outfile=build/assets/css/main.css --watch"
	},
	"author": "",
	"license": "ISC",
	"dependencies": {
	  "esbuild": "^0.13.8",
	  "font-awesome": "^4.7.0",
	  "jquery": "^3.6.0",
	  "jquery-ujs": "^1.2.3"
	},
	"rwt:project": {
	  "targets": {
	    "build/assets/js/main.js": "assets/js/main.js",
	    "build/assets/css/main.css": "assets/css/main.css"
	  },
	  "loaders": {
		"file": [".csv", ".tsv"]
	  }
	}
  }
`
