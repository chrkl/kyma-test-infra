// Code generated by mockery v2.33.1. DO NOT EDIT.

package automock

import (
	context "context"

	storage "github.com/kyma-project/test-infra/development/tools/pkg/gcscleaner/storage"
	mock "github.com/stretchr/testify/mock"
)

// BucketHandle is an autogenerated mock type for the BucketHandle type
type BucketHandle struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx
func (_m *BucketHandle) Delete(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Object provides a mock function with given fields: name
func (_m *BucketHandle) Object(name string) storage.ObjectHandle {
	ret := _m.Called(name)

	var r0 storage.ObjectHandle
	if rf, ok := ret.Get(0).(func(string) storage.ObjectHandle); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(storage.ObjectHandle)
		}
	}

	return r0
}

// Objects provides a mock function with given fields: ctx, q
func (_m *BucketHandle) Objects(ctx context.Context, q storage.Query) storage.ObjectIterator {
	ret := _m.Called(ctx, q)

	var r0 storage.ObjectIterator
	if rf, ok := ret.Get(0).(func(context.Context, storage.Query) storage.ObjectIterator); ok {
		r0 = rf(ctx, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(storage.ObjectIterator)
		}
	}

	return r0
}

// NewBucketHandle creates a new instance of BucketHandle. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBucketHandle(t interface {
	mock.TestingT
	Cleanup(func())
}) *BucketHandle {
	mock := &BucketHandle{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
