package cgcp

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/go-redis/redis/v7"
	"golang.org/x/xerrors"
)

type GCPContextResource struct {
	GCS *storage.Client
	GMS MemoryStoreClient
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
	return errs
}

type GCPContextResourceGenerator struct {
	newGCS      bool
	newGMS      bool
	redisClient *redis.Client
	redisTTL    int
}

func NewGCPContextResourceGenerator() *GCPContextResourceGenerator {
	return &GCPContextResourceGenerator{
		newGCS:      false,
		newGMS:      false,
		redisClient: nil,
	}
}

func (r *GCPContextResourceGenerator) GCS() {
	r.newGCS = true
}

func (r *GCPContextResourceGenerator) GMS(cli *redis.Client, ttl int) {
	r.redisClient = cli
	r.redisTTL = ttl
	r.newGMS = true
}

func (r *GCPContextResourceGenerator) Gen(ctx context.Context) (*GCPContextResource, error) {
	var err error
	ret := GCPContextResource{}
	if r.newGCS {
		ret.GCS, err = storage.NewClient(ctx)
		if err != nil {
			return nil, xerrors.Errorf("Cannot storage.NewClient : %w", err)
		}
	}
	if r.newGMS {
		ret.GMS = NewMemoryStoreClientRedis(r.redisClient, r.redisTTL)
	}
	return &ret, nil
}
