// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ilya-rusyanov/gophkeeper/internal/client/controller (interfaces: LogIner)
//
// Generated by this command:
//
//	mockgen -destination ./mock/loginer.go -package mock . LogIner
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	entity "github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
	gomock "go.uber.org/mock/gomock"
)

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
