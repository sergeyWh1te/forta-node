// Code generated by MockGen. DO NOT EDIT.
// Source: store/registry.go

// Package mock_store is a generated GoMock package.
package mock_store

import (
	reflect "reflect"

	config "github.com/forta-network/forta-node/config"
	gomock "github.com/golang/mock/gomock"
)

// MockRegistryStore is a mock of RegistryStore interface.
type MockRegistryStore struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryStoreMockRecorder
}

// MockRegistryStoreMockRecorder is the mock recorder for MockRegistryStore.
type MockRegistryStoreMockRecorder struct {
	mock *MockRegistryStore
}

// NewMockRegistryStore creates a new mock instance.
func NewMockRegistryStore(ctrl *gomock.Controller) *MockRegistryStore {
	mock := &MockRegistryStore{ctrl: ctrl}
	mock.recorder = &MockRegistryStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistryStore) EXPECT() *MockRegistryStoreMockRecorder {
	return m.recorder
}

// FindAgentGlobally mocks base method.
func (m *MockRegistryStore) FindAgentGlobally(agentID string) (*config.AgentConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAgentGlobally", agentID)
	ret0, _ := ret[0].(*config.AgentConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAgentGlobally indicates an expected call of FindAgentGlobally.
func (mr *MockRegistryStoreMockRecorder) FindAgentGlobally(agentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAgentGlobally", reflect.TypeOf((*MockRegistryStore)(nil).FindAgentGlobally), agentID)
}

// GetAgentsIfChanged mocks base method.
func (m *MockRegistryStore) GetAgentsIfChanged(scanner string) ([]config.AgentConfig, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentsIfChanged", scanner)
	ret0, _ := ret[0].([]config.AgentConfig)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAgentsIfChanged indicates an expected call of GetAgentsIfChanged.
func (mr *MockRegistryStoreMockRecorder) GetAgentsIfChanged(scanner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentsIfChanged", reflect.TypeOf((*MockRegistryStore)(nil).GetAgentsIfChanged), scanner)
}
