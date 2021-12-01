// Code generated by MockGen. DO NOT EDIT.
// Source: ignition.go

// Package ignition is a generated GoMock package.
package ignition

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	common "github.com/openshift/assisted-service/internal/common"
	models "github.com/openshift/assisted-service/models"
	auth "github.com/openshift/assisted-service/pkg/auth"
)

// MockGenerator is a mock of Generator interface.
type MockGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockGeneratorMockRecorder
}

// MockGeneratorMockRecorder is the mock recorder for MockGenerator.
type MockGeneratorMockRecorder struct {
	mock *MockGenerator
}

// NewMockGenerator creates a new mock instance.
func NewMockGenerator(ctrl *gomock.Controller) *MockGenerator {
	mock := &MockGenerator{ctrl: ctrl}
	mock.recorder = &MockGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenerator) EXPECT() *MockGeneratorMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockGenerator) Generate(ctx context.Context, installConfig []byte, platformType models.PlatformType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", ctx, installConfig, platformType)
	ret0, _ := ret[0].(error)
	return ret0
}

// Generate indicates an expected call of Generate.
func (mr *MockGeneratorMockRecorder) Generate(ctx, installConfig, platformType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockGenerator)(nil).Generate), ctx, installConfig, platformType)
}

// UpdateEtcHosts mocks base method.
func (m *MockGenerator) UpdateEtcHosts(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEtcHosts", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEtcHosts indicates an expected call of UpdateEtcHosts.
func (mr *MockGeneratorMockRecorder) UpdateEtcHosts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEtcHosts", reflect.TypeOf((*MockGenerator)(nil).UpdateEtcHosts), arg0)
}

// UploadToS3 mocks base method.
func (m *MockGenerator) UploadToS3(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadToS3", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadToS3 indicates an expected call of UploadToS3.
func (mr *MockGeneratorMockRecorder) UploadToS3(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadToS3", reflect.TypeOf((*MockGenerator)(nil).UploadToS3), ctx)
}

// MockIgnitionBuilder is a mock of IgnitionBuilder interface.
type MockIgnitionBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockIgnitionBuilderMockRecorder
}

// MockIgnitionBuilderMockRecorder is the mock recorder for MockIgnitionBuilder.
type MockIgnitionBuilderMockRecorder struct {
	mock *MockIgnitionBuilder
}

// NewMockIgnitionBuilder creates a new mock instance.
func NewMockIgnitionBuilder(ctrl *gomock.Controller) *MockIgnitionBuilder {
	mock := &MockIgnitionBuilder{ctrl: ctrl}
	mock.recorder = &MockIgnitionBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIgnitionBuilder) EXPECT() *MockIgnitionBuilderMockRecorder {
	return m.recorder
}

// FormatDiscoveryIgnitionFile mocks base method.
func (m *MockIgnitionBuilder) FormatDiscoveryIgnitionFile(ctx context.Context, infraEnv *common.InfraEnv, cfg IgnitionConfig, safeForLogs bool, authType auth.AuthType) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormatDiscoveryIgnitionFile", ctx, infraEnv, cfg, safeForLogs, authType)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FormatDiscoveryIgnitionFile indicates an expected call of FormatDiscoveryIgnitionFile.
func (mr *MockIgnitionBuilderMockRecorder) FormatDiscoveryIgnitionFile(ctx, infraEnv, cfg, safeForLogs, authType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormatDiscoveryIgnitionFile", reflect.TypeOf((*MockIgnitionBuilder)(nil).FormatDiscoveryIgnitionFile), ctx, infraEnv, cfg, safeForLogs, authType)
}

// FormatSecondDayWorkerIgnitionFile mocks base method.
func (m *MockIgnitionBuilder) FormatSecondDayWorkerIgnitionFile(url string, caCert *string, bearerToken string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormatSecondDayWorkerIgnitionFile", url, caCert, bearerToken)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FormatSecondDayWorkerIgnitionFile indicates an expected call of FormatSecondDayWorkerIgnitionFile.
func (mr *MockIgnitionBuilderMockRecorder) FormatSecondDayWorkerIgnitionFile(url, caCert, bearerToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormatSecondDayWorkerIgnitionFile", reflect.TypeOf((*MockIgnitionBuilder)(nil).FormatSecondDayWorkerIgnitionFile), url, caCert, bearerToken)
}
