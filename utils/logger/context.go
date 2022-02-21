package logger

import (
	"context"
)

func FromContext(ctx context.Context) Logger {
	l, ok := ctx.Value(loggerContextKey).(Logger)
	if ok {
		return l
	}

	return defaultLogger
}

type loggerContextKeyType struct{}

var loggerContextKey = loggerContextKeyType{}
