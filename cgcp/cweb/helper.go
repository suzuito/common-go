package cweb

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/common-go/application"
	"github.com/suzuito/common-go/cgcp"
	"github.com/suzuito/common-go/cgin"
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
	ctx *gin.Context,
	app application.ApplicationLogger,
	appFirebase cgcp.FirebaseApp,
	proc func(
		logger clogger.Logger,
		fcli cgcp.FirebaseFirestoreClient,
		fauth cgcp.FirebaseAuthClient,
	) error,
	opt *HO,
) {
	var logger clogger.Logger
	var fcli cgcp.FirebaseFirestoreClient
	var fauth cgcp.FirebaseAuthClient
	var err error
	if opt == nil || opt.LoggerNotUse == false {
		logger = app.Logger(ctx)
		defer logger.Close()
	}
	if opt == nil || opt.FirestoreClientNotUse == false || opt.AuthClientNotUse == false {
		if opt == nil || opt.FirestoreClientNotUse == false {
			fcli, err = appFirebase.Firestore(ctx)
			if err != nil {
				if logger != nil {
					logger.Errorf("%+v", err)
				}
				cgin.Abort(ctx, cgin.NewHTTPError(http.StatusInternalServerError, "InternalServerError", err))
				return
			}
			defer fcli.Close()
		}
		if opt == nil || opt.AuthClientNotUse == false {
			fauth, err = appFirebase.Auth(ctx)
			if err != nil {
				if logger != nil {
					logger.Errorf("%+v", err)
				}
				cgin.Abort(ctx, cgin.NewHTTPError(http.StatusInternalServerError, "InternalServerError", err))
				return
			}
		}
	}
	if err := proc(logger, fcli, fauth); err != nil {
		if logger != nil {
			logger.Errorf("%+v", err)
		}
	}
}
