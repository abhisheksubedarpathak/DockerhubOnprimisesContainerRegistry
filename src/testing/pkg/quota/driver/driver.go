// Code generated by mockery v1.0.0. DO NOT EDIT.

package driver

import (
	context "context"

	driver "github.com/goharbor/harbor/src/pkg/quota/driver"
	mock "github.com/stretchr/testify/mock"

	types "github.com/goharbor/harbor/src/pkg/quota/types"
)

// Driver is an autogenerated mock type for the Driver type
type Driver struct {
	mock.Mock
}

// CalculateUsage provides a mock function with given fields: ctx, key
func (_m *Driver) CalculateUsage(ctx context.Context, key string) (types.ResourceList, error) {
	ret := _m.Called(ctx, key)

	var r0 types.ResourceList
	if rf, ok := ret.Get(0).(func(context.Context, string) types.ResourceList); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.ResourceList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Enabled provides a mock function with given fields: ctx, key
func (_m *Driver) Enabled(ctx context.Context, key string) (bool, error) {
	ret := _m.Called(ctx, key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HardLimits provides a mock function with given fields: ctx
func (_m *Driver) HardLimits(ctx context.Context) types.ResourceList {
	ret := _m.Called(ctx)

	var r0 types.ResourceList
	if rf, ok := ret.Get(0).(func(context.Context) types.ResourceList); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.ResourceList)
		}
	}

	return r0
}

// Load provides a mock function with given fields: ctx, key
func (_m *Driver) Load(ctx context.Context, key string) (driver.RefObject, error) {
	ret := _m.Called(ctx, key)

	var r0 driver.RefObject
	if rf, ok := ret.Get(0).(func(context.Context, string) driver.RefObject); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.RefObject)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validate provides a mock function with given fields: hardLimits
func (_m *Driver) Validate(hardLimits types.ResourceList) error {
	ret := _m.Called(hardLimits)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.ResourceList) error); ok {
		r0 = rf(hardLimits)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
