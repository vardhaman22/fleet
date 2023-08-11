// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rancher/fleet/pkg/generated/controllers/fleet.cattle.io/v1alpha1 (interfaces: ClusterRegistrationController)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/rancher/fleet/pkg/apis/fleet.cattle.io/v1alpha1"
	v1alpha10 "github.com/rancher/fleet/pkg/generated/controllers/fleet.cattle.io/v1alpha1"
	generic "github.com/rancher/wrangler/pkg/generic"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	rest "k8s.io/client-go/rest"
)

// MockClusterRegistrationController is a mock of ClusterRegistrationController interface.
type MockClusterRegistrationController struct {
	ctrl     *gomock.Controller
	recorder *MockClusterRegistrationControllerMockRecorder
}

// MockClusterRegistrationControllerMockRecorder is the mock recorder for MockClusterRegistrationController.
type MockClusterRegistrationControllerMockRecorder struct {
	mock *MockClusterRegistrationController
}

// NewMockClusterRegistrationController creates a new mock instance.
func NewMockClusterRegistrationController(ctrl *gomock.Controller) *MockClusterRegistrationController {
	mock := &MockClusterRegistrationController{ctrl: ctrl}
	mock.recorder = &MockClusterRegistrationControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterRegistrationController) EXPECT() *MockClusterRegistrationControllerMockRecorder {
	return m.recorder
}

// AddGenericHandler mocks base method.
func (m *MockClusterRegistrationController) AddGenericHandler(arg0 context.Context, arg1 string, arg2 generic.Handler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddGenericHandler", arg0, arg1, arg2)
}

// AddGenericHandler indicates an expected call of AddGenericHandler.
func (mr *MockClusterRegistrationControllerMockRecorder) AddGenericHandler(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddGenericHandler", reflect.TypeOf((*MockClusterRegistrationController)(nil).AddGenericHandler), arg0, arg1, arg2)
}

// AddGenericRemoveHandler mocks base method.
func (m *MockClusterRegistrationController) AddGenericRemoveHandler(arg0 context.Context, arg1 string, arg2 generic.Handler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddGenericRemoveHandler", arg0, arg1, arg2)
}

// AddGenericRemoveHandler indicates an expected call of AddGenericRemoveHandler.
func (mr *MockClusterRegistrationControllerMockRecorder) AddGenericRemoveHandler(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddGenericRemoveHandler", reflect.TypeOf((*MockClusterRegistrationController)(nil).AddGenericRemoveHandler), arg0, arg1, arg2)
}

// Cache mocks base method.
func (m *MockClusterRegistrationController) Cache() generic.CacheInterface[*v1alpha1.ClusterRegistration] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cache")
	ret0, _ := ret[0].(v1alpha10.ClusterRegistrationCache)
	return ret0
}

// Cache indicates an expected call of Cache.
func (mr *MockClusterRegistrationControllerMockRecorder) Cache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cache", reflect.TypeOf((*MockClusterRegistrationController)(nil).Cache))
}

// Create mocks base method.
func (m *MockClusterRegistrationController) Create(arg0 *v1alpha1.ClusterRegistration) (*v1alpha1.ClusterRegistration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*v1alpha1.ClusterRegistration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockClusterRegistrationControllerMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClusterRegistrationController)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockClusterRegistrationController) Delete(arg0, arg1 string, arg2 *v1.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockClusterRegistrationControllerMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClusterRegistrationController)(nil).Delete), arg0, arg1, arg2)
}

// Enqueue mocks base method.
func (m *MockClusterRegistrationController) Enqueue(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Enqueue", arg0, arg1)
}

// Enqueue indicates an expected call of Enqueue.
func (mr *MockClusterRegistrationControllerMockRecorder) Enqueue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enqueue", reflect.TypeOf((*MockClusterRegistrationController)(nil).Enqueue), arg0, arg1)
}

// EnqueueAfter mocks base method.
func (m *MockClusterRegistrationController) EnqueueAfter(arg0, arg1 string, arg2 time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnqueueAfter", arg0, arg1, arg2)
}

// EnqueueAfter indicates an expected call of EnqueueAfter.
func (mr *MockClusterRegistrationControllerMockRecorder) EnqueueAfter(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnqueueAfter", reflect.TypeOf((*MockClusterRegistrationController)(nil).EnqueueAfter), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockClusterRegistrationController) Get(arg0, arg1 string, arg2 v1.GetOptions) (*v1alpha1.ClusterRegistration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1alpha1.ClusterRegistration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockClusterRegistrationControllerMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClusterRegistrationController)(nil).Get), arg0, arg1, arg2)
}

