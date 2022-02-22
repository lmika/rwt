# rwt - Redundant Web Tool

The frontend build tool that really shouldn't exist, but does.  More or less a glorified wrapper around [esbuild](https://esbuild.github.io).

## Usage

To make use of `rwt`, you will need to do the following:

1. Create a package.json file and install `esbuild` as a dependency.
2. Define a project (see below).
3. Run `rwt build` to build the project, or `rwt watch` to start watching for changes.

## Project

A project is defined in the "package.json" by adding the `rwt:project` node.
An example is provided below.  Note that this is not standard in any way, and is subject to change.

```
{
    "name": "frontend-project",
    "version": "1.0.0",
    "dependencies": {
        "esbuild": "^0.13.8"
    },
    "rwt:project": {
        "targets": {
            "build/main.js": "src/main.js",
            "build/main.css": "src/main.css",
        },
        "loaders": {
            "file": [".eot", ".svg", ".woff"]
        }
    }
}
```

The structure is as follows:

### Targets

Targets define what the artefacts of the build should be, and the source files from which those artefacts should
be constructed from.  Artefacts are set as the key of the object, and the top-level source file is the value.

From the example above, the build is expected to produce a bundled JavaScript file at `build/main.js`,
by running `esbuild` over `src/main.js`.  Likewise, it is expected to produce a bundled CSS file at `build/main.css` using
the source file `src/main.css`.  This is used for both `build` and `watch`.

### Loaders

Loaders configure which loaders should be used for files with a particular extension.  The loaders are defined as
keys of the object, with the value being the list of extensions the loader should be used for.  Currently, only
the `file` loader is supported, which saves the required file as an external file.

Loaders are not required, and include a default set of loaders for JavaScript and CSS files.