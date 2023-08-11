// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rancher/fleet/pkg/generated/controllers/fleet.cattle.io/v1alpha1 (interfaces: BundleDeploymentCache)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/rancher/fleet/pkg/apis/fleet.cattle.io/v1alpha1"
	"github.com/rancher/wrangler/pkg/generic"
	labels "k8s.io/apimachinery/pkg/labels"
	
)

// MockBundleDeploymentCache is a mock of BundleDeploymentCache interface.
type MockBundleDeploymentCache struct {
	ctrl     *gomock.Controller
	recorder *MockBundleDeploymentCacheMockRecorder
}

// MockBundleDeploymentCacheMockRecorder is the mock recorder for MockBundleDeploymentCache.
type MockBundleDeploymentCacheMockRecorder struct {
	mock *MockBundleDeploymentCache
}

// NewMockBundleDeploymentCache creates a new mock instance.
func NewMockBundleDeploymentCache(ctrl *gomock.Controller) *MockBundleDeploymentCache {
	mock := &MockBundleDeploymentCache{ctrl: ctrl}
	mock.recorder = &MockBundleDeploymentCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBundleDeploymentCache) EXPECT() *MockBundleDeploymentCacheMockRecorder {
	return m.recorder
}

// AddIndexer mocks base method.
func (m *MockBundleDeploymentCache) AddIndexer(arg0 string, arg1 generic.Indexer[*v1alpha1.BundleDeployment]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddIndexer", arg0, arg1)
}

// AddIndexer indicates an expected call of AddIndexer.
func (mr *MockBundleDeploymentCacheMockRecorder) AddIndexer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIndexer", reflect.TypeOf((*MockBundleDeploymentCache)(nil).AddIndexer), arg0, arg1)
}

// Get mocks base method.
func (m *MockBundleDeploymentCache) Get(arg0, arg1 string) (*v1alpha1.BundleDeployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.BundleDeployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockBundleDeploymentCacheMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBundleDeploymentCache)(nil).Get), arg0, arg1)
}

// GetByIndex mocks base method.
func (m *MockBundleDeploymentCache) GetByIndex(arg0, arg1 string) ([]*v1alpha1.BundleDeployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIndex", arg0, arg1)
	ret0, _ := ret[0].([]*v1alpha1.BundleDeployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIndex indicates an expected call of GetByIndex.
func (mr *MockBundleDeploymentCacheMockRecorder) GetByIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIndex", reflect.TypeOf((*MockBundleDeploymentCache)(nil).GetByIndex), arg0, arg1)
}

// List mocks base method.
func (m *MockBundleDeploymentCache) List(arg0 string, arg1 labels.Selector) ([]*v1alpha1.BundleDeployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*v1alpha1.BundleDeployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockBundleDeploymentCacheMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockBundleDeploymentCache)(nil).List), arg0, arg1)
}
