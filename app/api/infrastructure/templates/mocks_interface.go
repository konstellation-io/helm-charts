// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package templates is a generated GoMock package.
package templates

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/konstellation-io/kdl-server/app/api/entity"
)

// MockTemplating is a mock of Templating interface.
type MockTemplating struct {
	ctrl     *gomock.Controller
	recorder *MockTemplatingMockRecorder
}

// MockTemplatingMockRecorder is the mock recorder for MockTemplating.
type MockTemplatingMockRecorder struct {
	mock *MockTemplating
}

// NewMockTemplating creates a new mock instance.
func NewMockTemplating(ctrl *gomock.Controller) *MockTemplating {
	mock := &MockTemplating{ctrl: ctrl}
	mock.recorder = &MockTemplatingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTemplating) EXPECT() *MockTemplatingMockRecorder {
	return m.recorder
}

// GenerateInitialProjectContent mocks base method.
func (m *MockTemplating) GenerateInitialProjectContent(ctx context.Context, project entity.Project, user entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateInitialProjectContent", ctx, project, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateInitialProjectContent indicates an expected call of GenerateInitialProjectContent.
func (mr *MockTemplatingMockRecorder) GenerateInitialProjectContent(ctx, project, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateInitialProjectContent", reflect.TypeOf((*MockTemplating)(nil).GenerateInitialProjectContent), ctx, project, user)
}