// GroupVersionKind mocks base method.
func (m *MockClusterRegistrationController) GroupVersionKind() schema.GroupVersionKind {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GroupVersionKind")
	ret0, _ := ret[0].(schema.GroupVersionKind)
	return ret0
}

// GroupVersionKind indicates an expected call of GroupVersionKind.
func (mr *MockClusterRegistrationControllerMockRecorder) GroupVersionKind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupVersionKind", reflect.TypeOf((*MockClusterRegistrationController)(nil).GroupVersionKind))
}

// Informer mocks base method.
func (m *MockClusterRegistrationController) Informer() cache.SharedIndexInformer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Informer")
	ret0, _ := ret[0].(cache.SharedIndexInformer)
	return ret0
}

// Informer indicates an expected call of Informer.
func (mr *MockClusterRegistrationControllerMockRecorder) Informer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Informer", reflect.TypeOf((*MockClusterRegistrationController)(nil).Informer))
}

// List mocks base method.
func (m *MockClusterRegistrationController) List(arg0 string, arg1 v1.ListOptions) (*v1alpha1.ClusterRegistrationList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.ClusterRegistrationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockClusterRegistrationControllerMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClusterRegistrationController)(nil).List), arg0, arg1)
}

// OnChange mocks base method.
func (m *MockClusterRegistrationController) OnChange(arg0 context.Context, arg1 string, arg2 generic.ObjectHandler[*v1alpha1.ClusterRegistration]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnChange", arg0, arg1, arg2)
}

// OnChange indicates an expected call of OnChange.
func (mr *MockClusterRegistrationControllerMockRecorder) OnChange(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnChange", reflect.TypeOf((*MockClusterRegistrationController)(nil).OnChange), arg0, arg1, arg2)
}

// OnRemove mocks base method.
func (m *MockClusterRegistrationController) OnRemove(arg0 context.Context, arg1 string, arg2 generic.ObjectHandler[*v1alpha1.ClusterRegistration]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnRemove", arg0, arg1, arg2)
}

// OnRemove indicates an expected call of OnRemove.
func (mr *MockClusterRegistrationControllerMockRecorder) OnRemove(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnRemove", reflect.TypeOf((*MockClusterRegistrationController)(nil).OnRemove), arg0, arg1, arg2)
}

// Patch mocks base method.
func (m *MockClusterRegistrationController) Patch(arg0, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 ...string) (*v1alpha1.ClusterRegistration, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1alpha1.ClusterRegistration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockClusterRegistrationControllerMockRecorder) Patch(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockClusterRegistrationController)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockClusterRegistrationController) Update(arg0 *v1alpha1.ClusterRegistration) (*v1alpha1.ClusterRegistration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*v1alpha1.ClusterRegistration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockClusterRegistrationControllerMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClusterRegistrationController)(nil).Update), arg0)
}

// UpdateStatus mocks base method.
func (m *MockClusterRegistrationController) UpdateStatus(arg0 *v1alpha1.ClusterRegistration) (*v1alpha1.ClusterRegistration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0)
	ret0, _ := ret[0].(*v1alpha1.ClusterRegistration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockClusterRegistrationControllerMockRecorder) UpdateStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockClusterRegistrationController)(nil).UpdateStatus), arg0)
}

// Updater mocks base method.
func (m *MockClusterRegistrationController) Updater() generic.Updater {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Updater")
	ret0, _ := ret[0].(generic.Updater)
	return ret0
}

// Updater indicates an expected call of Updater.
func (mr *MockClusterRegistrationControllerMockRecorder) Updater() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Updater", reflect.TypeOf((*MockClusterRegistrationController)(nil).Updater))
}

// Watch mocks base method.
func (m *MockClusterRegistrationController) Watch(arg0 string, arg1 v1.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockClusterRegistrationControllerMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockClusterRegistrationController)(nil).Watch), arg0, arg1)
}


func (m *MockClusterRegistrationController) WithImpersonation(arg0 rest.ImpersonationConfig) (generic.ClientInterface[*v1alpha1.ClusterRegistration, *v1alpha1.ClusterRegistrationList], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithImpersonation", arg0)
	ret0, _ := ret[0].(generic.ClientInterface[*v1alpha1.ClusterRegistration, *v1alpha1.ClusterRegistrationList])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockClusterRegistrationControllerMockRecorder) WithImpersonation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithImpersonation", reflect.TypeOf((*MockClusterRegistrationController)(nil).WithImpersonation), arg0)
}
