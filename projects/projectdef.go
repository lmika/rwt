package projects

type TargetType string

const (
	JSTargetType TargetType = "js"
	CSSTargetType = "css"
)

type Project struct {
	Targets []Target
}

type Target struct {
	Type TargetType
	Source string
	Target string
}