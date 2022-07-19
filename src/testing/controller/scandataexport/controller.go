// Code generated by mockery v2.12.3. DO NOT EDIT.

package scandataexport

import (
	context "context"

	export "github.com/goharbor/harbor/src/pkg/scan/export"
	mock "github.com/stretchr/testify/mock"

	task "github.com/goharbor/harbor/src/pkg/task"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// DeleteExecution provides a mock function with given fields: ctx, executionID
func (_m *Controller) DeleteExecution(ctx context.Context, executionID int64) error {
	ret := _m.Called(ctx, executionID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, executionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetExecution provides a mock function with given fields: ctx, executionID
func (_m *Controller) GetExecution(ctx context.Context, executionID int64) (*export.Execution, error) {
	ret := _m.Called(ctx, executionID)

	var r0 *export.Execution
	if rf, ok := ret.Get(0).(func(context.Context, int64) *export.Execution); ok {
		r0 = rf(ctx, executionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*export.Execution)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, executionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTask provides a mock function with given fields: ctx, executionID
func (_m *Controller) GetTask(ctx context.Context, executionID int64) (*task.Task, error) {
	ret := _m.Called(ctx, executionID)

	var r0 *task.Task
	if rf, ok := ret.Get(0).(func(context.Context, int64) *task.Task); ok {
		r0 = rf(ctx, executionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*task.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, executionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListExecutions provides a mock function with given fields: ctx, userName
func (_m *Controller) ListExecutions(ctx context.Context, userName string) ([]*export.Execution, error) {
	ret := _m.Called(ctx, userName)

	var r0 []*export.Execution
	if rf, ok := ret.Get(0).(func(context.Context, string) []*export.Execution); ok {
		r0 = rf(ctx, userName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*export.Execution)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields: ctx, criteria
func (_m *Controller) Start(ctx context.Context, criteria export.Request) (int64, error) {
	ret := _m.Called(ctx, criteria)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, export.Request) int64); ok {
		r0 = rf(ctx, criteria)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, export.Request) error); ok {
		r1 = rf(ctx, criteria)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewControllerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewController creates a new instance of Controller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewController(t NewControllerT) *Controller {
	mock := &Controller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
