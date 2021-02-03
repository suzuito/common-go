package cgcp

import (
	"context"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"golang.org/x/xerrors"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

type SecretClientGCP struct {
	cli *secretmanager.Client
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
	return string(b)
}

func NewSecretClientGCP(ctx context.Context) (*SecretClientGCP, error) {
	cli, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, xerrors.Errorf("Cannot new secret manager client")
	}
	return &SecretClientGCP{
		cli: cli,
	}, nil
}
