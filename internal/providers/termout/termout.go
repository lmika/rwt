package termout

import (
	"context"
	"fmt"
)

type TermOut struct {
}

func New() *TermOut {
	return &TermOut{}
}

func (t *TermOut) Verbosef(pattern string, args ...interface{}) {
	fmt.Printf(pattern, args...)
	fmt.Println()
}

func FromCtx(ctx context.Context) *TermOut {
	t, _ := ctx.Value(termOutKey).(*TermOut)
	return t
}

func WithCtx(ctx context.Context, out *TermOut) context.Context {
	return context.WithValue(ctx, termOutKey, out)
}

type termOutKeyType struct{}

var termOutKey = termOutKeyType{}
