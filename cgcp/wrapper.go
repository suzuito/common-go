package cgcp

import (
	"context"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
)

// FirebaseApp ...
type FirebaseApp interface {
	Firestore(ctx context.Context) (FirebaseFirestoreClient, error)
	Auth(ctx context.Context) (FirebaseAuthClient, error)
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
