// Code generated by MockGen. DO NOT EDIT.
// Source: repository/issue_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	request "go-sandbox/libs/dto/request"
	response "go-sandbox/libs/dto/response"
	model "go-sandbox/libs/model"
	gomock "github.com/golang/mock/gomock"
)

// MockIssueRepository is a mock of IssueRepository interface.
type MockIssueRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIssueRepositoryMockRecorder
}

// MockIssueRepositoryMockRecorder is the mock recorder for MockIssueRepository.
type MockIssueRepositoryMockRecorder struct {
	mock *MockIssueRepository
}

// NewMockIssueRepository creates a new mock instance.
func NewMockIssueRepository(ctrl *gomock.Controller) *MockIssueRepository {
	mock := &MockIssueRepository{ctrl: ctrl}
	mock.recorder = &MockIssueRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIssueRepository) EXPECT() *MockIssueRepositoryMockRecorder {
	return m.recorder
}

// AssignieIssueToUser mocks base method.
func (m *MockIssueRepository) AssignieIssueToUser(issue model.Issue, user model.User) (*model.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignieIssueToUser", issue, user)
	ret0, _ := ret[0].(*model.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignieIssueToUser indicates an expected call of AssignieIssueToUser.
func (mr *MockIssueRepositoryMockRecorder) AssignieIssueToUser(issue, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignieIssueToUser", reflect.TypeOf((*MockIssueRepository)(nil).AssignieIssueToUser), issue, user)
}

// FindIssue mocks base method.
func (m *MockIssueRepository) FindIssue(issue_id uint64) (*model.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindIssue", issue_id)
	ret0, _ := ret[0].(*model.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindIssue indicates an expected call of FindIssue.
func (mr *MockIssueRepositoryMockRecorder) FindIssue(issue_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindIssue", reflect.TypeOf((*MockIssueRepository)(nil).FindIssue), issue_id)
}

// FindIssueByAccess mocks base method.
func (m *MockIssueRepository) FindIssueByAccess(issue_id, user_id uint64) (*model.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindIssueByAccess", issue_id, user_id)
	ret0, _ := ret[0].(*model.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindIssueByAccess indicates an expected call of FindIssueByAccess.
func (mr *MockIssueRepositoryMockRecorder) FindIssueByAccess(issue_id, user_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindIssueByAccess", reflect.TypeOf((*MockIssueRepository)(nil).FindIssueByAccess), issue_id, user_id)
}

// GetIssues mocks base method.
func (m *MockIssueRepository) GetIssues(issueGetQuery *request.IssueGetQuery, userID uint64) ([]response.IssueResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssues", issueGetQuery, userID)
	ret0, _ := ret[0].([]response.IssueResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssues indicates an expected call of GetIssues.
func (mr *MockIssueRepositoryMockRecorder) GetIssues(issueGetQuery, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssues", reflect.TypeOf((*MockIssueRepository)(nil).GetIssues), issueGetQuery, userID)
}

// InsertDependentIssueAssociation mocks base method.
func (m *MockIssueRepository) InsertDependentIssueAssociation(issue, dependentIssue model.Issue) (*model.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertDependentIssueAssociation", issue, dependentIssue)
	ret0, _ := ret[0].(*model.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertDependentIssueAssociation indicates an expected call of InsertDependentIssueAssociation.
func (mr *MockIssueRepositoryMockRecorder) InsertDependentIssueAssociation(issue, dependentIssue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertDependentIssueAssociation", reflect.TypeOf((*MockIssueRepository)(nil).InsertDependentIssueAssociation), issue, dependentIssue)
}

// InsertIssue mocks base method.
func (m *MockIssueRepository) InsertIssue(issue model.Issue) (*model.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertIssue", issue)
	ret0, _ := ret[0].(*model.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertIssue indicates an expected call of InsertIssue.
func (mr *MockIssueRepositoryMockRecorder) InsertIssue(issue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertIssue", reflect.TypeOf((*MockIssueRepository)(nil).InsertIssue), issue)
}
