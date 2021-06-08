// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package k8s is a generated GoMock package.
package k8s

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/konstellation-io/kdl-server/app/api/entity"
)

// MockK8sClient is a mock of K8sClient interface.
type MockK8sClient struct {
	ctrl     *gomock.Controller
	recorder *MockK8sClientMockRecorder
}

// MockK8sClientMockRecorder is the mock recorder for MockK8sClient.
type MockK8sClientMockRecorder struct {
	mock *MockK8sClient
}

// NewMockK8sClient creates a new mock instance.
func NewMockK8sClient(ctrl *gomock.Controller) *MockK8sClient {
	mock := &MockK8sClient{ctrl: ctrl}
	mock.recorder = &MockK8sClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockK8sClient) EXPECT() *MockK8sClientMockRecorder {
	return m.recorder
}

// CreateKDLProjectCR mocks base method.
func (m *MockK8sClient) CreateKDLProjectCR(ctx context.Context, projectID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKDLProjectCR", ctx, projectID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateKDLProjectCR indicates an expected call of CreateKDLProjectCR.
func (mr *MockK8sClientMockRecorder) CreateKDLProjectCR(ctx, projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKDLProjectCR", reflect.TypeOf((*MockK8sClient)(nil).CreateKDLProjectCR), ctx, projectID)
}

// CreateSecret mocks base method.
func (m *MockK8sClient) CreateSecret(ctx context.Context, name string, values map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecret", ctx, name, values)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSecret indicates an expected call of CreateSecret.
func (mr *MockK8sClientMockRecorder) CreateSecret(ctx, name, values interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecret", reflect.TypeOf((*MockK8sClient)(nil).CreateSecret), ctx, name, values)
}

// CreateUserSSHKeySecret mocks base method.
func (m *MockK8sClient) CreateUserSSHKeySecret(ctx context.Context, user entity.User, public, private string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserSSHKeySecret", ctx, user, public, private)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUserSSHKeySecret indicates an expected call of CreateUserSSHKeySecret.
func (mr *MockK8sClientMockRecorder) CreateUserSSHKeySecret(ctx, user, public, private interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserSSHKeySecret", reflect.TypeOf((*MockK8sClient)(nil).CreateUserSSHKeySecret), ctx, user, public, private)
}

// CreateUserToolsCR mocks base method.
func (m *MockK8sClient) CreateUserToolsCR(ctx context.Context, username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserToolsCR", ctx, username)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUserToolsCR indicates an expected call of CreateUserToolsCR.
func (mr *MockK8sClientMockRecorder) CreateUserToolsCR(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserToolsCR", reflect.TypeOf((*MockK8sClient)(nil).CreateUserToolsCR), ctx, username)
}

// DeleteUserToolsCR mocks base method.
func (m *MockK8sClient) DeleteUserToolsCR(ctx context.Context, username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserToolsCR", ctx, username)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserToolsCR indicates an expected call of DeleteUserToolsCR.
func (mr *MockK8sClientMockRecorder) DeleteUserToolsCR(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserToolsCR", reflect.TypeOf((*MockK8sClient)(nil).DeleteUserToolsCR), ctx, username)
}

// GetSecret mocks base method.
func (m *MockK8sClient) GetSecret(ctx context.Context, name string) (map[string][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", ctx, name)
	ret0, _ := ret[0].(map[string][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret.
func (mr *MockK8sClientMockRecorder) GetSecret(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockK8sClient)(nil).GetSecret), ctx, name)
}

// GetUserSSHKeyPublic mocks base method.
func (m *MockK8sClient) GetUserSSHKeyPublic(ctx context.Context, usernameSlug string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSSHKeyPublic", ctx, usernameSlug)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSSHKeyPublic indicates an expected call of GetUserSSHKeyPublic.
func (mr *MockK8sClientMockRecorder) GetUserSSHKeyPublic(ctx, usernameSlug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSSHKeyPublic", reflect.TypeOf((*MockK8sClient)(nil).GetUserSSHKeyPublic), ctx, usernameSlug)
}

// GetUserSSHKeySecret mocks base method.
func (m *MockK8sClient) GetUserSSHKeySecret(ctx context.Context, usernameSlug string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSSHKeySecret", ctx, usernameSlug)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSSHKeySecret indicates an expected call of GetUserSSHKeySecret.
func (mr *MockK8sClientMockRecorder) GetUserSSHKeySecret(ctx, usernameSlug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSSHKeySecret", reflect.TypeOf((*MockK8sClient)(nil).GetUserSSHKeySecret), ctx, usernameSlug)
}

// IsUserToolPODRunning mocks base method.
func (m *MockK8sClient) IsUserToolPODRunning(ctx context.Context, username string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserToolPODRunning", ctx, username)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUserToolPODRunning indicates an expected call of IsUserToolPODRunning.
func (mr *MockK8sClientMockRecorder) IsUserToolPODRunning(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserToolPODRunning", reflect.TypeOf((*MockK8sClient)(nil).IsUserToolPODRunning), ctx, username)
}

// UpdateSecret mocks base method.
func (m *MockK8sClient) UpdateSecret(ctx context.Context, name string, values map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecret", ctx, name, values)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecret indicates an expected call of UpdateSecret.
func (mr *MockK8sClientMockRecorder) UpdateSecret(ctx, name, values interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecret", reflect.TypeOf((*MockK8sClient)(nil).UpdateSecret), ctx, name, values)
}

// UpdateUserSSHKeySecret mocks base method.
func (m *MockK8sClient) UpdateUserSSHKeySecret(ctx context.Context, user entity.User, public, private string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserSSHKeySecret", ctx, user, public, private)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserSSHKeySecret indicates an expected call of UpdateUserSSHKeySecret.
func (mr *MockK8sClientMockRecorder) UpdateUserSSHKeySecret(ctx, user, public, private interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserSSHKeySecret", reflect.TypeOf((*MockK8sClient)(nil).UpdateUserSSHKeySecret), ctx, user, public, private)
}
