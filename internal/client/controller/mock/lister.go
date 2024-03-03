// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ilya-rusyanov/gophkeeper/internal/client/controller (interfaces: Lister)
//
// Generated by this command:
//
//	mockgen -destination ./mock/lister.go -package mock . Lister
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

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
func (m *MockLister) List(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockListerMockRecorder) List(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLister)(nil).List), arg0)
}
