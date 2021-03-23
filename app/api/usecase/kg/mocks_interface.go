// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package kg is a generated GoMock package.
package kg

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/konstellation-io/kdl-server/app/api/entity"
	reflect "reflect"
)

// MockUseCase is a mock of UseCase interface
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockUseCase) Get(ctx context.Context, description string) (entity.KnowledgeGraph, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, description)
	ret0, _ := ret[0].(entity.KnowledgeGraph)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockUseCaseMockRecorder) Get(ctx, description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUseCase)(nil).Get), ctx, description)
}

// GetItem mocks base method
func (m *MockUseCase) GetItem(ctx context.Context, id string) (entity.KnowledgeGraphItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItem", ctx, id)
	ret0, _ := ret[0].(entity.KnowledgeGraphItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItem indicates an expected call of GetItem
func (mr *MockUseCaseMockRecorder) GetItem(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItem", reflect.TypeOf((*MockUseCase)(nil).GetItem), ctx, id)
}
