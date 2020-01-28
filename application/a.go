package application

import (
	"context"

	"github.com/suzuito/common-go/clogger"
)

type ApplicationLogger interface {
	Logger(ctx context.Context) clogger.Logger
}
