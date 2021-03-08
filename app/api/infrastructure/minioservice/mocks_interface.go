// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package minioservice is a generated GoMock package.
package minioservice

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMinioService is a mock of MinioService interface.
type MockMinioService struct {
	ctrl     *gomock.Controller
	recorder *MockMinioServiceMockRecorder
}

// MockMinioServiceMockRecorder is the mock recorder for MockMinioService.
type MockMinioServiceMockRecorder struct {
	mock *MockMinioService
}

// NewMockMinioService creates a new mock instance.
func NewMockMinioService(ctrl *gomock.Controller) *MockMinioService {
	mock := &MockMinioService{ctrl: ctrl}
	mock.recorder = &MockMinioServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMinioService) EXPECT() *MockMinioServiceMockRecorder {
	return m.recorder
}

// CreateBucket mocks base method.
func (m *MockMinioService) CreateBucket(bucketName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket", bucketName)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBucket indicates an expected call of CreateBucket.
func (mr *MockMinioServiceMockRecorder) CreateBucket(bucketName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockMinioService)(nil).CreateBucket), bucketName)
}
