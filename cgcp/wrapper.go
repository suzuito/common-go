package cgcp

import (
	"context"

	"cloud.google.com/go/firestore"
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
	Batch() *firestore.WriteBatch
	Close() error
	Collection(path string) *firestore.CollectionRef
	CollectionGroup(collectionID string) *firestore.CollectionGroupRef
	Collections(ctx context.Context) *firestore.CollectionIterator
	Doc(path string) *firestore.DocumentRef
	GetAll(ctx context.Context, docRefs []*firestore.DocumentRef) (_ []*firestore.DocumentSnapshot, err error)
	RunTransaction(ctx context.Context, f func(context.Context, *firestore.Transaction) error, opts ...firestore.TransactionOption) (err error)
}

// FirebaseAuthClient ...
type FirebaseAuthClient interface {
	VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error)
}
