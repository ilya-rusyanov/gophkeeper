// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ilya-rusyanov/gophkeeper/internal/client/usecase/register (interfaces: Storager)
//
// Generated by this command:
//
//	mockgen -package mock -destination ./mock/storager.go . Storager
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	entity "github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
	gomock "go.uber.org/mock/gomock"
)

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

// Store mocks base method.
func (m *MockStorager) Store(arg0 context.Context, arg1 entity.MyAuthentication) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockStoragerMockRecorder) Store(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockStorager)(nil).Store), arg0, arg1)
}
