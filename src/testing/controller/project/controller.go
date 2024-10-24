// Code generated by mockery v2.46.2. DO NOT EDIT.

package project

import (
	context "context"

	commonmodels "github.com/goharbor/harbor/src/common/models"

	mock "github.com/stretchr/testify/mock"

	models "github.com/goharbor/harbor/src/pkg/project/models"

	project "github.com/goharbor/harbor/src/controller/project"

	q "github.com/goharbor/harbor/src/lib/q"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, query
func (_m *Controller) Count(ctx context.Context, query *q.Query) (int64, error) {
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

// Create provides a mock function with given fields: ctx, _a1
func (_m *Controller) Create(ctx context.Context, _a1 *models.Project) (int64, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Project) (int64, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.Project) int64); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.Project) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Controller) Delete(ctx context.Context, id int64) error {
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

// Exists provides a mock function with given fields: ctx, projectIDOrName
func (_m *Controller) Exists(ctx context.Context, projectIDOrName interface{}) (bool, error) {
	ret := _m.Called(ctx, projectIDOrName)

	if len(ret) == 0 {
		panic("no return value specified for Exists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) (bool, error)); ok {
		return rf(ctx, projectIDOrName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) bool); ok {
		r0 = rf(ctx, projectIDOrName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, projectIDOrName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, projectIDOrName, options
func (_m *Controller) Get(ctx context.Context, projectIDOrName interface{}, options ...project.Option) (*models.Project, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, projectIDOrName)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *models.Project
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...project.Option) (*models.Project, error)); ok {
		return rf(ctx, projectIDOrName, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...project.Option) *models.Project); ok {
		r0 = rf(ctx, projectIDOrName, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Project)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...project.Option) error); ok {
		r1 = rf(ctx, projectIDOrName, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: ctx, projectName, options
func (_m *Controller) GetByName(ctx context.Context, projectName string, options ...project.Option) (*models.Project, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, projectName)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *models.Project
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...project.Option) (*models.Project, error)); ok {
		return rf(ctx, projectName, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...project.Option) *models.Project); ok {
		r0 = rf(ctx, projectName, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Project)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...project.Option) error); ok {
		r1 = rf(ctx, projectName, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, query, options
func (_m *Controller) List(ctx context.Context, query *q.Query, options ...project.Option) ([]*models.Project, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*models.Project
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query, ...project.Option) ([]*models.Project, error)); ok {
		return rf(ctx, query, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query, ...project.Option) []*models.Project); ok {
		r0 = rf(ctx, query, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Project)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query, ...project.Option) error); ok {
		r1 = rf(ctx, query, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRoles provides a mock function with given fields: ctx, projectID, u
func (_m *Controller) ListRoles(ctx context.Context, projectID int64, u *commonmodels.User) ([]int, error) {
	ret := _m.Called(ctx, projectID, u)

	if len(ret) == 0 {
		panic("no return value specified for ListRoles")
	}

	var r0 []int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *commonmodels.User) ([]int, error)); ok {
		return rf(ctx, projectID, u)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, *commonmodels.User) []int); ok {
		r0 = rf(ctx, projectID, u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, *commonmodels.User) error); ok {
		r1 = rf(ctx, projectID, u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, _a1
func (_m *Controller) Update(ctx context.Context, _a1 *models.Project) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Project) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewController creates a new instance of Controller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewController(t interface {
	mock.TestingT
	Cleanup(func())
}) *Controller {
	mock := &Controller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
