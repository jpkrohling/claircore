// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quay/claircore/internal/vulnstore (interfaces: Updater)

// Package vulnstore is a generated GoMock package.
package vulnstore

import (
	gomock "github.com/golang/mock/gomock"
	claircore "github.com/quay/claircore"
	reflect "reflect"
)

// MockUpdater is a mock of Updater interface
type MockUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockUpdaterMockRecorder
}

// MockUpdaterMockRecorder is the mock recorder for MockUpdater
type MockUpdaterMockRecorder struct {
	mock *MockUpdater
}

// NewMockUpdater creates a new mock instance
func NewMockUpdater(ctrl *gomock.Controller) *MockUpdater {
	mock := &MockUpdater{ctrl: ctrl}
	mock.recorder = &MockUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpdater) EXPECT() *MockUpdaterMockRecorder {
	return m.recorder
}

// GetHash mocks base method
func (m *MockUpdater) GetHash(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHash", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHash indicates an expected call of GetHash
func (mr *MockUpdaterMockRecorder) GetHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHash", reflect.TypeOf((*MockUpdater)(nil).GetHash), arg0)
}

// PutVulnerabilities mocks base method
func (m *MockUpdater) PutVulnerabilities(arg0, arg1 string, arg2 []*claircore.Vulnerability) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutVulnerabilities", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutVulnerabilities indicates an expected call of PutVulnerabilities
func (mr *MockUpdaterMockRecorder) PutVulnerabilities(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutVulnerabilities", reflect.TypeOf((*MockUpdater)(nil).PutVulnerabilities), arg0, arg1, arg2)
}