// Code generated by MockGen. DO NOT EDIT.
// Source: examples.go

// Package examples is a generated GoMock package.
package examples

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockExampleInterface is a mock of ExampleInterface interface
type MockExampleInterface struct {
	ctrl     *gomock.Controller
	recorder *MockExampleInterfaceMockRecorder
}

// MockExampleInterfaceMockRecorder is the mock recorder for MockExampleInterface
type MockExampleInterfaceMockRecorder struct {
	mock *MockExampleInterface
}

// NewMockExampleInterface creates a new mock instance
func NewMockExampleInterface(ctrl *gomock.Controller) *MockExampleInterface {
	mock := &MockExampleInterface{ctrl: ctrl}
	mock.recorder = &MockExampleInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExampleInterface) EXPECT() *MockExampleInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockExampleInterface) Create(name string, hashedPassword []byte, createdAt, updatedAt time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", name, hashedPassword, createdAt, updatedAt)
}

// Create indicates an expected call of Create
func (mr *MockExampleInterfaceMockRecorder) Create(name, hashedPassword, createdAt, updatedAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockExampleInterface)(nil).Create), name, hashedPassword, createdAt, updatedAt)
}
