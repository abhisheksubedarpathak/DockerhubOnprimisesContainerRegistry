// Code generated by mockery v2.43.2. DO NOT EDIT.

package scheduler

import (
	context "context"

	q "github.com/goharbor/harbor/src/lib/q"
	mock "github.com/stretchr/testify/mock"

	scheduler "github.com/goharbor/harbor/src/pkg/scheduler"
)

// Scheduler is an autogenerated mock type for the Scheduler type
type Scheduler struct {
	mock.Mock
}

// CountSchedules provides a mock function with given fields: ctx, query
func (_m *Scheduler) CountSchedules(ctx context.Context, query *q.Query) (int64, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for CountSchedules")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) (int64, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) int64); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSchedule provides a mock function with given fields: ctx, id
func (_m *Scheduler) GetSchedule(ctx context.Context, id int64) (*scheduler.Schedule, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetSchedule")
	}

	var r0 *scheduler.Schedule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*scheduler.Schedule, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *scheduler.Schedule); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scheduler.Schedule)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSchedules provides a mock function with given fields: ctx, query
func (_m *Scheduler) ListSchedules(ctx context.Context, query *q.Query) ([]*scheduler.Schedule, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for ListSchedules")
	}

	var r0 []*scheduler.Schedule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) ([]*scheduler.Schedule, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) []*scheduler.Schedule); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*scheduler.Schedule)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Schedule provides a mock function with given fields: ctx, vendorType, vendorID, cronType, cron, callbackFuncName, callbackFuncParams, extraAttrs
func (_m *Scheduler) Schedule(ctx context.Context, vendorType string, vendorID int64, cronType string, cron string, callbackFuncName string, callbackFuncParams interface{}, extraAttrs map[string]interface{}) (int64, error) {
	ret := _m.Called(ctx, vendorType, vendorID, cronType, cron, callbackFuncName, callbackFuncParams, extraAttrs)

	if len(ret) == 0 {
		panic("no return value specified for Schedule")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, string, string, string, interface{}, map[string]interface{}) (int64, error)); ok {
		return rf(ctx, vendorType, vendorID, cronType, cron, callbackFuncName, callbackFuncParams, extraAttrs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, string, string, string, interface{}, map[string]interface{}) int64); ok {
		r0 = rf(ctx, vendorType, vendorID, cronType, cron, callbackFuncName, callbackFuncParams, extraAttrs)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64, string, string, string, interface{}, map[string]interface{}) error); ok {
		r1 = rf(ctx, vendorType, vendorID, cronType, cron, callbackFuncName, callbackFuncParams, extraAttrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnScheduleByID provides a mock function with given fields: ctx, id
func (_m *Scheduler) UnScheduleByID(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for UnScheduleByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnScheduleByVendor provides a mock function with given fields: ctx, vendorType, vendorID
func (_m *Scheduler) UnScheduleByVendor(ctx context.Context, vendorType string, vendorID int64) error {
	ret := _m.Called(ctx, vendorType, vendorID)

	if len(ret) == 0 {
		panic("no return value specified for UnScheduleByVendor")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) error); ok {
		r0 = rf(ctx, vendorType, vendorID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewScheduler creates a new instance of Scheduler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewScheduler(t interface {
	mock.TestingT
	Cleanup(func())
}) *Scheduler {
	mock := &Scheduler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
