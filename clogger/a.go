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
	fmt.Fprintf(os.Stdout, format+fmt.Sprintf("%s (%s:%d)\n", l.ID, fname, line), a...)
}

// Errorf ...
func (l *LoggerPrint) Errorf(format string, a ...interface{}) {
	_, fname, line, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stderr, format+fmt.Sprintf("%s (%s:%d)\n", l.ID, fname, line), a...)
}
