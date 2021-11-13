package csecret

import "context"

type SecretClient interface {
	Get(ctx context.Context, name string) ([]byte, error)
	GetString(ctx context.Context, name string) (string, error)
	ReplaceEnv(ctx context.Context, ekey string) error
	ReplaceAllEnvs(ctx context.Context) error
}
