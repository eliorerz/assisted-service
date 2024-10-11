// Code generated by MockGen. DO NOT EDIT.
// Source: builder.go

// Package builder is a generated GoMock package.
package builder

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/openshift/assisted-service/api/hiveextension/v1beta1"
	common "github.com/openshift/assisted-service/internal/common"
)

// MockInstallConfigBuilder is a mock of InstallConfigBuilder interface.
type MockInstallConfigBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockInstallConfigBuilderMockRecorder
}

// MockInstallConfigBuilderMockRecorder is the mock recorder for MockInstallConfigBuilder.
type MockInstallConfigBuilderMockRecorder struct {
	mock *MockInstallConfigBuilder
}

// NewMockInstallConfigBuilder creates a new mock instance.
func NewMockInstallConfigBuilder(ctrl *gomock.Controller) *MockInstallConfigBuilder {
	mock := &MockInstallConfigBuilder{ctrl: ctrl}
	mock.recorder = &MockInstallConfigBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstallConfigBuilder) EXPECT() *MockInstallConfigBuilderMockRecorder {
	return m.recorder
}

// GetInstallConfig mocks base method.
func (m *MockInstallConfigBuilder) GetInstallConfig(cluster *common.Cluster, clusterInfraenvs []*common.InfraEnv, rhRootCA string, mrConfiguration *v1beta1.MirrorRegistryConfiguration) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstallConfig", cluster, clusterInfraenvs, rhRootCA, mrConfiguration)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstallConfig indicates an expected call of GetInstallConfig.
func (mr *MockInstallConfigBuilderMockRecorder) GetInstallConfig(cluster, clusterInfraenvs, rhRootCA, mrConfiguration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstallConfig", reflect.TypeOf((*MockInstallConfigBuilder)(nil).GetInstallConfig), cluster, clusterInfraenvs, rhRootCA, mrConfiguration)
}

// ValidateInstallConfigPatch mocks base method.
func (m *MockInstallConfigBuilder) ValidateInstallConfigPatch(cluster *common.Cluster, clusterInfraenvs []*common.InfraEnv, patch string, mrConfiguration *v1beta1.MirrorRegistryConfiguration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateInstallConfigPatch", cluster, clusterInfraenvs, patch, mrConfiguration)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateInstallConfigPatch indicates an expected call of ValidateInstallConfigPatch.
func (mr *MockInstallConfigBuilderMockRecorder) ValidateInstallConfigPatch(cluster, clusterInfraenvs, patch, mrConfiguration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateInstallConfigPatch", reflect.TypeOf((*MockInstallConfigBuilder)(nil).ValidateInstallConfigPatch), cluster, clusterInfraenvs, patch, mrConfiguration)
}
