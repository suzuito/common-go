package cgcp

import (
	"context"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/go-redis/redis/v7"
	"golang.org/x/xerrors"
)

type GCPResource struct {
	GCS  *storage.Client
	GCF  *firestore.Client
	GCA  *auth.Client
	GCPS *pubsub.Client
	GMS  MemoryStoreClient
}

func (r *GCPResource) Close() []error {
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

type GCPResourceGenerator struct {
	newGCS        bool
	newGCF        bool
	projectIDGCF  string
	newGCA        bool
	newGMS        bool
	newGCPS       bool
	projectIDGCPS string
	redisClient   *redis.Client
	redisTTL      int
}

func NewGCPResourceGenerator() *GCPResourceGenerator {
	return &GCPResourceGenerator{
		newGCS:      false,
		newGMS:      false,
		newGCF:      false,
		newGCA:      false,
		newGCPS:     false,
		redisClient: nil,
	}
}

func (r *GCPResourceGenerator) GCS() *GCPResourceGenerator {
	r.newGCS = true
	return r
}

func (r *GCPResourceGenerator) GMS(cli *redis.Client, ttl int) *GCPResourceGenerator {
	r.redisClient = cli
	r.redisTTL = ttl
	r.newGMS = true
	return r
}

func (r *GCPResourceGenerator) GCF(projectID string) *GCPResourceGenerator {
	r.projectIDGCF = projectID
	r.newGCF = true
	return r
}

func (r *GCPResourceGenerator) GCPS(projectID string) *GCPResourceGenerator {
	r.newGCPS = true
	r.projectIDGCPS = projectID
	return r
}

func (r *GCPResourceGenerator) GCA() *GCPResourceGenerator {
	r.newGCA = true
	return r
}

func (r *GCPResourceGenerator) Gen(ctx context.Context) (*GCPResource, error) {
	var err error
	ret := GCPResource{}
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
		ret.GCF, err = firestore.NewClient(ctx, r.projectIDGCF)
		if err != nil {
			return nil, xerrors.Errorf("Cannot firestore.NewClient : %w", err)
		}
	}
	if r.newGCA {
		var app *firebase.App
		app, err = firebase.NewApp(ctx, nil)
		if err != nil {
			return nil, xerrors.Errorf("Cannot firebase.NewApp : %w", err)
		}
		if r.newGCA {
			ret.GCA, err = app.Auth(ctx)
			if err != nil {
				return nil, xerrors.Errorf("Cannot app.Auth : %w", err)
			}
		}
	}
	if r.newGMS {
		ret.GMS = NewMemoryStoreClientRedis(r.redisClient, r.redisTTL)
	}
	if r.newGCPS {
		ret.GCPS, err = pubsub.NewClient(ctx, r.projectIDGCPS)
		if err != nil {
			return nil, xerrors.Errorf("Cannot pubsub.NewClient : %w", err)
		}
	}
	return &ret, nil
}
