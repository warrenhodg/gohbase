// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tsuna/gohbase (interfaces: AdminClient)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	hrpc "github.com/tsuna/gohbase/hrpc"
	pb "github.com/tsuna/gohbase/pb"
)

// MockAdminClient is a mock of AdminClient interface.
type MockAdminClient struct {
	ctrl     *gomock.Controller
	recorder *MockAdminClientMockRecorder
}

// MockAdminClientMockRecorder is the mock recorder for MockAdminClient.
type MockAdminClientMockRecorder struct {
	mock *MockAdminClient
}

// NewMockAdminClient creates a new mock instance.
func NewMockAdminClient(ctrl *gomock.Controller) *MockAdminClient {
	mock := &MockAdminClient{ctrl: ctrl}
	mock.recorder = &MockAdminClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminClient) EXPECT() *MockAdminClientMockRecorder {
	return m.recorder
}

// ClusterStatus mocks base method.
func (m *MockAdminClient) ClusterStatus() (*pb.ClusterStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterStatus")
	ret0, _ := ret[0].(*pb.ClusterStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClusterStatus indicates an expected call of ClusterStatus.
func (mr *MockAdminClientMockRecorder) ClusterStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterStatus", reflect.TypeOf((*MockAdminClient)(nil).ClusterStatus))
}

// CreateSnapshot mocks base method.
func (m *MockAdminClient) CreateSnapshot(arg0 *hrpc.Snapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSnapshot", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSnapshot indicates an expected call of CreateSnapshot.
func (mr *MockAdminClientMockRecorder) CreateSnapshot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnapshot", reflect.TypeOf((*MockAdminClient)(nil).CreateSnapshot), arg0)
}

// CreateTable mocks base method.
func (m *MockAdminClient) CreateTable(arg0 *hrpc.CreateTable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTable", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTable indicates an expected call of CreateTable.
func (mr *MockAdminClientMockRecorder) CreateTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTable", reflect.TypeOf((*MockAdminClient)(nil).CreateTable), arg0)
}

// DeleteSnapshot mocks base method.
func (m *MockAdminClient) DeleteSnapshot(arg0 *hrpc.Snapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSnapshot", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSnapshot indicates an expected call of DeleteSnapshot.
func (mr *MockAdminClientMockRecorder) DeleteSnapshot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnapshot", reflect.TypeOf((*MockAdminClient)(nil).DeleteSnapshot), arg0)
}

// DeleteTable mocks base method.
func (m *MockAdminClient) DeleteTable(arg0 *hrpc.DeleteTable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTable", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTable indicates an expected call of DeleteTable.
func (mr *MockAdminClientMockRecorder) DeleteTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTable", reflect.TypeOf((*MockAdminClient)(nil).DeleteTable), arg0)
}

// DisableTable mocks base method.
func (m *MockAdminClient) DisableTable(arg0 *hrpc.DisableTable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableTable", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableTable indicates an expected call of DisableTable.
func (mr *MockAdminClientMockRecorder) DisableTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableTable", reflect.TypeOf((*MockAdminClient)(nil).DisableTable), arg0)
}

// EnableTable mocks base method.
func (m *MockAdminClient) EnableTable(arg0 *hrpc.EnableTable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableTable", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableTable indicates an expected call of EnableTable.
func (mr *MockAdminClientMockRecorder) EnableTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableTable", reflect.TypeOf((*MockAdminClient)(nil).EnableTable), arg0)
}

// ListSnapshots mocks base method.
func (m *MockAdminClient) ListSnapshots(arg0 *hrpc.ListSnapshots) ([]*pb.SnapshotDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSnapshots", arg0)
	ret0, _ := ret[0].([]*pb.SnapshotDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSnapshots indicates an expected call of ListSnapshots.
func (mr *MockAdminClientMockRecorder) ListSnapshots(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSnapshots", reflect.TypeOf((*MockAdminClient)(nil).ListSnapshots), arg0)
}

// ListTableNames mocks base method.
func (m *MockAdminClient) ListTableNames(arg0 *hrpc.ListTableNames) ([]*pb.TableName, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTableNames", arg0)
	ret0, _ := ret[0].([]*pb.TableName)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTableNames indicates an expected call of ListTableNames.
func (mr *MockAdminClientMockRecorder) ListTableNames(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTableNames", reflect.TypeOf((*MockAdminClient)(nil).ListTableNames), arg0)
}

// MoveRegion mocks base method.
func (m *MockAdminClient) MoveRegion(arg0 *hrpc.MoveRegion) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveRegion", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveRegion indicates an expected call of MoveRegion.
func (mr *MockAdminClientMockRecorder) MoveRegion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveRegion", reflect.TypeOf((*MockAdminClient)(nil).MoveRegion), arg0)
}

// RestoreSnapshot mocks base method.
func (m *MockAdminClient) RestoreSnapshot(arg0 *hrpc.Snapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreSnapshot", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RestoreSnapshot indicates an expected call of RestoreSnapshot.
func (mr *MockAdminClientMockRecorder) RestoreSnapshot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreSnapshot", reflect.TypeOf((*MockAdminClient)(nil).RestoreSnapshot), arg0)
}

// SetBalancer mocks base method.
func (m *MockAdminClient) SetBalancer(arg0 *hrpc.SetBalancer) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBalancer", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetBalancer indicates an expected call of SetBalancer.
func (mr *MockAdminClientMockRecorder) SetBalancer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBalancer", reflect.TypeOf((*MockAdminClient)(nil).SetBalancer), arg0)
}
