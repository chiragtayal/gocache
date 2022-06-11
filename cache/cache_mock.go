// Code generated by MockGen. DO NOT EDIT.
// Source: cache.go

// Package cache is a generated GoMock package.
package cache

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockCache) Get(key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCacheMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCache)(nil).Get), key)
}

// Put mocks base method.
func (m *MockCache) Put(key, value string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Put", key, value)
}

// Put indicates an expected call of Put.
func (mr *MockCacheMockRecorder) Put(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockCache)(nil).Put), key, value)
}

// MockInitializer is a mock of Initializer interface.
type MockInitializer struct {
	ctrl     *gomock.Controller
	recorder *MockInitializerMockRecorder
}

// MockInitializerMockRecorder is the mock recorder for MockInitializer.
type MockInitializerMockRecorder struct {
	mock *MockInitializer
}

// NewMockInitializer creates a new mock instance.
func NewMockInitializer(ctrl *gomock.Controller) *MockInitializer {
	mock := &MockInitializer{ctrl: ctrl}
	mock.recorder = &MockInitializerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInitializer) EXPECT() *MockInitializerMockRecorder {
	return m.recorder
}

// Initialize mocks base method.
func (m *MockInitializer) Initialize(size int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Initialize", size)
}

// Initialize indicates an expected call of Initialize.
func (mr *MockInitializerMockRecorder) Initialize(size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockInitializer)(nil).Initialize), size)
}

// MockComposer is a mock of Composer interface.
type MockComposer struct {
	ctrl     *gomock.Controller
	recorder *MockComposerMockRecorder
}

// MockComposerMockRecorder is the mock recorder for MockComposer.
type MockComposerMockRecorder struct {
	mock *MockComposer
}

// NewMockComposer creates a new mock instance.
func NewMockComposer(ctrl *gomock.Controller) *MockComposer {
	mock := &MockComposer{ctrl: ctrl}
	mock.recorder = &MockComposerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComposer) EXPECT() *MockComposerMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockComposer) Get(key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockComposerMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockComposer)(nil).Get), key)
}

// Initialize mocks base method.
func (m *MockComposer) Initialize(size int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Initialize", size)
}

// Initialize indicates an expected call of Initialize.
func (mr *MockComposerMockRecorder) Initialize(size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockComposer)(nil).Initialize), size)
}

// Put mocks base method.
func (m *MockComposer) Put(key, value string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Put", key, value)
}

// Put indicates an expected call of Put.
func (mr *MockComposerMockRecorder) Put(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockComposer)(nil).Put), key, value)
}
