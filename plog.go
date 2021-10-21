package log

import (
	"fmt"
	"io"
	"sync"
)

type Logger struct {
	Level Level
	mu    sync.Mutex
	w     io.Writer
}

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func New(level Level) *Logger {
	return &Logger{
		Level: level,
	}
}

func Default() *Logger {
	return &Logger{Level: INFO}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	format = "[Debug]" + format
	fmt.Fprintf(l.w, format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	format = "[Info]" + format
	fmt.Fprintf(l.w, format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	format = "[Warn]" + format
	fmt.Fprintf(l.w, format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	format = "[Error]" + format
	fmt.Fprintf(l.w, format, v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	format = "[Fatal]" + format
	fmt.Fprintf(l.w, format, v...)
}

func (l *Logger) Debug(msg string) {
	msg = "[Debug]" + msg
	fmt.Fprint(l.w, msg)
}

func (l *Logger) Info(msg string) {
	msg = "[Info]" + msg
	fmt.Fprint(l.w, msg)
}

func (l *Logger) Warn(msg string) {
	msg = "[Warn]" + msg
	fmt.Fprint(l.w, msg)
}

func (l *Logger) Error(msg string) {
	msg = "[Error]" + msg
	fmt.Fprint(l.w, msg)
}

func (l *Logger) Fatal(msg string) {
	msg = "[Fatal]" + msg
	fmt.Fprint(l.w, msg)
}