// Code generated by MockGen. DO NOT EDIT.
// Source: cgcp/wrapper.go

// Package cgcp is a generated GoMock package.
package cgcp

import (
	context "context"
	auth "firebase.google.com/go/auth"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFirebaseApp is a mock of FirebaseApp interface
type MockFirebaseApp struct {
	ctrl     *gomock.Controller
	recorder *MockFirebaseAppMockRecorder
}

// MockFirebaseAppMockRecorder is the mock recorder for MockFirebaseApp
type MockFirebaseAppMockRecorder struct {
	mock *MockFirebaseApp
}

// NewMockFirebaseApp creates a new mock instance
func NewMockFirebaseApp(ctrl *gomock.Controller) *MockFirebaseApp {
	mock := &MockFirebaseApp{ctrl: ctrl}
	mock.recorder = &MockFirebaseAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFirebaseApp) EXPECT() *MockFirebaseAppMockRecorder {
	return m.recorder
}

// Firestore mocks base method
func (m *MockFirebaseApp) Firestore(ctx context.Context) (FirebaseFirestoreClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Firestore", ctx)
	ret0, _ := ret[0].(FirebaseFirestoreClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Firestore indicates an expected call of Firestore
func (mr *MockFirebaseAppMockRecorder) Firestore(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Firestore", reflect.TypeOf((*MockFirebaseApp)(nil).Firestore), ctx)
}

// Auth mocks base method
func (m *MockFirebaseApp) Auth(ctx context.Context) (FirebaseAuthClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auth", ctx)
	ret0, _ := ret[0].(FirebaseAuthClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Auth indicates an expected call of Auth
func (mr *MockFirebaseAppMockRecorder) Auth(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auth", reflect.TypeOf((*MockFirebaseApp)(nil).Auth), ctx)
}

// MockFirebaseFirestoreClient is a mock of FirebaseFirestoreClient interface
type MockFirebaseFirestoreClient struct {
	ctrl     *gomock.Controller
	recorder *MockFirebaseFirestoreClientMockRecorder
}

// MockFirebaseFirestoreClientMockRecorder is the mock recorder for MockFirebaseFirestoreClient
type MockFirebaseFirestoreClientMockRecorder struct {
	mock *MockFirebaseFirestoreClient
}

// NewMockFirebaseFirestoreClient creates a new mock instance
func NewMockFirebaseFirestoreClient(ctrl *gomock.Controller) *MockFirebaseFirestoreClient {
	mock := &MockFirebaseFirestoreClient{ctrl: ctrl}
	mock.recorder = &MockFirebaseFirestoreClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFirebaseFirestoreClient) EXPECT() *MockFirebaseFirestoreClientMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockFirebaseFirestoreClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockFirebaseFirestoreClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockFirebaseFirestoreClient)(nil).Close))
}

// MockFirebaseAuthClient is a mock of FirebaseAuthClient interface
type MockFirebaseAuthClient struct {
	ctrl     *gomock.Controller
	recorder *MockFirebaseAuthClientMockRecorder
}

// MockFirebaseAuthClientMockRecorder is the mock recorder for MockFirebaseAuthClient
type MockFirebaseAuthClientMockRecorder struct {
	mock *MockFirebaseAuthClient
}

// NewMockFirebaseAuthClient creates a new mock instance
func NewMockFirebaseAuthClient(ctrl *gomock.Controller) *MockFirebaseAuthClient {
	mock := &MockFirebaseAuthClient{ctrl: ctrl}
	mock.recorder = &MockFirebaseAuthClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFirebaseAuthClient) EXPECT() *MockFirebaseAuthClientMockRecorder {
	return m.recorder
}

// VerifyIDToken mocks base method
func (m *MockFirebaseAuthClient) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyIDToken", ctx, idToken)
	ret0, _ := ret[0].(*auth.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyIDToken indicates an expected call of VerifyIDToken
func (mr *MockFirebaseAuthClientMockRecorder) VerifyIDToken(ctx, idToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyIDToken", reflect.TypeOf((*MockFirebaseAuthClient)(nil).VerifyIDToken), ctx, idToken)
}
