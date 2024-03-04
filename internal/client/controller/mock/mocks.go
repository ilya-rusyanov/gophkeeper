// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ilya-rusyanov/gophkeeper/internal/client/controller (interfaces: Registerer,Storer,BinStorer,LogIner,Lister,Shower)
//
// Generated by this command:
//
//	mockgen -destination ./mock/mocks.go -package mock . Registerer,Storer,BinStorer,LogIner,Lister,Shower
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	entity "github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockRegisterer is a mock of Registerer interface.
type MockRegisterer struct {
	ctrl     *gomock.Controller
	recorder *MockRegistererMockRecorder
}

// MockRegistererMockRecorder is the mock recorder for MockRegisterer.
type MockRegistererMockRecorder struct {
	mock *MockRegisterer
}

// NewMockRegisterer creates a new mock instance.
func NewMockRegisterer(ctrl *gomock.Controller) *MockRegisterer {
	mock := &MockRegisterer{ctrl: ctrl}
	mock.recorder = &MockRegistererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegisterer) EXPECT() *MockRegistererMockRecorder {
	return m.recorder
}

// Register mocks base method.
func (m *MockRegisterer) Register(arg0 context.Context, arg1 entity.MyCredentials) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockRegistererMockRecorder) Register(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRegisterer)(nil).Register), arg0, arg1)
}

// MockStorer is a mock of Storer interface.
type MockStorer struct {
	ctrl     *gomock.Controller
	recorder *MockStorerMockRecorder
}

// MockStorerMockRecorder is the mock recorder for MockStorer.
type MockStorerMockRecorder struct {
	mock *MockStorer
}

// NewMockStorer creates a new mock instance.
func NewMockStorer(ctrl *gomock.Controller) *MockStorer {
	mock := &MockStorer{ctrl: ctrl}
	mock.recorder = &MockStorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorer) EXPECT() *MockStorerMockRecorder {
	return m.recorder
}

// Store mocks base method.
func (m *MockStorer) Store(arg0 context.Context, arg1 entity.Record) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockStorerMockRecorder) Store(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockStorer)(nil).Store), arg0, arg1)
}

// MockBinStorer is a mock of BinStorer interface.
type MockBinStorer struct {
	ctrl     *gomock.Controller
	recorder *MockBinStorerMockRecorder
}

// MockBinStorerMockRecorder is the mock recorder for MockBinStorer.
type MockBinStorerMockRecorder struct {
	mock *MockBinStorer
}

// NewMockBinStorer creates a new mock instance.
func NewMockBinStorer(ctrl *gomock.Controller) *MockBinStorer {
	mock := &MockBinStorer{ctrl: ctrl}
	mock.recorder = &MockBinStorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBinStorer) EXPECT() *MockBinStorerMockRecorder {
	return m.recorder
}

// StoreBin mocks base method.
func (m *MockBinStorer) StoreBin(arg0 context.Context, arg1 entity.Record, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreBin", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreBin indicates an expected call of StoreBin.
func (mr *MockBinStorerMockRecorder) StoreBin(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreBin", reflect.TypeOf((*MockBinStorer)(nil).StoreBin), arg0, arg1, arg2)
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
func (m *MockLogIner) LogIn(arg0 context.Context, arg1 entity.MyCredentials) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogIn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// LogIn indicates an expected call of LogIn.
func (mr *MockLogInerMockRecorder) LogIn(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogIn", reflect.TypeOf((*MockLogIner)(nil).LogIn), arg0, arg1)
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
func (m *MockLister) List(arg0 context.Context) (entity.DataList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(entity.DataList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockListerMockRecorder) List(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLister)(nil).List), arg0)
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
func (m *MockShower) Show(arg0 context.Context, arg1 entity.ShowIn) (entity.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show", arg0, arg1)
	ret0, _ := ret[0].(entity.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Show indicates an expected call of Show.
func (mr *MockShowerMockRecorder) Show(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockShower)(nil).Show), arg0, arg1)
}
