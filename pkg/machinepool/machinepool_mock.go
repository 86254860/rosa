// Code generated by MockGen. DO NOT EDIT.
// Source: machinepool.go
//
// Generated by this command:
//
//	mockgen -source=machinepool.go -package=machinepool -destination=machinepool_mock.go
//

// Package machinepool is a generated GoMock package.
package machinepool

import (
	reflect "reflect"

	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	rosa "github.com/openshift/rosa/pkg/rosa"
	gomock "go.uber.org/mock/gomock"
)

// MockMachinePoolService is a mock of MachinePoolService interface.
type MockMachinePoolService struct {
	ctrl     *gomock.Controller
	recorder *MockMachinePoolServiceMockRecorder
}

// MockMachinePoolServiceMockRecorder is the mock recorder for MockMachinePoolService.
type MockMachinePoolServiceMockRecorder struct {
	mock *MockMachinePoolService
}

// NewMockMachinePoolService creates a new mock instance.
func NewMockMachinePoolService(ctrl *gomock.Controller) *MockMachinePoolService {
	mock := &MockMachinePoolService{ctrl: ctrl}
	mock.recorder = &MockMachinePoolServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachinePoolService) EXPECT() *MockMachinePoolServiceMockRecorder {
	return m.recorder
}

// DescribeMachinePool mocks base method.
func (m *MockMachinePoolService) DescribeMachinePool(r *rosa.Runtime, cluster *v1.Cluster, clusterKey string, isHypershift bool, machinePoolId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeMachinePool", r, cluster, clusterKey, isHypershift, machinePoolId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeMachinePool indicates an expected call of DescribeMachinePool.
func (mr *MockMachinePoolServiceMockRecorder) DescribeMachinePool(r, cluster, clusterKey, isHypershift, machinePoolId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMachinePool", reflect.TypeOf((*MockMachinePoolService)(nil).DescribeMachinePool), r, cluster, clusterKey, isHypershift, machinePoolId)
}
