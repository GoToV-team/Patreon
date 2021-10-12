// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_creator is a generated GoMock package.
package mock_repository

import (
	"patreon/internal/app/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(arg0 *models.Creator) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), arg0)
}

// GetCreator mocks base method.
func (m *MockRepository) GetCreator(arg0 int64) (*models.Creator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreator", arg0)
	ret0, _ := ret[0].(*models.Creator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCreator indicates an expected call of GetCreator.
func (mr *MockRepositoryMockRecorder) GetCreator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreator", reflect.TypeOf((*MockRepository)(nil).GetCreator), arg0)
}

// GetCreators mocks base method.
func (m *MockRepository) GetCreators() ([]models.Creator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreators")
	ret0, _ := ret[0].([]models.Creator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCreators indicates an expected call of GetCreators.
func (mr *MockRepositoryMockRecorder) GetCreators() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreators", reflect.TypeOf((*MockRepository)(nil).GetCreators))
}
