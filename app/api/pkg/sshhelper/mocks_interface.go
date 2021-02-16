// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package sshhelper is a generated GoMock package.
package sshhelper

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/konstellation-io/kdl-server/app/api/entity"
	reflect "reflect"
)

// MockSSHKeyGenerator is a mock of SSHKeyGenerator interface
type MockSSHKeyGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockSSHKeyGeneratorMockRecorder
}

// MockSSHKeyGeneratorMockRecorder is the mock recorder for MockSSHKeyGenerator
type MockSSHKeyGeneratorMockRecorder struct {
	mock *MockSSHKeyGenerator
}

// NewMockSSHKeyGenerator creates a new mock instance
func NewMockSSHKeyGenerator(ctrl *gomock.Controller) *MockSSHKeyGenerator {
	mock := &MockSSHKeyGenerator{ctrl: ctrl}
	mock.recorder = &MockSSHKeyGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSSHKeyGenerator) EXPECT() *MockSSHKeyGeneratorMockRecorder {
	return m.recorder
}

// NewKeys mocks base method
func (m *MockSSHKeyGenerator) NewKeys() (entity.SSHKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewKeys")
	ret0, _ := ret[0].(entity.SSHKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewKeys indicates an expected call of NewKeys
func (mr *MockSSHKeyGeneratorMockRecorder) NewKeys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewKeys", reflect.TypeOf((*MockSSHKeyGenerator)(nil).NewKeys))
}