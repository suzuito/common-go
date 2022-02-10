package cgcp

import (
	"context"
	"os"
	"strings"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"golang.org/x/xerrors"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

type SecretClientGCP struct {
	cli *secretmanager.Client
}

func (c *SecretClientGCP) Close() error {
	return c.cli.Close()
}

func (c *SecretClientGCP) Get(ctx context.Context, name string) ([]byte, error) {
	req := secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}
	result, err := c.cli.AccessSecretVersion(ctx, &req)
	if err != nil {
		return nil, xerrors.Errorf("Cannot get secret '%s' : %w", name, err)
	}
	return result.Payload.Data, nil
}

func (c *SecretClientGCP) GetString(ctx context.Context, name string) (string, error) {
	b, err := c.Get(ctx, name)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *SecretClientGCP) ReplaceEnv(ctx context.Context, ekey string) error {
	evalue := os.Getenv(ekey)
	b, err := c.Get(ctx, evalue)
	if err != nil {
		return err
	}
	if err := os.Setenv(ekey, string(b)); err != nil {
		return err
	}
	return nil
}

func (c *SecretClientGCP) ReplaceAllEnvs(ctx context.Context) error {
	prefix := "GCP_SECRET__"
	envs := os.Environ()
	for _, env := range envs {
		parts := strings.Split(env, "=")
		if len(parts) < 2 {
			continue
		}
		prevKey := parts[0]
		if !strings.HasPrefix(prevKey, prefix) {
			continue
		}
		secretName := parts[1]
		nextKey := strings.Replace(prevKey, prefix, "", -1)
		os.Setenv(nextKey, secretName)
		if err := c.ReplaceEnv(ctx, nextKey); err != nil {
			return xerrors.Errorf("cannot replace env : %w", err)
		}
	}
	return nil
}

func NewSecretClientGCP(ctx context.Context) (*SecretClientGCP, error) {
	cli, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, xerrors.Errorf("Cannot new secret manager client : %w", err)
	}
	return &SecretClientGCP{
		cli: cli,
	}, nil
}
