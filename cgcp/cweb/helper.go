package cweb

import (
	"net/http"

	"github.com/gin-gonic/gin"
	env "github.com/suzuito/common-env"
	"github.com/suzuito/common-go/application"
	"github.com/suzuito/common-go/cgcp"
	"github.com/suzuito/common-go/cgin"
	"github.com/suzuito/common-go/clogger"
)

// HO ...
type HO struct {
	FirestoreClientNotUse   bool
	AuthClientNotUse        bool
	LoggerNotUse            bool
	PubSubClientNotUse      bool
	MemoryStoreClientNotUse bool
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
		mcli cgcp.MemoryStoreClient,
	) error,
	opt *HO,
) {
	var logger clogger.Logger
	var fcli cgcp.FirebaseFirestoreClient
	var fauth cgcp.FirebaseAuthClient
	var pcli cgcp.GCPPubSubClient
	var mcli cgcp.MemoryStoreClient
	var err error
	if opt == nil || opt.LoggerNotUse == false {
		logger, err = cgcp.NewLoggerGCP2(ctx, env.GetenvAsString("GOOGLE_CLOUD_PROJECT", ""), ctx.Request)
		if err != nil {
			logger = &clogger.LoggerPrint{}
		}
		defer func() {
			logger.Request(ctx.Request)
			logger.Close()
		}()
	}
	if opt == nil || opt.FirestoreClientNotUse == false || opt.AuthClientNotUse == false || opt.PubSubClientNotUse == false || opt.MemoryStoreClientNotUse == false {
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
		if opt == nil || opt.MemoryStoreClientNotUse == false {
			mcli, err = appGCP.MemoryStore(ctx)
			if err != nil {
				if logger != nil {
					logger.Errorf("%+v", err)
				}
				cgin.Abort(ctx, cgin.NewHTTPError(http.StatusInternalServerError, "InternalServerError", err))
				return
			}
			defer mcli.Close()
		}
	}
	if err := proc(logger, fcli, fauth, pcli, mcli); err != nil {
		if logger != nil {
			logger.Errorf("%+v", err)
		}
	}
}
