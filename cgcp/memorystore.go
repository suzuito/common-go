package cgcp

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/suzuito/common-go/werror"
)

var (
	// ErrNotFound ...
	ErrNotFound = fmt.Errorf("Not found key")
)

// MemoryStoreClient ...
type MemoryStoreClient interface {
	GetJSON(ctx context.Context, key string, value interface{}) error
	SetJSON(ctx context.Context, key string, value interface{}) error
	Delete(ctx context.Context, keys ...string) error
	Close() error
}

// MemoryStoreClientRedis ...
type MemoryStoreClientRedis struct {
	cli *redis.Client
	ttl int
}

// NewMemoryStoreClientRedis ...
func NewMemoryStoreClientRedis(cli *redis.Client, ttl int) *MemoryStoreClientRedis {
	return &MemoryStoreClientRedis{
		cli: cli,
		ttl: ttl,
	}
}

func (c *MemoryStoreClientRedis) getString(ctx context.Context, key string, value *string) error {
	var err error
	*value, err = c.cli.Get(key).Result()
	if err == redis.Nil {
		return werror.Newf(ErrNotFound, "Not found key '%s'", key)
	} else if err != nil {
		return err
	}
	return nil
}

// GetJSON ...
func (c *MemoryStoreClientRedis) GetJSON(ctx context.Context, key string, value interface{}) error {
	valueString := ""
	if err := c.getString(ctx, key, &valueString); err != nil {
		return err
	}
	return json.Unmarshal([]byte(valueString), value)
}

func (c *MemoryStoreClientRedis) setString(ctx context.Context, key string, value string) error {
	return c.cli.Set(key, value, time.Duration(c.ttl)*time.Second).Err()
}

// SetJSON ...
func (c *MemoryStoreClientRedis) SetJSON(ctx context.Context, key string, value interface{}) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.setString(ctx, key, string(b))
}

// Delete ...
func (c *MemoryStoreClientRedis) Delete(ctx context.Context, keys ...string) error {
	if err := c.cli.Del(keys...).Err(); err != nil {
		if err == redis.Nil {
			return nil
		}
		return err
	}
	return nil
}

// Close ...
func (c *MemoryStoreClientRedis) Close() error {
	return nil
}
