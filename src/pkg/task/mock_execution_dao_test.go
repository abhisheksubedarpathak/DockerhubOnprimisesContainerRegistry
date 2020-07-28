// Code generated by mockery v2.1.0. DO NOT EDIT.

package task

import (
	context "context"

	dao "github.com/goharbor/harbor/src/pkg/task/dao"
	mock "github.com/stretchr/testify/mock"

	q "github.com/goharbor/harbor/src/lib/q"
)

// mockExecutionDAO is an autogenerated mock type for the ExecutionDAO type
type mockExecutionDAO struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, query
func (_m *mockExecutionDAO) Count(ctx context.Context, query *q.Query) (int64, error) {
	ret := _m.Called(ctx, query)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) int64); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, execution
func (_m *mockExecutionDAO) Create(ctx context.Context, execution *dao.Execution) (int64, error) {
	ret := _m.Called(ctx, execution)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *dao.Execution) int64); ok {
		r0 = rf(ctx, execution)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dao.Execution) error); ok {
		r1 = rf(ctx, execution)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *mockExecutionDAO) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *mockExecutionDAO) Get(ctx context.Context, id int64) (*dao.Execution, error) {
	ret := _m.Called(ctx, id)

	var r0 *dao.Execution
	if rf, ok := ret.Get(0).(func(context.Context, int64) *dao.Execution); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dao.Execution)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, query
func (_m *mockExecutionDAO) List(ctx context.Context, query *q.Query) ([]*dao.Execution, error) {
	ret := _m.Called(ctx, query)

	var r0 []*dao.Execution
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) []*dao.Execution); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dao.Execution)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, execution, props
func (_m *mockExecutionDAO) Update(ctx context.Context, execution *dao.Execution, props ...string) error {
	_va := make([]interface{}, len(props))
	for _i := range props {
		_va[_i] = props[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, execution)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dao.Execution, ...string) error); ok {
		r0 = rf(ctx, execution, props...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
