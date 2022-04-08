// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scylladb/scylla-manager/pkg/service/scheduler (interfaces: Runner)

// Package scheduler is a generated GoMock package.
package scheduler

import (
	context "context"
	json "encoding/json"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/scylladb/scylla-manager/v3/pkg/util/uuid"
	reflect "reflect"
)

// mockRunner is a mock of Runner interface
type mockRunner struct {
	ctrl     *gomock.Controller
	recorder *mockRunnerMockRecorder
}

// mockRunnerMockRecorder is the mock recorder for mockRunner
type mockRunnerMockRecorder struct {
	mock *mockRunner
}

// NewmockRunner creates a new mock instance
func NewmockRunner(ctrl *gomock.Controller) *mockRunner {
	mock := &mockRunner{ctrl: ctrl}
	mock.recorder = &mockRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *mockRunner) EXPECT() *mockRunnerMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *mockRunner) Run(arg0 context.Context, arg1, arg2, arg3 uuid.UUID, arg4 json.RawMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *mockRunnerMockRecorder) Run(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*mockRunner)(nil).Run), arg0, arg1, arg2, arg3, arg4)
}
