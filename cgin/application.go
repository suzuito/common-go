package cgin

import (
	"context"
	env "github.com/suzuito/common-env"
	"os"
	"strings"
)

// ApplicationGin ...
type ApplicationGin interface {
	AllowedOrigins() []string
	AllowedMethods() []string
	AllowedHeaders() []string
	AllowedCredential() bool
}

// ApplicationGinImpl ...
type ApplicationGinImpl struct {
	allowedOrigins    []string
	allowedMethods    []string
	allowedHeaders    []string
	allowedCredential bool
}

// AllowedOrigins ...
func (a *ApplicationGinImpl) AllowedOrigins() []string {
	return a.allowedOrigins
}

// AllowedMethods ...
func (a *ApplicationGinImpl) AllowedMethods() []string {
	return a.allowedMethods
}

// AllowedHeaders ...
func (a *ApplicationGinImpl) AllowedHeaders() []string {
	return a.allowedHeaders
}

// AllowedCredential ...
func (a *ApplicationGinImpl) AllowedCredential() bool {
	return a.allowedCredential
}

// NewApplicationGinImpl ...
func NewApplicationGinImpl(ctx context.Context) (*ApplicationGinImpl, error) {
	return &ApplicationGinImpl{
		allowedOrigins:    strings.Split(os.Getenv("ALLOWED_ORIGINS"), ","),
		allowedMethods:    strings.Split(os.Getenv("ALLOWED_METHODS"), ","),
		allowedHeaders:    strings.Split(os.Getenv("ALLOWED_HEADERS"), ","),
		allowedCredential: env.GetenvAsBool("ALLOWED_CREDENTIAL", false),
	}, nil
}
