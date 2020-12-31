package cgcp

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v7"
	"golang.org/x/xerrors"
)

var (
	// ErrNotFound ...
	ErrNotFound = fmt.Errorf("Not found key")
)

// MemoryStoreClient ...
type MemoryStoreClient interface {
	GetJSON(ctx context.Context, key string, value interface{}) error
	SetJSON(ctx context.Context, key string, value interface{}) error
	GetInt(ctx context.Context, key string, value *int) error
	SetInt(ctx context.Context, key string, value int) error
	GetString(ctx context.Context, key string, value *string) error
	SetString(ctx context.Context, key string, value string) error
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
		return xerrors.Errorf("Not found key '%s' : %w", ErrNotFound)
	} else if err != nil {
		return err
	}
	return nil
}

func (c *MemoryStoreClientRedis) setString(ctx context.Context, key string, value string) error {
	return c.cli.Set(key, value, time.Duration(c.ttl)*time.Second).Err()
}

// GetJSON ...
func (c *MemoryStoreClientRedis) GetJSON(ctx context.Context, key string, value interface{}) error {
	valueString := ""
	if err := c.getString(ctx, key, &valueString); err != nil {
		return err
	}
	return json.Unmarshal([]byte(valueString), value)
}

// SetJSON ...
func (c *MemoryStoreClientRedis) SetJSON(ctx context.Context, key string, value interface{}) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.setString(ctx, key, string(b))
}

// GetInt ...
func (c *MemoryStoreClientRedis) GetInt(ctx context.Context, key string, value *int) error {
	var err error
	v := ""
	if err = c.getString(ctx, key, &v); err != nil {
		return err
	}
	*value, err = strconv.Atoi(v)
	return err
}

// SetInt ...
func (c *MemoryStoreClientRedis) SetInt(ctx context.Context, key string, value int) error {
	var err error
	v := strconv.Itoa(value)
	if err = c.setString(ctx, key, v); err != nil {
		return err
	}
	return err
}

// GetString ...
func (c *MemoryStoreClientRedis) GetString(ctx context.Context, key string, value *string) error {
	return c.getString(ctx, key, value)
}

// SetString ...
func (c *MemoryStoreClientRedis) SetString(ctx context.Context, key string, value string) error {
	return c.setString(ctx, key, value)
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
