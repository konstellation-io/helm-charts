// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package project is a generated GoMock package.
package project

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/konstellation-io/kdl-server/app/api/entity"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockRepository) Get(ctx context.Context, id string) (entity.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(entity.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), ctx, id)
}

// Create mocks base method
func (m *MockRepository) Create(ctx context.Context, project entity.Project) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, project)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockRepositoryMockRecorder) Create(ctx, project interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), ctx, project)
}

// FindByUserID mocks base method
func (m *MockRepository) FindByUserID(ctx context.Context, userID string) ([]entity.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserID", ctx, userID)
	ret0, _ := ret[0].([]entity.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserID indicates an expected call of FindByUserID
func (mr *MockRepositoryMockRecorder) FindByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserID", reflect.TypeOf((*MockRepository)(nil).FindByUserID), ctx, userID)
}

// AddMembers mocks base method
func (m *MockRepository) AddMembers(ctx context.Context, projectID string, members []entity.Member) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMembers", ctx, projectID, members)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMembers indicates an expected call of AddMembers
func (mr *MockRepositoryMockRecorder) AddMembers(ctx, projectID, members interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMembers", reflect.TypeOf((*MockRepository)(nil).AddMembers), ctx, projectID, members)
}

// RemoveMember mocks base method
func (m *MockRepository) RemoveMember(ctx context.Context, projectID, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMember", ctx, projectID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveMember indicates an expected call of RemoveMember
func (mr *MockRepositoryMockRecorder) RemoveMember(ctx, projectID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMember", reflect.TypeOf((*MockRepository)(nil).RemoveMember), ctx, projectID, userID)
}

// UpdateMemberAccessLevel mocks base method
func (m *MockRepository) UpdateMemberAccessLevel(ctx context.Context, projectID, userID string, accessLevel entity.AccessLevel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMemberAccessLevel", ctx, projectID, userID, accessLevel)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMemberAccessLevel indicates an expected call of UpdateMemberAccessLevel
func (mr *MockRepositoryMockRecorder) UpdateMemberAccessLevel(ctx, projectID, userID, accessLevel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMemberAccessLevel", reflect.TypeOf((*MockRepository)(nil).UpdateMemberAccessLevel), ctx, projectID, userID, accessLevel)
}

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

// Create mocks base method
func (m *MockUseCase) Create(ctx context.Context, opt CreateProjectOption) (entity.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, opt)
	ret0, _ := ret[0].(entity.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockUseCaseMockRecorder) Create(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUseCase)(nil).Create), ctx, opt)
}

// FindByUserID mocks base method
func (m *MockUseCase) FindByUserID(ctx context.Context, userID string) ([]entity.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserID", ctx, userID)
	ret0, _ := ret[0].([]entity.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserID indicates an expected call of FindByUserID
func (mr *MockUseCaseMockRecorder) FindByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserID", reflect.TypeOf((*MockUseCase)(nil).FindByUserID), ctx, userID)
}

// GetByID mocks base method
func (m *MockUseCase) GetByID(ctx context.Context, id string) (entity.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(entity.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockUseCaseMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUseCase)(nil).GetByID), ctx, id)
}

// AddMembers mocks base method
func (m *MockUseCase) AddMembers(ctx context.Context, opt AddMembersOption) (entity.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMembers", ctx, opt)
	ret0, _ := ret[0].(entity.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMembers indicates an expected call of AddMembers
func (mr *MockUseCaseMockRecorder) AddMembers(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMembers", reflect.TypeOf((*MockUseCase)(nil).AddMembers), ctx, opt)
}

// RemoveMember mocks base method
func (m *MockUseCase) RemoveMember(ctx context.Context, opt RemoveMemberOption) (entity.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMember", ctx, opt)
	ret0, _ := ret[0].(entity.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveMember indicates an expected call of RemoveMember
func (mr *MockUseCaseMockRecorder) RemoveMember(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMember", reflect.TypeOf((*MockUseCase)(nil).RemoveMember), ctx, opt)
}

// UpdateMember mocks base method
func (m *MockUseCase) UpdateMember(ctx context.Context, opt UpdateMemberOption) (entity.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMember", ctx, opt)
	ret0, _ := ret[0].(entity.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMember indicates an expected call of UpdateMember
func (mr *MockUseCaseMockRecorder) UpdateMember(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMember", reflect.TypeOf((*MockUseCase)(nil).UpdateMember), ctx, opt)
}
