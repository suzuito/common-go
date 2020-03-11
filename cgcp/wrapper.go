package cgcp

import (
	"context"

	"cloud.google.com/go/pubsub"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/go-redis/redis/v7"
	env "github.com/suzuito/common-env"
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
	MemoryStore(ctx context.Context) (MemoryStoreClient, error)
}

// GCPAppImpl ...
type GCPAppImpl struct {
	redisClient *redis.Client
	redisTTL    int
	projectID   string
}

// NewGCPAppImpl ...
func NewGCPAppImpl(
	ctx context.Context,
	redisClient *redis.Client,
	projectID string,
) (*GCPAppImpl, error) {
	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr: env.GetenvAsString("REDIS_ADDR", "localhost:6379"),
	// })
	return &GCPAppImpl{
		redisClient: redisClient,
		redisTTL:    env.GetenvAsInt("REDIS_TTL", 1800),
		projectID:   projectID,
	}, nil
}

// PubSub ...
func (f *GCPAppImpl) PubSub(ctx context.Context) (GCPPubSubClient, error) {
	return pubsub.NewClient(ctx, f.projectID)
}

// MemoryStore ...
func (f *GCPAppImpl) MemoryStore(ctx context.Context) (MemoryStoreClient, error) {
	return NewMemoryStoreClientRedis(f.redisClient, f.redisTTL), nil
}

// GCPPubSubClient ...
type GCPPubSubClient interface {
	Close() error
}
