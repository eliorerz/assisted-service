// Code generated by MockGen. DO NOT EDIT.
// Source: generator.go

// Package mirrorregistries is a generated GoMock package.
package mirrorregistries

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockServiceMirrorRegistriesConfigBuilder is a mock of ServiceMirrorRegistriesConfigBuilder interface.
type MockServiceMirrorRegistriesConfigBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMirrorRegistriesConfigBuilderMockRecorder
}

// MockServiceMirrorRegistriesConfigBuilderMockRecorder is the mock recorder for MockServiceMirrorRegistriesConfigBuilder.
type MockServiceMirrorRegistriesConfigBuilderMockRecorder struct {
	mock *MockServiceMirrorRegistriesConfigBuilder
}

// NewMockServiceMirrorRegistriesConfigBuilder creates a new mock instance.
func NewMockServiceMirrorRegistriesConfigBuilder(ctrl *gomock.Controller) *MockServiceMirrorRegistriesConfigBuilder {
	mock := &MockServiceMirrorRegistriesConfigBuilder{ctrl: ctrl}
	mock.recorder = &MockServiceMirrorRegistriesConfigBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceMirrorRegistriesConfigBuilder) EXPECT() *MockServiceMirrorRegistriesConfigBuilderMockRecorder {
	return m.recorder
}

// ExtractLocationMirrorDataFromRegistries mocks base method.
func (m *MockServiceMirrorRegistriesConfigBuilder) ExtractLocationMirrorDataFromRegistries() ([]RegistriesConf, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractLocationMirrorDataFromRegistries")
	ret0, _ := ret[0].([]RegistriesConf)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExtractLocationMirrorDataFromRegistries indicates an expected call of ExtractLocationMirrorDataFromRegistries.
func (mr *MockServiceMirrorRegistriesConfigBuilderMockRecorder) ExtractLocationMirrorDataFromRegistries() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractLocationMirrorDataFromRegistries", reflect.TypeOf((*MockServiceMirrorRegistriesConfigBuilder)(nil).ExtractLocationMirrorDataFromRegistries))
}

// GenerateInsecurePolicyJSON mocks base method.
func (m *MockServiceMirrorRegistriesConfigBuilder) GenerateInsecurePolicyJSON() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateInsecurePolicyJSON")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateInsecurePolicyJSON indicates an expected call of GenerateInsecurePolicyJSON.
func (mr *MockServiceMirrorRegistriesConfigBuilderMockRecorder) GenerateInsecurePolicyJSON() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateInsecurePolicyJSON", reflect.TypeOf((*MockServiceMirrorRegistriesConfigBuilder)(nil).GenerateInsecurePolicyJSON))
}

// GetMirrorCA mocks base method.
func (m *MockServiceMirrorRegistriesConfigBuilder) GetMirrorCA() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMirrorCA")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMirrorCA indicates an expected call of GetMirrorCA.
func (mr *MockServiceMirrorRegistriesConfigBuilderMockRecorder) GetMirrorCA() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMirrorCA", reflect.TypeOf((*MockServiceMirrorRegistriesConfigBuilder)(nil).GetMirrorCA))
}

// GetMirrorRegistries mocks base method.
func (m *MockServiceMirrorRegistriesConfigBuilder) GetMirrorRegistries() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMirrorRegistries")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMirrorRegistries indicates an expected call of GetMirrorRegistries.
func (mr *MockServiceMirrorRegistriesConfigBuilderMockRecorder) GetMirrorRegistries() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMirrorRegistries", reflect.TypeOf((*MockServiceMirrorRegistriesConfigBuilder)(nil).GetMirrorRegistries))
}

// IsMirrorRegistriesConfigured mocks base method.
func (m *MockServiceMirrorRegistriesConfigBuilder) IsMirrorRegistriesConfigured() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMirrorRegistriesConfigured")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsMirrorRegistriesConfigured indicates an expected call of IsMirrorRegistriesConfigured.
func (mr *MockServiceMirrorRegistriesConfigBuilderMockRecorder) IsMirrorRegistriesConfigured() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMirrorRegistriesConfigured", reflect.TypeOf((*MockServiceMirrorRegistriesConfigBuilder)(nil).IsMirrorRegistriesConfigured))
}
