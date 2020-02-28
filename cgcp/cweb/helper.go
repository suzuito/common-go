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
	PubSubClientNotUse    bool
}

// H ...
func H(
	ctx *gin.Context,
	app application.ApplicationLogger,
	appFirebase cgcp.FirebaseApp,
	appGCP cgcp.GCPApp,
	proc func(
		logger clogger.Logger,
		fcli cgcp.FirebaseFirestoreClient,
		fauth cgcp.FirebaseAuthClient,
		pcli cgcp.GCPPubSubClient,
	) error,
	opt *HO,
) {
	var logger clogger.Logger
	var fcli cgcp.FirebaseFirestoreClient
	var fauth cgcp.FirebaseAuthClient
	var pcli cgcp.GCPPubSubClient
	var err error
	if opt == nil || opt.LoggerNotUse == false {
		logger = app.Logger(ctx)
		defer logger.Close()
	}
	if opt == nil || opt.FirestoreClientNotUse == false || opt.AuthClientNotUse == false || opt.PubSubClientNotUse == false {
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
		if opt == nil || opt.PubSubClientNotUse == false {
			pcli, err = appGCP.PubSub(ctx)
			if err != nil {
				if logger != nil {
					logger.Errorf("%+v", err)
				}
				cgin.Abort(ctx, cgin.NewHTTPError(http.StatusInternalServerError, "InternalServerError", err))
				return
			}
			defer pcli.Close()
		}
	}
	if err := proc(logger, fcli, fauth, pcli); err != nil {
		if logger != nil {
			logger.Errorf("%+v", err)
		}
	}
}
