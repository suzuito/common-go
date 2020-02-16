package cweb

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/suzuito/common-go/application"
	"github.com/suzuito/common-go/cgcp"
	"github.com/suzuito/common-go/clogger"
	"gopkg.in/go-playground/assert.v1"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc      string
		setUpMock func(
			app *application.MockApplicationLogger,
			logger *clogger.MockLogger,
			appFirebase *cgcp.MockFirebaseApp,
			dcliFirebase *cgcp.MockFirebaseFirestoreClient,
			acliFirebase *cgcp.MockFirebaseAuthClient,
		)
		inputProc func(
			logger clogger.Logger,
			fcli cgcp.FirebaseFirestoreClient,
			fauth cgcp.FirebaseAuthClient,
		) error
		inputOpt         *HO
		expectedResponse *http.Response
	}{
		{
			desc: "If opt is nil",
			setUpMock: func(
				app *application.MockApplicationLogger,
				logger *clogger.MockLogger,
				appFirebase *cgcp.MockFirebaseApp,
				dcliFirebase *cgcp.MockFirebaseFirestoreClient,
				acliFirebase *cgcp.MockFirebaseAuthClient,
			) {
				app.
					EXPECT().Logger(gomock.Any()).Return(logger)
				logger.
					EXPECT().Close()
				appFirebase.
					EXPECT().Firestore(gomock.Any()).Return(dcliFirebase, nil)
				dcliFirebase.
					EXPECT().Close()
				appFirebase.
					EXPECT().Auth(gomock.Any()).Return(acliFirebase, nil)
			},
			inputProc: func(
				logger clogger.Logger,
				fcli cgcp.FirebaseFirestoreClient,
				fauth cgcp.FirebaseAuthClient,
			) error {
				return nil
			},
			inputOpt: nil,
		},
		{
			desc: "Use all",
			setUpMock: func(
				app *application.MockApplicationLogger,
				logger *clogger.MockLogger,
				appFirebase *cgcp.MockFirebaseApp,
				dcliFirebase *cgcp.MockFirebaseFirestoreClient,
				acliFirebase *cgcp.MockFirebaseAuthClient,
			) {
				app.
					EXPECT().Logger(gomock.Any()).Return(logger)
				logger.
					EXPECT().Close()
				appFirebase.
					EXPECT().Firestore(gomock.Any()).Return(dcliFirebase, nil)
				dcliFirebase.
					EXPECT().Close()
				appFirebase.
					EXPECT().Auth(gomock.Any()).Return(acliFirebase, nil)
			},
			inputProc: func(
				logger clogger.Logger,
				fcli cgcp.FirebaseFirestoreClient,
				fauth cgcp.FirebaseAuthClient,
			) error {
				return nil
			},
			inputOpt: &HO{
				FirestoreClientNotUse: false,
				AuthClientNotUse:      false,
				LoggerNotUse:          false,
			},
		},
		{
			desc: "Use firestore client only",
			setUpMock: func(
				app *application.MockApplicationLogger,
				logger *clogger.MockLogger,
				appFirebase *cgcp.MockFirebaseApp,
				dcliFirebase *cgcp.MockFirebaseFirestoreClient,
				acliFirebase *cgcp.MockFirebaseAuthClient,
			) {
				appFirebase.
					EXPECT().Firestore(gomock.Any()).Return(dcliFirebase, nil)
				dcliFirebase.
					EXPECT().Close()
			},
			inputProc: func(
				logger clogger.Logger,
				fcli cgcp.FirebaseFirestoreClient,
				fauth cgcp.FirebaseAuthClient,
			) error {
				return nil
			},
			inputOpt: &HO{
				FirestoreClientNotUse: false,
				AuthClientNotUse:      true,
				LoggerNotUse:          true,
			},
		},
		{
			desc: "Use firestore auth only",
			setUpMock: func(
				app *application.MockApplicationLogger,
				logger *clogger.MockLogger,
				appFirebase *cgcp.MockFirebaseApp,
				dcliFirebase *cgcp.MockFirebaseFirestoreClient,
				acliFirebase *cgcp.MockFirebaseAuthClient,
			) {
				appFirebase.
					EXPECT().Auth(gomock.Any()).Return(acliFirebase, nil)
			},
			inputProc: func(
				logger clogger.Logger,
				fcli cgcp.FirebaseFirestoreClient,
				fauth cgcp.FirebaseAuthClient,
			) error {
				return nil
			},
			inputOpt: &HO{
				FirestoreClientNotUse: true,
				AuthClientNotUse:      false,
				LoggerNotUse:          true,
			},
		},
		{
			desc: "Use firestore logger only",
			setUpMock: func(
				app *application.MockApplicationLogger,
				logger *clogger.MockLogger,
				appFirebase *cgcp.MockFirebaseApp,
				dcliFirebase *cgcp.MockFirebaseFirestoreClient,
				acliFirebase *cgcp.MockFirebaseAuthClient,
			) {
				app.
					EXPECT().Logger(gomock.Any()).Return(logger)
				logger.
					EXPECT().Close()
			},
			inputProc: func(
				logger clogger.Logger,
				fcli cgcp.FirebaseFirestoreClient,
				fauth cgcp.FirebaseAuthClient,
			) error {
				return nil
			},
			inputOpt: &HO{
				FirestoreClientNotUse: true,
				AuthClientNotUse:      true,
				LoggerNotUse:          false,
			},
		},
		{
			desc: "Error Firestore",
			setUpMock: func(
				app *application.MockApplicationLogger,
				logger *clogger.MockLogger,
				appFirebase *cgcp.MockFirebaseApp,
				dcliFirebase *cgcp.MockFirebaseFirestoreClient,
				acliFirebase *cgcp.MockFirebaseAuthClient,
			) {
				app.
					EXPECT().Logger(gomock.Any()).Return(logger)
				logger.
					EXPECT().Close()
				appFirebase.
					EXPECT().Firestore(gomock.Any()).Return(dcliFirebase, fmt.Errorf("Dummy error"))
				logger.
					EXPECT().Errorf(gomock.Any(), gomock.Any())
			},
			inputProc: func(
				logger clogger.Logger,
				fcli cgcp.FirebaseFirestoreClient,
				fauth cgcp.FirebaseAuthClient,
			) error {
				return nil
			},
			inputOpt: &HO{
				FirestoreClientNotUse: false,
				AuthClientNotUse:      true,
				LoggerNotUse:          false,
			},
			expectedResponse: &http.Response{
				StatusCode: http.StatusInternalServerError,
			},
		},
		{
			desc: "Error Auth",
			setUpMock: func(
				app *application.MockApplicationLogger,
				logger *clogger.MockLogger,
				appFirebase *cgcp.MockFirebaseApp,
				dcliFirebase *cgcp.MockFirebaseFirestoreClient,
				acliFirebase *cgcp.MockFirebaseAuthClient,
			) {
				app.
					EXPECT().Logger(gomock.Any()).Return(logger)
				logger.
					EXPECT().Close()
				appFirebase.
					EXPECT().Auth(gomock.Any()).Return(acliFirebase, fmt.Errorf("Dummy error"))
				logger.
					EXPECT().Errorf(gomock.Any(), gomock.Any())
			},
			inputProc: func(
				logger clogger.Logger,
				fcli cgcp.FirebaseFirestoreClient,
				fauth cgcp.FirebaseAuthClient,
			) error {
				return nil
			},
			inputOpt: &HO{
				FirestoreClientNotUse: true,
				AuthClientNotUse:      false,
				LoggerNotUse:          false,
			},
			expectedResponse: &http.Response{
				StatusCode: http.StatusInternalServerError,
			},
		},
		{
			desc: "Error proc",
			setUpMock: func(
				app *application.MockApplicationLogger,
				logger *clogger.MockLogger,
				appFirebase *cgcp.MockFirebaseApp,
				dcliFirebase *cgcp.MockFirebaseFirestoreClient,
				acliFirebase *cgcp.MockFirebaseAuthClient,
			) {
				app.
					EXPECT().Logger(gomock.Any()).Return(logger)
				logger.
					EXPECT().Close()
				logger.
					EXPECT().Errorf(gomock.Any(), gomock.Any())
			},
			inputProc: func(
				logger clogger.Logger,
				fcli cgcp.FirebaseFirestoreClient,
				fauth cgcp.FirebaseAuthClient,
			) error {
				return fmt.Errorf("Dummy error")
			},
			inputOpt: &HO{
				FirestoreClientNotUse: true,
				AuthClientNotUse:      true,
				LoggerNotUse:          false,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrlApp := gomock.NewController(t)
			defer ctrlApp.Finish()
			app := application.NewMockApplicationLogger(ctrlApp)
			ctrlLogger := gomock.NewController(t)
			defer ctrlLogger.Finish()
			logger := clogger.NewMockLogger(ctrlLogger)
			ctrlAppFirebase := gomock.NewController(t)
			defer ctrlAppFirebase.Finish()
			appFirebase := cgcp.NewMockFirebaseApp(ctrlAppFirebase)
			ctrlFirestoreFirebase := gomock.NewController(t)
			defer ctrlFirestoreFirebase.Finish()
			cliFirestoreFirebase := cgcp.NewMockFirebaseFirestoreClient(ctrlFirestoreFirebase)
			ctrlAuthFirebase := gomock.NewController(t)
			defer ctrlAuthFirebase.Finish()
			cliAuthFirebase := cgcp.NewMockFirebaseAuthClient(ctrlAuthFirebase)
			tC.setUpMock(
				app,
				logger,
				appFirebase,
				cliFirestoreFirebase,
				cliAuthFirebase,
			)
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
			result := rec.Result()
			if result.StatusCode != http.StatusOK {
				assert.NotEqual(
					t,
					nil,
					tC.expectedResponse,
				)
				assert.Equal(
					t,
					tC.expectedResponse.StatusCode,
					result.StatusCode,
				)
			}
		})
	}
}
