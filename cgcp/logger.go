package cgcp

import (
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/logging"
)

// LoggerGCP ...
type LoggerGCP struct {
	cli    *logging.Client
	logger *logging.Logger
}

func (l *LoggerGCP) log(serv logging.Severity, format string, a ...interface{}) {
	e := logging.Entry{
		Payload:  fmt.Sprintf(format, a...),
		Severity: serv,
	}
	l.logger.Log(e)
}

// Infof ...
func (l *LoggerGCP) Infof(format string, a ...interface{}) {
	l.log(logging.Info, format, a...)
}

// Errorf ...
func (l *LoggerGCP) Errorf(format string, a ...interface{}) {
	l.log(logging.Error, format, a...)
}

// Request ...
func (l *LoggerGCP) Request(req *http.Request) {
	e := logging.Entry{
		HTTPRequest: &logging.HTTPRequest{
			Request: req,
		},
	}
	if req.Response != nil {
		e.HTTPRequest.Status = req.Response.StatusCode
	}
	logger := l.cli.Logger("api")
	logger.Log(e)
}

// Close ...
func (l *LoggerGCP) Close() {
	l.cli.Close()
}

// NewLoggerGCP ...
func NewLoggerGCP(ctx context.Context, projectID string) (*LoggerGCP, error) {
	cli, err := logging.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return &LoggerGCP{
		cli:    cli,
		logger: cli.Logger("api"),
	}, nil
}

// NewLoggerGCP2 ...
func NewLoggerGCP2(ctx context.Context, projectID string, req *http.Request) (*LoggerGCP, error) {
	cli, err := logging.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	child := cli.Logger("api-log")
	return &LoggerGCP{
		cli:    cli,
		logger: child,
	}, nil
}
