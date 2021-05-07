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

// Infof ...
func (l *LoggerPrint) Infof(format string, a ...interface{}) {
	_, fname, line, _ := runtime.Caller(1)
	s := fmt.Sprintf(format+fmt.Sprintf(" (%s:%d)\n", fname, line), a...)
	if l.ID != "" {
		s = l.ID + " " + s
	}
	fmt.Fprintf(os.Stdout, s)
}

// Errorf ...
func (l *LoggerPrint) Errorf(format string, a ...interface{}) {
	_, fname, line, _ := runtime.Caller(1)
	s := fmt.Sprintf(format+fmt.Sprintf(" (%s:%d)\n", fname, line), a...)
	if l.ID != "" {
		s = l.ID + " " + s
	}
	fmt.Fprintf(os.Stderr, s)
}
