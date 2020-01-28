package cgcp

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/suzuito/common-go/application"
)

type ApplicationGCP interface {
	application.Application
	AppFirebase() *firebase.App
}

type ApplicationGCPImpl struct {
	appFirebase *firebase.App
}

func NewApplicationGCPImpl(ctx context.Context) (*ApplicationGCPImpl, error) {
	appFirebase, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &ApplicationGCPImpl{
		appFirebase: appFirebase,
	}, nil
}

func (a *ApplicationGCPImpl) AppFirebase() *firebase.App {
	return a.appFirebase
}
