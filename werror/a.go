package werror

import (
	"fmt"
	"runtime"
)

type werror struct {
	appendedMsg string
	err         error
	file        string
	line        int
}

// New ...
func New(err error) error {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "?"
		line = 0
	}
	return &werror{
		appendedMsg: "",
		err:         err,
		file:        file,
		line:        line,
	}
}

// Newf ...
func Newf(err error, format string, a ...interface{}) error {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "?"
		line = 0
	}
	return &werror{
		appendedMsg: fmt.Sprintf(format, a...),
		err:         err,
		file:        file,
		line:        line,
	}
}

func (e *werror) Unwrap() error {
	return e.err
}

func (e *werror) Error() string {
	msg := ""
	if e.err != nil {
		msg = e.err.Error()
	}
	return fmt.Sprintf("%s, %s, at %s:%d", msg, e.appendedMsg, e.file, e.line)
}