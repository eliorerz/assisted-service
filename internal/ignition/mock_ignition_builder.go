// Code generated by MockGen. DO NOT EDIT.
// Source: discovery.go

// Package ignition is a generated GoMock package.
package ignition

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/openshift/assisted-service/api/hiveextension/v1beta1"
	common "github.com/openshift/assisted-service/internal/common"
	models "github.com/openshift/assisted-service/models"
	auth "github.com/openshift/assisted-service/pkg/auth"
)

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
func (m *MockIgnitionBuilder) FormatDiscoveryIgnitionFile(ctx context.Context, infraEnv *common.InfraEnv, cfg IgnitionConfig, safeForLogs bool, authType auth.AuthType, overrideDiscoveryISOType string, mirrorRegistryConfiguration *v1beta1.MirrorRegistryConfiguration) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormatDiscoveryIgnitionFile", ctx, infraEnv, cfg, safeForLogs, authType, overrideDiscoveryISOType, mirrorRegistryConfiguration)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FormatDiscoveryIgnitionFile indicates an expected call of FormatDiscoveryIgnitionFile.
func (mr *MockIgnitionBuilderMockRecorder) FormatDiscoveryIgnitionFile(ctx, infraEnv, cfg, safeForLogs, authType, overrideDiscoveryISOType, mirrorRegistryConfiguration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormatDiscoveryIgnitionFile", reflect.TypeOf((*MockIgnitionBuilder)(nil).FormatDiscoveryIgnitionFile), ctx, infraEnv, cfg, safeForLogs, authType, overrideDiscoveryISOType, mirrorRegistryConfiguration)
}

// FormatSecondDayWorkerIgnitionFile mocks base method.
func (m *MockIgnitionBuilder) FormatSecondDayWorkerIgnitionFile(url string, caCert *string, bearerToken, ignitionEndpointHTTPHeaders string, host *models.Host) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormatSecondDayWorkerIgnitionFile", url, caCert, bearerToken, ignitionEndpointHTTPHeaders, host)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FormatSecondDayWorkerIgnitionFile indicates an expected call of FormatSecondDayWorkerIgnitionFile.
func (mr *MockIgnitionBuilderMockRecorder) FormatSecondDayWorkerIgnitionFile(url, caCert, bearerToken, ignitionEndpointHTTPHeaders, host interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormatSecondDayWorkerIgnitionFile", reflect.TypeOf((*MockIgnitionBuilder)(nil).FormatSecondDayWorkerIgnitionFile), url, caCert, bearerToken, ignitionEndpointHTTPHeaders, host)
}
