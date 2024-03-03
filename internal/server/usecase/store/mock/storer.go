// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ilya-rusyanov/gophkeeper/internal/server/usecase/store (interfaces: Storer)
//
// Generated by this command:
//
//	mockgen -destination ./mock/storer.go -package mock . Storer
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	entity "github.com/ilya-rusyanov/gophkeeper/internal/server/entity"
	gomock "go.uber.org/mock/gomock"
)

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
func (m *MockStorer) Store(arg0 context.Context, arg1 *entity.StoreIn) error {
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
