package projects

type TargetType string

type LoaderType string

const (
	JSTargetType  TargetType = "js"
	CSSTargetType            = "css"
)

const (
	FileLoader LoaderType = "file"
)

type Project struct {
	Targets []Target
	Loaders []Loader
}

type Target struct {
	Type   TargetType
	Source string
	Target string
}

type Loader struct {
	Type    LoaderType
	Pattern string
}
