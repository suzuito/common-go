package application

import (
	"context"

	"github.com/suzuito/common-go/clogger"
)

type Application interface {
	Logger(ctx context.Context) clogger.Logger
}
