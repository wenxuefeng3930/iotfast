// Code generated by MockGen. DO NOT EDIT.
// Source: persistence/subscription/subscription.go

// Package subscription is a generated GoMock package.
package subscription

import (
	gmqtt  "github.com/xiaodingding/iotfast/server/mqtt"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockStore) Init(clientIDs []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", clientIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockStoreMockRecorder) Init(clientIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockStore)(nil).Init), clientIDs)
}

// Subscribe mocks base method
func (m *MockStore) Subscribe(clientID string, subscriptions ...*gmqtt.Subscription) (SubscribeResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{clientID}
	for _, a := range subscriptions {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(SubscribeResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockStoreMockRecorder) Subscribe(clientID interface{}, subscriptions ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{clientID}, subscriptions...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockStore)(nil).Subscribe), varargs...)
}

// Unsubscribe mocks base method
func (m *MockStore) Unsubscribe(clientID string, topics ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{clientID}
	for _, a := range topics {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Unsubscribe", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe
func (mr *MockStoreMockRecorder) Unsubscribe(clientID interface{}, topics ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{clientID}, topics...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockStore)(nil).Unsubscribe), varargs...)
}

// UnsubscribeAll mocks base method
func (m *MockStore) UnsubscribeAll(clientID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsubscribeAll", clientID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsubscribeAll indicates an expected call of UnsubscribeAll
func (mr *MockStoreMockRecorder) UnsubscribeAll(clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeAll", reflect.TypeOf((*MockStore)(nil).UnsubscribeAll), clientID)
}

// Iterate mocks base method
func (m *MockStore) Iterate(fn IterateFn, options IterationOptions) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Iterate", fn, options)
}

// Iterate indicates an expected call of Iterate
func (mr *MockStoreMockRecorder) Iterate(fn, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iterate", reflect.TypeOf((*MockStore)(nil).Iterate), fn, options)
}

// Close mocks base method
func (m *MockStore) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockStoreMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStore)(nil).Close))
}

// GetStats mocks base method
func (m *MockStore) GetStats() Stats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStats")
	ret0, _ := ret[0].(Stats)
	return ret0
}

// GetStats indicates an expected call of GetStats
func (mr *MockStoreMockRecorder) GetStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStats", reflect.TypeOf((*MockStore)(nil).GetStats))
}

// GetClientStats mocks base method
func (m *MockStore) GetClientStats(clientID string) (Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientStats", clientID)
	ret0, _ := ret[0].(Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientStats indicates an expected call of GetClientStats
func (mr *MockStoreMockRecorder) GetClientStats(clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientStats", reflect.TypeOf((*MockStore)(nil).GetClientStats), clientID)
}

// MockStatsReader is a mock of StatsReader interface
type MockStatsReader struct {
	ctrl     *gomock.Controller
	recorder *MockStatsReaderMockRecorder
}

// MockStatsReaderMockRecorder is the mock recorder for MockStatsReader
type MockStatsReaderMockRecorder struct {
	mock *MockStatsReader
}

// NewMockStatsReader creates a new mock instance
func NewMockStatsReader(ctrl *gomock.Controller) *MockStatsReader {
	mock := &MockStatsReader{ctrl: ctrl}
	mock.recorder = &MockStatsReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStatsReader) EXPECT() *MockStatsReaderMockRecorder {
	return m.recorder
}

// GetStats mocks base method
func (m *MockStatsReader) GetStats() Stats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStats")
	ret0, _ := ret[0].(Stats)
	return ret0
}

// GetStats indicates an expected call of GetStats
func (mr *MockStatsReaderMockRecorder) GetStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStats", reflect.TypeOf((*MockStatsReader)(nil).GetStats))
}

// GetClientStats mocks base method
func (m *MockStatsReader) GetClientStats(clientID string) (Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientStats", clientID)
	ret0, _ := ret[0].(Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientStats indicates an expected call of GetClientStats
func (mr *MockStatsReaderMockRecorder) GetClientStats(clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientStats", reflect.TypeOf((*MockStatsReader)(nil).GetClientStats), clientID)
}
