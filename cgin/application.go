package cgin

import (
	"context"
	"os"
	"strings"

	env "github.com/suzuito/common-env"
)

// ApplicationGin ...
// Deprecated:
type ApplicationGin interface {
	AllowedOrigins() []string
	AllowedMethods() []string
	AllowedHeaders() []string
	ExposeHeaders() []string
	AllowedCredential() bool
}

// ApplicationGinImpl ...
// Deprecated:
type ApplicationGinImpl struct {
	allowedOrigins    []string
	allowedMethods    []string
	allowedHeaders    []string
	exposeHeaders     []string
	allowedCredential bool
}

// AllowedOrigins ...
// Deprecated:
func (a *ApplicationGinImpl) AllowedOrigins() []string {
	return a.allowedOrigins
}

// AllowedMethods ...
// Deprecated:
func (a *ApplicationGinImpl) AllowedMethods() []string {
	return a.allowedMethods
}

// AllowedHeaders ...
// Deprecated:
func (a *ApplicationGinImpl) AllowedHeaders() []string {
	return a.allowedHeaders
}

// ExposeHeaders ...
// Deprecated:
func (a *ApplicationGinImpl) ExposeHeaders() []string {
	return a.exposeHeaders
}

// AllowedCredential ...
// Deprecated:
func (a *ApplicationGinImpl) AllowedCredential() bool {
	return a.allowedCredential
}

// NewApplicationGinImpl ...
// Deprecated:
func NewApplicationGinImpl(ctx context.Context) (*ApplicationGinImpl, error) {
	return &ApplicationGinImpl{
		allowedOrigins:    strings.Split(os.Getenv("ALLOWED_ORIGINS"), ","),
		allowedMethods:    strings.Split(os.Getenv("ALLOWED_METHODS"), ","),
		allowedHeaders:    strings.Split(os.Getenv("ALLOWED_HEADERS"), ","),
		exposeHeaders:     strings.Split(os.Getenv("EXPOSE_HEADERS"), ","),
		allowedCredential: env.GetenvAsBool("ALLOWED_CREDENTIAL", false),
	}, nil
}
