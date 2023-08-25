// Code generated by MockGen. DO NOT EDIT.
// Source: broker.go

// Package broker is a generated GoMock package.
package broker

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIBroker is a mock of IBroker interface.
type MockIBroker struct {
	ctrl     *gomock.Controller
	recorder *MockIBrokerMockRecorder
}

// MockIBrokerMockRecorder is the mock recorder for MockIBroker.
type MockIBrokerMockRecorder struct {
	mock *MockIBroker
}

// NewMockIBroker creates a new mock instance.
func NewMockIBroker(ctrl *gomock.Controller) *MockIBroker {
	mock := &MockIBroker{ctrl: ctrl}
	mock.recorder = &MockIBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBroker) EXPECT() *MockIBrokerMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *MockIBroker) Publish(ctx context.Context, message *BrokerMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockIBrokerMockRecorder) Publish(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockIBroker)(nil).Publish), ctx, message)
}
