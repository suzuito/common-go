package gcp

import (
	firebase "firebase.google.com/go"
	"github.com/suzuito/common-go/application"
)

type ApplicationGCP interface {
	application.Application
	AppFirebase() *firebase.App
}

// func NewApplicationGCP(ctx context.Context) (*ApplicationGCP, error) {
// 	appFirebase, err := firebase.NewApp(ctx, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &ApplicationGCP{
// 		appFirebase: appFirebase,
// 	}, nil
// }
//
// func (a *ApplicationGCP) AppFirebase() *firebase.App {
// 	return a.appFirebase
// }
