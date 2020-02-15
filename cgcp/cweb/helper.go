package cweb

import (
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/suzuito/common-go/application"
	"github.com/suzuito/common-go/clogger"
)

// HO ...
type HO struct {
	FirestoreClientNotUse bool
	AuthClientNotUse      bool
	LoggerNotUse          bool
}

// H ...
func H(
	app application.ApplicationLogger,
	ctx *gin.Context,
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
	var appFirebase *firebase.App
	if opt == nil || opt.FirestoreClientNotUse == false || opt.AuthClientNotUse == false {
		appFirebase, err = firebase.NewApp(ctx, nil)
		if err != nil {
			logger.Errorf("%+v", err)
			Abort(ctx, NewHTTPError(http.StatusInternalServerError, "InternalServerError", err))
			return
		}
		if opt.FirestoreClientNotUse == false {
			fcli, err = appFirebase.Firestore(ctx)
			if !opt.FirestoreClientNotUse {
				if err != nil {
					logger.Errorf("%+v", err)
					Abort(ctx, NewHTTPError(http.StatusInternalServerError, "InternalServerError", err))
					return
				}
				defer fcli.Close()
			}
			if opt.AuthClientNotUse == false {
				if !opt.AuthClientNotUse {
					fauth, err = appFirebase.Auth(ctx)
					if err != nil {
						logger.Errorf("%+v", err)
						Abort(ctx, NewHTTPError(http.StatusInternalServerError, "InternalServerError", err))
						return
					}
				}
			}
		}
		if err := proc(logger, fcli, fauth); err != nil {
			logger.Errorf("%+v", err)
		}
	}
}
