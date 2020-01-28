package cweb

import (
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/suzuito/common-go/clogger"
	"github.com/suzuito/common-go/cgcp"
)

// HO ...
type HO struct {
	FirestoreClientNotUse bool
	LoggerNotUse          bool
}

// H ...
func H(
	app cgcp.ApplicationGCP,
	ctx *gin.Context,
	proc func(
		logger clogger.Logger,
		fcli *firestore.Client,
	) error,
	opt *HO,
) {
	var logger clogger.Logger
	var fcli *firestore.Client
	var err error
	if opt == nil || opt.LoggerNotUse == false {
		logger = app.Logger(ctx)
		defer logger.Close()
	}
	if opt == nil || opt.FirestoreClientNotUse == false {
		fcli, err = app.AppFirebase().Firestore(ctx)
		if err != nil {
			logger.Errorf("%+v", err)
			Abort(ctx, NewHTTPError(http.StatusInternalServerError, "InternalServerError", err))
			return
		}
		defer fcli.Close()
	}
	if err := proc(logger, fcli); err != nil {
		logger.Errorf("%+v", err)
	}
}
