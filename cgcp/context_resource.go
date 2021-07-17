package cgcp

import (
	"context"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/go-redis/redis/v7"
	"golang.org/x/xerrors"
)

type GCPContextResource struct {
	GCS  *storage.Client
	GCF  *firestore.Client
	GCPS *pubsub.Client
	GMS  MemoryStoreClient
}

func (r *GCPContextResource) Close() []error {
	errs := []error{}
	if r.GCS != nil {
		if err := r.GCS.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	if r.GMS != nil {
		if err := r.GMS.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	if r.GCF != nil {
		if err := r.GCF.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	if r.GCPS != nil {
		if err := r.GCPS.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

type GCPContextResourceGenerator struct {
	newGCS        bool
	newGCF        bool
	newGMS        bool
	newGCPS       bool
	ProjectIDGCPS string
	redisClient   *redis.Client
	redisTTL      int
}

func NewGCPContextResourceGenerator() *GCPContextResourceGenerator {
	return &GCPContextResourceGenerator{
		newGCS:      false,
		newGMS:      false,
		newGCF:      false,
		newGCPS:     false,
		redisClient: nil,
	}
}

func (r *GCPContextResourceGenerator) GCS() *GCPContextResourceGenerator {
	r.newGCS = true
}

func (r *GCPContextResourceGenerator) GMS(cli *redis.Client, ttl int) *GCPContextResourceGenerator {
	r.redisClient = cli
	r.redisTTL = ttl
	r.newGMS = true
}

func (r *GCPContextResourceGenerator) GCF() *GCPContextResourceGenerator {
	r.newGCF = true
}

func (r *GCPContextResourceGenerator) GCPS(projectID string) *GCPContextResourceGenerator {
	r.newGCPS = true
	r.ProjectIDGCPS = projectID
}

func (r *GCPContextResourceGenerator) Gen(ctx context.Context) (*GCPContextResource, error) {
	var err error
	ret := GCPContextResource{}
	defer func() {
		if err != nil {
			ret.Close()
		}
	}()
	if r.newGCS {
		ret.GCS, err = storage.NewClient(ctx)
		if err != nil {
			return nil, xerrors.Errorf("Cannot storage.NewClient : %w", err)
		}
	}
	if r.newGCF {
		var app *firebase.App
		app, err = firebase.NewApp(ctx, nil)
		if err != nil {
			return nil, xerrors.Errorf("Cannot firebase.NewApp : %w", err)
		}
		if r.newGCF {
			ret.GCF, err = app.Firestore(ctx)
			if err != nil {
				return nil, xerrors.Errorf("Cannot app.Firestore : %w", err)
			}
		}
	}
	if r.newGMS {
		ret.GMS = NewMemoryStoreClientRedis(r.redisClient, r.redisTTL)
	}
	if r.newGCPS {
		ret.GCPS, err = pubsub.NewClient(ctx, r.ProjectIDGCPS)
		if err != nil {
			return nil, xerrors.Errorf("Cannot pubsub.NewClient : %w", err)
		}
	}
	return &ret, nil
}
