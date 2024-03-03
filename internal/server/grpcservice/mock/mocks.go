// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ilya-rusyanov/gophkeeper/internal/server/grpcservice (interfaces: IRegisterUC,LogIner,IStoreUC,Lister,Shower)
//
// Generated by this command:
//
//	mockgen -destination ./mock/mocks.go -package mock . IRegisterUC,LogIner,IStoreUC,Lister,Shower
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	entity "github.com/ilya-rusyanov/gophkeeper/internal/server/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockIRegisterUC is a mock of IRegisterUC interface.
type MockIRegisterUC struct {
	ctrl     *gomock.Controller
	recorder *MockIRegisterUCMockRecorder
}

// MockIRegisterUCMockRecorder is the mock recorder for MockIRegisterUC.
type MockIRegisterUCMockRecorder struct {
	mock *MockIRegisterUC
}

// NewMockIRegisterUC creates a new mock instance.
func NewMockIRegisterUC(ctrl *gomock.Controller) *MockIRegisterUC {
	mock := &MockIRegisterUC{ctrl: ctrl}
	mock.recorder = &MockIRegisterUCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRegisterUC) EXPECT() *MockIRegisterUCMockRecorder {
	return m.recorder
}

// Register mocks base method.
func (m *MockIRegisterUC) Register(arg0 context.Context, arg1 entity.UserCredentials) (entity.AuthToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(entity.AuthToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockIRegisterUCMockRecorder) Register(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockIRegisterUC)(nil).Register), arg0, arg1)
}

// MockLogIner is a mock of LogIner interface.
type MockLogIner struct {
	ctrl     *gomock.Controller
	recorder *MockLogInerMockRecorder
}

// MockLogInerMockRecorder is the mock recorder for MockLogIner.
type MockLogInerMockRecorder struct {
	mock *MockLogIner
}

// NewMockLogIner creates a new mock instance.
func NewMockLogIner(ctrl *gomock.Controller) *MockLogIner {
	mock := &MockLogIner{ctrl: ctrl}
	mock.recorder = &MockLogInerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogIner) EXPECT() *MockLogInerMockRecorder {
	return m.recorder
}

// LogIn mocks base method.
func (m *MockLogIner) LogIn(arg0 context.Context, arg1 entity.UserCredentials) (entity.AuthToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogIn", arg0, arg1)
	ret0, _ := ret[0].(entity.AuthToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogIn indicates an expected call of LogIn.
func (mr *MockLogInerMockRecorder) LogIn(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogIn", reflect.TypeOf((*MockLogIner)(nil).LogIn), arg0, arg1)
}

// MockIStoreUC is a mock of IStoreUC interface.
type MockIStoreUC struct {
	ctrl     *gomock.Controller
	recorder *MockIStoreUCMockRecorder
}

// MockIStoreUCMockRecorder is the mock recorder for MockIStoreUC.
type MockIStoreUCMockRecorder struct {
	mock *MockIStoreUC
}

// NewMockIStoreUC creates a new mock instance.
func NewMockIStoreUC(ctrl *gomock.Controller) *MockIStoreUC {
	mock := &MockIStoreUC{ctrl: ctrl}
	mock.recorder = &MockIStoreUCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStoreUC) EXPECT() *MockIStoreUCMockRecorder {
	return m.recorder
}

// Store mocks base method.
func (m *MockIStoreUC) Store(arg0 context.Context, arg1 *entity.StoreIn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockIStoreUCMockRecorder) Store(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockIStoreUC)(nil).Store), arg0, arg1)
}

// MockLister is a mock of Lister interface.
type MockLister struct {
	ctrl     *gomock.Controller
	recorder *MockListerMockRecorder
}

// MockListerMockRecorder is the mock recorder for MockLister.
type MockListerMockRecorder struct {
	mock *MockLister
}

// NewMockLister creates a new mock instance.
func NewMockLister(ctrl *gomock.Controller) *MockLister {
	mock := &MockLister{ctrl: ctrl}
	mock.recorder = &MockListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLister) EXPECT() *MockListerMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockLister) List(arg0 context.Context, arg1 string) (entity.DataListing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(entity.DataListing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockListerMockRecorder) List(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLister)(nil).List), arg0, arg1)
}

// MockShower is a mock of Shower interface.
type MockShower struct {
	ctrl     *gomock.Controller
	recorder *MockShowerMockRecorder
}

// MockShowerMockRecorder is the mock recorder for MockShower.
type MockShowerMockRecorder struct {
	mock *MockShower
}

// NewMockShower creates a new mock instance.
func NewMockShower(ctrl *gomock.Controller) *MockShower {
	mock := &MockShower{ctrl: ctrl}
	mock.recorder = &MockShowerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShower) EXPECT() *MockShowerMockRecorder {
	return m.recorder
}

// Show mocks base method.
func (m *MockShower) Show(arg0 context.Context, arg1 entity.ShowIn) (entity.ShowResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show", arg0, arg1)
	ret0, _ := ret[0].(entity.ShowResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Show indicates an expected call of Show.
func (mr *MockShowerMockRecorder) Show(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockShower)(nil).Show), arg0, arg1)
}
