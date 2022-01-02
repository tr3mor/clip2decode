// Code generated by MockGen. DO NOT EDIT.
// Source: clipboard.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockClipboardInterface is a mock of ClipboardInterface interface.
type MockClipboardInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClipboardInterfaceMockRecorder
}

// MockClipboardInterfaceMockRecorder is the mock recorder for MockClipboardInterface.
type MockClipboardInterfaceMockRecorder struct {
	mock *MockClipboardInterface
}

// NewMockClipboardInterface creates a new mock instance.
func NewMockClipboardInterface(ctrl *gomock.Controller) *MockClipboardInterface {
	mock := &MockClipboardInterface{ctrl: ctrl}
	mock.recorder = &MockClipboardInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClipboardInterface) EXPECT() *MockClipboardInterfaceMockRecorder {
	return m.recorder
}

// GetData mocks base method.
func (m *MockClipboardInterface) GetData() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetData indicates an expected call of GetData.
func (mr *MockClipboardInterfaceMockRecorder) GetData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockClipboardInterface)(nil).GetData))
}

// WriteData mocks base method.
func (m *MockClipboardInterface) WriteData(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteData", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteData indicates an expected call of WriteData.
func (mr *MockClipboardInterfaceMockRecorder) WriteData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteData", reflect.TypeOf((*MockClipboardInterface)(nil).WriteData), arg0)
}