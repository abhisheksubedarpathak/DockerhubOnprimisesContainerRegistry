// Code generated by mockery v2.43.2. DO NOT EDIT.

package scheduler

import (
	context "context"

	q "github.com/goharbor/harbor/src/lib/q"
	mock "github.com/stretchr/testify/mock"
)

// mockDAO is an autogenerated mock type for the DAO type
type mockDAO struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, query
func (_m *mockDAO) Count(ctx context.Context, query *q.Query) (int64, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for Count")
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

// Create provides a mock function with given fields: ctx, s
func (_m *mockDAO) Create(ctx context.Context, s *schedule) (int64, error) {
	ret := _m.Called(ctx, s)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *schedule) (int64, error)); ok {
		return rf(ctx, s)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *schedule) int64); ok {
		r0 = rf(ctx, s)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *schedule) error); ok {
		r1 = rf(ctx, s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *mockDAO) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *mockDAO) Get(ctx context.Context, id int64) (*schedule, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *schedule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*schedule, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *schedule); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schedule)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, query
func (_m *mockDAO) List(ctx context.Context, query *q.Query) ([]*schedule, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*schedule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) ([]*schedule, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) []*schedule); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*schedule)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, s, props
func (_m *mockDAO) Update(ctx context.Context, s *schedule, props ...string) error {
	_va := make([]interface{}, len(props))
	for _i := range props {
		_va[_i] = props[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, s)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *schedule, ...string) error); ok {
		r0 = rf(ctx, s, props...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRevision provides a mock function with given fields: ctx, id, revision
func (_m *mockDAO) UpdateRevision(ctx context.Context, id int64, revision int64) (int64, error) {
	ret := _m.Called(ctx, id, revision)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRevision")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) (int64, error)); ok {
		return rf(ctx, id, revision)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) int64); ok {
		r0 = rf(ctx, id, revision)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, id, revision)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newMockDAO creates a new instance of mockDAO. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockDAO(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockDAO {
	mock := &mockDAO{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
