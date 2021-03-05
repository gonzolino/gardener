// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/pkg/chartrenderer (interfaces: Interface,Factory)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	chartrenderer "github.com/gardener/gardener/pkg/chartrenderer"
	gomock "github.com/golang/mock/gomock"
	rest "k8s.io/client-go/rest"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Render mocks base method.
func (m *MockInterface) Render(arg0, arg1, arg2 string, arg3 interface{}) (*chartrenderer.RenderedChart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Render", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*chartrenderer.RenderedChart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Render indicates an expected call of Render.
func (mr *MockInterfaceMockRecorder) Render(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Render", reflect.TypeOf((*MockInterface)(nil).Render), arg0, arg1, arg2, arg3)
}

// RenderArchive mocks base method.
func (m *MockInterface) RenderArchive(arg0 []byte, arg1, arg2 string, arg3 interface{}) (*chartrenderer.RenderedChart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenderArchive", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*chartrenderer.RenderedChart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenderArchive indicates an expected call of RenderArchive.
func (mr *MockInterfaceMockRecorder) RenderArchive(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenderArchive", reflect.TypeOf((*MockInterface)(nil).RenderArchive), arg0, arg1, arg2, arg3)
}

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// NewForConfig mocks base method.
func (m *MockFactory) NewForConfig(arg0 *rest.Config) (chartrenderer.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewForConfig", arg0)
	ret0, _ := ret[0].(chartrenderer.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewForConfig indicates an expected call of NewForConfig.
func (mr *MockFactoryMockRecorder) NewForConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewForConfig", reflect.TypeOf((*MockFactory)(nil).NewForConfig), arg0)
}