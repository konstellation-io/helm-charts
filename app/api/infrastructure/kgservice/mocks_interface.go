// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package kgservice is a generated GoMock package.
package kgservice

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/konstellation-io/kdl-server/app/api/entity"
	reflect "reflect"
)

// MockKGService is a mock of KGService interface
type MockKGService struct {
	ctrl     *gomock.Controller
	recorder *MockKGServiceMockRecorder
}

// MockKGServiceMockRecorder is the mock recorder for MockKGService
type MockKGServiceMockRecorder struct {
	mock *MockKGService
}

// NewMockKGService creates a new mock instance
func NewMockKGService(ctrl *gomock.Controller) *MockKGService {
	mock := &MockKGService{ctrl: ctrl}
	mock.recorder = &MockKGServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKGService) EXPECT() *MockKGServiceMockRecorder {
	return m.recorder
}

// GetGraph mocks base method
func (m *MockKGService) GetGraph(ctx context.Context, description string) (entity.KnowledgeGraph, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGraph", ctx, description)
	ret0, _ := ret[0].(entity.KnowledgeGraph)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGraph indicates an expected call of GetGraph
func (mr *MockKGServiceMockRecorder) GetGraph(ctx, description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGraph", reflect.TypeOf((*MockKGService)(nil).GetGraph), ctx, description)
}
