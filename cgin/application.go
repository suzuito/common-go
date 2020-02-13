package cgin

import (
	"context"
	"os"
	"strings"
)

// ApplicationGin ...
type ApplicationGin interface {
	AllowedOrigins() []string
}

// ApplicationGinImpl ...
type ApplicationGinImpl struct {
	allowedOrigins []string
}

// NewApplicationGinImpl ...
func NewApplicationGinImpl(ctx context.Context) (*ApplicationGinImpl, error) {
	return &ApplicationGinImpl{
		allowedOrigins: strings.Split(os.Getenv("ALLOWED_ORIGINS"), ","),
	}, nil
}
