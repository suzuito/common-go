package cweb

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/suzuito/common-go/application"
	"github.com/suzuito/common-go/cgin"
	"github.com/suzuito/common-go/clogger"
)

// FirebaseApp ...
type FirebaseApp interface {
	Firestore(ctx context.Context) (*firestore.Client, error)
	Auth(ctx context.Context) (*auth.Client, error)
}

// HO ...
type HO struct {
	FirestoreClientNotUse bool
	AuthClientNotUse      bool
	LoggerNotUse          bool
}

// H ...
func H(
	ctx *gin.Context,
	app application.ApplicationLogger,
	appFirebase FirebaseApp,
	proc func(
		logger clogger.Logger,
		fcli *firestore.Client,
		fauth *auth.Client,
	) error,
	opt *HO,
) {
	var logger clogger.Logger
	var fcli *firestore.Client
	var fauth *auth.Client
	var err error
	if opt == nil || opt.LoggerNotUse == false {
		logger = app.Logger(ctx)
		defer logger.Close()
	}
	if opt == nil || opt.FirestoreClientNotUse == false || opt.AuthClientNotUse == false {
		if err != nil {
			logger.Errorf("%+v", err)
			cgin.Abort(ctx, cgin.NewHTTPError(http.StatusInternalServerError, "InternalServerError", err))
			return
		}
		if opt == nil || opt.FirestoreClientNotUse == false {
			fcli, err = appFirebase.Firestore(ctx)
			if err != nil {
				logger.Errorf("%+v", err)
				cgin.Abort(ctx, cgin.NewHTTPError(http.StatusInternalServerError, "InternalServerError", err))
				return
			}
			defer fcli.Close()
		}
		if opt == nil || opt.AuthClientNotUse == false {
			fauth, err = appFirebase.Auth(ctx)
			if err != nil {
				logger.Errorf("%+v", err)
				cgin.Abort(ctx, cgin.NewHTTPError(http.StatusInternalServerError, "InternalServerError", err))
				return
			}
		}
		if err := proc(logger, fcli, fauth); err != nil {
			logger.Errorf("%+v", err)
		}
	}
}
