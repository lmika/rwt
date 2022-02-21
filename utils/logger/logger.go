package logger

type Logger interface {
	Debugf(msg string, args ...interface{})
	Warnf(msg string, args ...interface{})
}
