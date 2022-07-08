// Code generated by MockGen. DO NOT EDIT.
// Source: server/server.go

// Package server is a generated GoMock package.
package server

import (
	context "context"
	config "github.com/xiaodingding/iotfast/server/mqtt/config"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockServer is a mock of Server interface
type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
}

// MockServerMockRecorder is the mock recorder for MockServer
type MockServerMockRecorder struct {
	mock *MockServer
}

// NewMockServer creates a new mock instance
func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}

// Publisher mocks base method
func (m *MockServer) Publisher() Publisher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publisher")
	ret0, _ := ret[0].(Publisher)
	return ret0
}

// Publisher indicates an expected call of Publisher
func (mr *MockServerMockRecorder) Publisher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publisher", reflect.TypeOf((*MockServer)(nil).Publisher))
}

// GetConfig mocks base method
func (m *MockServer) GetConfig() config.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(config.Config)
	return ret0
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockServerMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockServer)(nil).GetConfig))
}

// StatsManager mocks base method
func (m *MockServer) StatsManager() StatsReader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatsManager")
	ret0, _ := ret[0].(StatsReader)
	return ret0
}

// StatsManager indicates an expected call of StatsManager
func (mr *MockServerMockRecorder) StatsManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatsManager", reflect.TypeOf((*MockServer)(nil).StatsManager))
}

// Stop mocks base method
func (m *MockServer) Stop(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockServerMockRecorder) Stop(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockServer)(nil).Stop), ctx)
}

// ApplyConfig mocks base method
func (m *MockServer) ApplyConfig(config config.Config) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ApplyConfig", config)
}

// ApplyConfig indicates an expected call of ApplyConfig
func (mr *MockServerMockRecorder) ApplyConfig(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyConfig", reflect.TypeOf((*MockServer)(nil).ApplyConfig), config)
}

// ClientService mocks base method
func (m *MockServer) ClientService() ClientService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientService")
	ret0, _ := ret[0].(ClientService)
	return ret0
}

// ClientService indicates an expected call of ClientService
func (mr *MockServerMockRecorder) ClientService() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientService", reflect.TypeOf((*MockServer)(nil).ClientService))
}

// SubscriptionService mocks base method
func (m *MockServer) SubscriptionService() SubscriptionService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionService")
	ret0, _ := ret[0].(SubscriptionService)
	return ret0
}

// SubscriptionService indicates an expected call of SubscriptionService
func (mr *MockServerMockRecorder) SubscriptionService() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionService", reflect.TypeOf((*MockServer)(nil).SubscriptionService))
}

// RetainedService mocks base method
func (m *MockServer) RetainedService() RetainedService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetainedService")
	ret0, _ := ret[0].(RetainedService)
	return ret0
}

// RetainedService indicates an expected call of RetainedService
func (mr *MockServerMockRecorder) RetainedService() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetainedService", reflect.TypeOf((*MockServer)(nil).RetainedService))
}

// Plugins mocks base method
func (m *MockServer) Plugins() []Plugin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Plugins")
	ret0, _ := ret[0].([]Plugin)
	return ret0
}

// Plugins indicates an expected call of Plugins
func (mr *MockServerMockRecorder) Plugins() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Plugins", reflect.TypeOf((*MockServer)(nil).Plugins))
}

// APIRegistrar mocks base method
func (m *MockServer) APIRegistrar() APIRegistrar {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIRegistrar")
	ret0, _ := ret[0].(APIRegistrar)
	return ret0
}

// APIRegistrar indicates an expected call of APIRegistrar
func (mr *MockServerMockRecorder) APIRegistrar() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIRegistrar", reflect.TypeOf((*MockServer)(nil).APIRegistrar))
}
