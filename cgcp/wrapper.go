package cgcp

import (
	"context"

	"cloud.google.com/go/pubsub"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

// FirebaseApp ...
type FirebaseApp interface {
	Firestore(ctx context.Context) (FirebaseFirestoreClient, error)
	Auth(ctx context.Context) (FirebaseAuthClient, error)
}

// FirebaseAppImpl ...
type FirebaseAppImpl struct {
	app *firebase.App
}

// Firestore ...
func (f *FirebaseAppImpl) Firestore(ctx context.Context) (FirebaseFirestoreClient, error) {
	return f.app.Firestore(ctx)
}

// Auth ...
func (f *FirebaseAppImpl) Auth(ctx context.Context) (FirebaseAuthClient, error) {
	return f.app.Auth(ctx)
}

// NewFirebaseApp ...
func NewFirebaseApp(ctx context.Context) (*FirebaseAppImpl, error) {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &FirebaseAppImpl{
		app: app,
	}, nil
}

// FirebaseFirestoreClient ...
type FirebaseFirestoreClient interface {
	Close() error
}

// FirebaseAuthClient ...
type FirebaseAuthClient interface {
	VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error)
}

// GCPApp ...
type GCPApp interface {
	PubSub(ctx context.Context) (GCPPubSubClient, error)
}

// GCPAppImpl ...
type GCPAppImpl struct {
	projectID string
}

// NewGCPAppImpl ...
func NewGCPAppImpl(ctx context.Context, projectID string) (*GCPAppImpl, error) {
	return &GCPAppImpl{
		projectID: projectID,
	}, nil
}

// PubSub ...
func (f *GCPAppImpl) PubSub(ctx context.Context) (GCPPubSubClient, error) {
	return pubsub.NewClient(ctx, f.projectID)
}

// GCPPubSubClient ...
type GCPPubSubClient interface {
	Close() error
}
