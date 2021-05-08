package clogger

import (
	"fmt"
	"os"
	"runtime"
)

// Logger ...
type Logger interface {
	Infof(format string, a ...interface{})
	Errorf(format string, a ...interface{})
}

// LoggerPrint ...
type LoggerPrint struct {
	ID string
}

func (l *LoggerPrint) msg(level string, format string, a ...interface{}) string {
	_, fname, line, _ := runtime.Caller(2)
	s := fmt.Sprintf(format+fmt.Sprintf(" (%s:%d)\n", fname, line), a...)
	if l.ID != "" {
		s = l.ID + " " + s
	}
	s = level + " " + s
	return s
}

// Infof ...
func (l *LoggerPrint) Infof(format string, a ...interface{}) {
	fmt.Fprintf(os.Stdout, l.msg("info", format, a...))
}

// Errorf ...
func (l *LoggerPrint) Errorf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, l.msg("info", format, a...))
}
