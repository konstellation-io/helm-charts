// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package k8s is a generated GoMock package.
package k8s

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockK8sClient is a mock of K8sClient interface
type MockK8sClient struct {
	ctrl     *gomock.Controller
	recorder *MockK8sClientMockRecorder
}

// MockK8sClientMockRecorder is the mock recorder for MockK8sClient
type MockK8sClientMockRecorder struct {
	mock *MockK8sClient
}

// NewMockK8sClient creates a new mock instance
func NewMockK8sClient(ctrl *gomock.Controller) *MockK8sClient {
	mock := &MockK8sClient{ctrl: ctrl}
	mock.recorder = &MockK8sClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockK8sClient) EXPECT() *MockK8sClientMockRecorder {
	return m.recorder
}

// CreateSecret mocks base method
func (m *MockK8sClient) CreateSecret(name string, values map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecret", name, values)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSecret indicates an expected call of CreateSecret
func (mr *MockK8sClientMockRecorder) CreateSecret(name, values interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecret", reflect.TypeOf((*MockK8sClient)(nil).CreateSecret), name, values)
}

// CreateUserToolsCR mocks base method
func (m *MockK8sClient) CreateUserToolsCR(ctx context.Context, username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserToolsCR", ctx, username)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUserToolsCR indicates an expected call of CreateUserToolsCR
func (mr *MockK8sClientMockRecorder) CreateUserToolsCR(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserToolsCR", reflect.TypeOf((*MockK8sClient)(nil).CreateUserToolsCR), ctx, username)
}

// DeleteUserToolsCR mocks base method
func (m *MockK8sClient) DeleteUserToolsCR(ctx context.Context, username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserToolsCR", ctx, username)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserToolsCR indicates an expected call of DeleteUserToolsCR
func (mr *MockK8sClientMockRecorder) DeleteUserToolsCR(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserToolsCR", reflect.TypeOf((*MockK8sClient)(nil).DeleteUserToolsCR), ctx, username)
}

// IsUserToolPODRunning mocks base method
func (m *MockK8sClient) IsUserToolPODRunning(username string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserToolPODRunning", username)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUserToolPODRunning indicates an expected call of IsUserToolPODRunning
func (mr *MockK8sClientMockRecorder) IsUserToolPODRunning(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserToolPODRunning", reflect.TypeOf((*MockK8sClient)(nil).IsUserToolPODRunning), username)
}