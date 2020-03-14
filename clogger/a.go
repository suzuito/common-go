package clogger

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

// Logger ...
type Logger interface {
	Infof(format string, a ...interface{})
	Errorf(format string, a ...interface{})
	Request(req *http.Request)
	Close()
}

// LoggerPrint ...
type LoggerPrint struct {
}

// Infof ...
func (l *LoggerPrint) Infof(format string, a ...interface{}) {
	_, fname, line, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stdout, format+fmt.Sprintf(" (%s:%d)\n", fname, line), a...)
}

// Errorf ...
func (l *LoggerPrint) Errorf(format string, a ...interface{}) {
	_, fname, line, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stderr, format+fmt.Sprintf(" (%s:%d)\n", fname, line), a...)
}

// Request ...
func (l *LoggerPrint) Request(req *http.Request) {}

// Close ...
func (l *LoggerPrint) Close() {
}
