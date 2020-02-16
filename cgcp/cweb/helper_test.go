package cweb

import (
	"net/http/httptest"
	"testing"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/suzuito/common-go/application_mock"
	"github.com/suzuito/common-go/clogger"
	"github.com/suzuito/common-go/clogger_mock"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc      string
		setUpMock func(
			app *application_mock.MockApplicationLogger,
			logger *clogger_mock.MockLogger,
			appFirebase *cweb_mock.MockAppFirebaseWrapper,
		)
		inputProc func(
			logger clogger.Logger,
			fcli *firestore.Client,
			fauth *auth.Client,
		) error
		inputOpt *HO
	}{
		{
			desc: "",
			setUpMock: func(
				app *application_mock.MockApplicationLogger,
				logger *clogger_mock.MockLogger,
				appFirebase *cweb_mock.MockAppFirebaseWrapper,
			) {
				app.
					EXPECT().
					Logger(gomock.Any()).
					Return(logger)
				logger.
					EXPECT().
					Close()
			},
			inputProc: func(
				logger clogger.Logger,
				fcli *firestore.Client,
				fauth *auth.Client,
			) error {
				return nil
			},
			inputOpt: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrlApp := gomock.NewController(t)
			defer ctrlApp.Finish()
			app := application_mock.NewMockApplicationLogger(ctrlApp)
			ctrlLogger := gomock.NewController(t)
			defer ctrlLogger.Finish()
			logger := clogger_mock.NewMockLogger(ctrlLogger)
			ctrlAppFirebase := gomock.NewController(t)
			defer ctrlAppFirebase.Finish()
			appFirebase := cweb_mock.NewMockAppFirebaseWrapper(ctrlAppFirebase)
			tC.setUpMock(app, logger, appFirebase)
			// _, _ := http.NewRequest("GET", "/dummy", nil)
			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)
			H(
				ctx,
				app,
				appFirebase,
				tC.inputProc,
				tC.inputOpt,
			)
		})
	}
}
