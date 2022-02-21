package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type simpleLogger struct {
	f io.Writer
}

var defaultLogger = simpleLogger{os.Stderr}

func (sl simpleLogger) Debugf(msg string, args ...interface{}) {
	fmt.Fprintf(sl.f, "debug: "+msg, args...)
	if !strings.HasSuffix(msg, "\n") {
		fmt.Fprintln(sl.f)
	}
}

func (sl simpleLogger) Warnf(msg string, args ...interface{}) {
	fmt.Fprintf(sl.f, "warn: "+msg, args...)
	if !strings.HasSuffix(msg, "\n") {
		fmt.Fprintln(sl.f)
	}
}
