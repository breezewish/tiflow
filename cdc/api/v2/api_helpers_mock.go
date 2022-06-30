// Code generated by MockGen. DO NOT EDIT.
// Source: cdc/api/v2/api_helpers.go

// Package v2 is a generated GoMock package.
package v2

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	kv "github.com/pingcap/tidb/kv"
	model "github.com/pingcap/tiflow/cdc/model"
	owner "github.com/pingcap/tiflow/cdc/owner"
	security "github.com/pingcap/tiflow/pkg/security"
	client "github.com/tikv/pd/client"
)

// MockAPIV2Helpers is a mock of APIV2Helpers interface.
type MockAPIV2Helpers struct {
	ctrl     *gomock.Controller
	recorder *MockAPIV2HelpersMockRecorder
}

// MockAPIV2HelpersMockRecorder is the mock recorder for MockAPIV2Helpers.
type MockAPIV2HelpersMockRecorder struct {
	mock *MockAPIV2Helpers
}

// NewMockAPIV2Helpers creates a new mock instance.
func NewMockAPIV2Helpers(ctrl *gomock.Controller) *MockAPIV2Helpers {
	mock := &MockAPIV2Helpers{ctrl: ctrl}
	mock.recorder = &MockAPIV2HelpersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIV2Helpers) EXPECT() *MockAPIV2HelpersMockRecorder {
	return m.recorder
}

// createTiStore mocks base method.
func (m *MockAPIV2Helpers) createTiStore(arg0 []string, arg1 *security.Credential) (kv.Storage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "createTiStore", arg0, arg1)
	ret0, _ := ret[0].(kv.Storage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// createTiStore indicates an expected call of createTiStore.
func (mr *MockAPIV2HelpersMockRecorder) createTiStore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "createTiStore", reflect.TypeOf((*MockAPIV2Helpers)(nil).createTiStore), arg0, arg1)
}

// getPDClient mocks base method.
func (m *MockAPIV2Helpers) getPDClient(arg0 context.Context, arg1 []string, arg2 *security.Credential) (client.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getPDClient", arg0, arg1, arg2)
	ret0, _ := ret[0].(client.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getPDClient indicates an expected call of getPDClient.
func (mr *MockAPIV2HelpersMockRecorder) getPDClient(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getPDClient", reflect.TypeOf((*MockAPIV2Helpers)(nil).getPDClient), arg0, arg1, arg2)
}

// verifyCreateChangefeedConfig mocks base method.
func (m *MockAPIV2Helpers) verifyCreateChangefeedConfig(arg0 context.Context, arg1 *ChangefeedConfig, arg2 client.Client, arg3 owner.StatusProvider, arg4 string, arg5 kv.Storage) (*model.ChangeFeedInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "verifyCreateChangefeedConfig", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(*model.ChangeFeedInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// verifyCreateChangefeedConfig indicates an expected call of verifyCreateChangefeedConfig.
func (mr *MockAPIV2HelpersMockRecorder) verifyCreateChangefeedConfig(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "verifyCreateChangefeedConfig", reflect.TypeOf((*MockAPIV2Helpers)(nil).verifyCreateChangefeedConfig), arg0, arg1, arg2, arg3, arg4, arg5)
}

// verifyResumeChangefeedConfig mocks base method.
func (m *MockAPIV2Helpers) verifyResumeChangefeedConfig(arg0 context.Context, arg1 client.Client, arg2 string, arg3 model.ChangeFeedID, arg4 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "verifyResumeChangefeedConfig", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// verifyResumeChangefeedConfig indicates an expected call of verifyResumeChangefeedConfig.
func (mr *MockAPIV2HelpersMockRecorder) verifyResumeChangefeedConfig(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "verifyResumeChangefeedConfig", reflect.TypeOf((*MockAPIV2Helpers)(nil).verifyResumeChangefeedConfig), arg0, arg1, arg2, arg3, arg4)
}

// verifyUpdateChangefeedConfig mocks base method.
func (m *MockAPIV2Helpers) verifyUpdateChangefeedConfig(arg0 context.Context, arg1 *ChangefeedConfig, arg2 *model.ChangeFeedInfo, arg3 *model.UpstreamInfo) (*model.ChangeFeedInfo, *model.UpstreamInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "verifyUpdateChangefeedConfig", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*model.ChangeFeedInfo)
	ret1, _ := ret[1].(*model.UpstreamInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// verifyUpdateChangefeedConfig indicates an expected call of verifyUpdateChangefeedConfig.
func (mr *MockAPIV2HelpersMockRecorder) verifyUpdateChangefeedConfig(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "verifyUpdateChangefeedConfig", reflect.TypeOf((*MockAPIV2Helpers)(nil).verifyUpdateChangefeedConfig), arg0, arg1, arg2, arg3)
}

// verifyUpstream mocks base method.
func (m *MockAPIV2Helpers) verifyUpstream(arg0 context.Context, arg1 *ChangefeedConfig, arg2 *model.ChangeFeedInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "verifyUpstream", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// verifyUpstream indicates an expected call of verifyUpstream.
func (mr *MockAPIV2HelpersMockRecorder) verifyUpstream(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "verifyUpstream", reflect.TypeOf((*MockAPIV2Helpers)(nil).verifyUpstream), arg0, arg1, arg2)
}
