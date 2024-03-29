// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ilya-rusyanov/gophkeeper/internal/client/usecase/show (interfaces: Servicer,Storager,FileSaver)
//
// Generated by this command:
//
//	mockgen -destination ./mock/mocks.go -package mock . Servicer,Storager,FileSaver
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	entity "github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockServicer is a mock of Servicer interface.
type MockServicer struct {
	ctrl     *gomock.Controller
	recorder *MockServicerMockRecorder
}

// MockServicerMockRecorder is the mock recorder for MockServicer.
type MockServicerMockRecorder struct {
	mock *MockServicer
}

// NewMockServicer creates a new mock instance.
func NewMockServicer(ctrl *gomock.Controller) *MockServicer {
	mock := &MockServicer{ctrl: ctrl}
	mock.recorder = &MockServicerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicer) EXPECT() *MockServicerMockRecorder {
	return m.recorder
}

// Show mocks base method.
func (m *MockServicer) Show(arg0 context.Context, arg1 entity.ServiceShowRequest) (entity.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show", arg0, arg1)
	ret0, _ := ret[0].(entity.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Show indicates an expected call of Show.
func (mr *MockServicerMockRecorder) Show(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockServicer)(nil).Show), arg0, arg1)
}

// MockStorager is a mock of Storager interface.
type MockStorager struct {
	ctrl     *gomock.Controller
	recorder *MockStoragerMockRecorder
}

// MockStoragerMockRecorder is the mock recorder for MockStorager.
type MockStoragerMockRecorder struct {
	mock *MockStorager
}

// NewMockStorager creates a new mock instance.
func NewMockStorager(ctrl *gomock.Controller) *MockStorager {
	mock := &MockStorager{ctrl: ctrl}
	mock.recorder = &MockStoragerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorager) EXPECT() *MockStoragerMockRecorder {
	return m.recorder
}

// Load mocks base method.
func (m *MockStorager) Load() (entity.MyAuthentication, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load")
	ret0, _ := ret[0].(entity.MyAuthentication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Load indicates an expected call of Load.
func (mr *MockStoragerMockRecorder) Load() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockStorager)(nil).Load))
}

// MockFileSaver is a mock of FileSaver interface.
type MockFileSaver struct {
	ctrl     *gomock.Controller
	recorder *MockFileSaverMockRecorder
}

// MockFileSaverMockRecorder is the mock recorder for MockFileSaver.
type MockFileSaverMockRecorder struct {
	mock *MockFileSaver
}

// NewMockFileSaver creates a new mock instance.
func NewMockFileSaver(ctrl *gomock.Controller) *MockFileSaver {
	mock := &MockFileSaver{ctrl: ctrl}
	mock.recorder = &MockFileSaverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileSaver) EXPECT() *MockFileSaverMockRecorder {
	return m.recorder
}

// SaveFile mocks base method.
func (m *MockFileSaver) SaveFile(arg0 context.Context, arg1 entity.FileSaveIn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveFile indicates an expected call of SaveFile.
func (mr *MockFileSaverMockRecorder) SaveFile(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveFile", reflect.TypeOf((*MockFileSaver)(nil).SaveFile), arg0, arg1)
}
