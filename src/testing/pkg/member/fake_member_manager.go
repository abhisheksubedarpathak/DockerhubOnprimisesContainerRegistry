// Code generated by mockery v2.43.2. DO NOT EDIT.

package member

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/goharbor/harbor/src/pkg/member/models"

	q "github.com/goharbor/harbor/src/lib/q"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// AddProjectMember provides a mock function with given fields: ctx, _a1
func (_m *Manager) AddProjectMember(ctx context.Context, _a1 models.Member) (int, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for AddProjectMember")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Member) (int, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.Member) int); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.Member) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, projectID, memberID
func (_m *Manager) Delete(ctx context.Context, projectID int64, memberID int) error {
	ret := _m.Called(ctx, projectID, memberID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int) error); ok {
		r0 = rf(ctx, projectID, memberID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMemberByProjectID provides a mock function with given fields: ctx, projectID
func (_m *Manager) DeleteMemberByProjectID(ctx context.Context, projectID int64) error {
	ret := _m.Called(ctx, projectID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMemberByProjectID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, projectID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMemberByUserID provides a mock function with given fields: ctx, uid
func (_m *Manager) DeleteMemberByUserID(ctx context.Context, uid int) error {
	ret := _m.Called(ctx, uid)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMemberByUserID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, projectID, memberID
func (_m *Manager) Get(ctx context.Context, projectID int64, memberID int) (*models.Member, error) {
	ret := _m.Called(ctx, projectID, memberID)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *models.Member
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int) (*models.Member, error)); ok {
		return rf(ctx, projectID, memberID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int) *models.Member); ok {
		r0 = rf(ctx, projectID, memberID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Member)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int) error); ok {
		r1 = rf(ctx, projectID, memberID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalOfProjectMembers provides a mock function with given fields: ctx, projectID, query, roles
func (_m *Manager) GetTotalOfProjectMembers(ctx context.Context, projectID int64, query *q.Query, roles ...int) (int, error) {
	_va := make([]interface{}, len(roles))
	for _i := range roles {
		_va[_i] = roles[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, projectID, query)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTotalOfProjectMembers")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *q.Query, ...int) (int, error)); ok {
		return rf(ctx, projectID, query, roles...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, *q.Query, ...int) int); ok {
		r0 = rf(ctx, projectID, query, roles...)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, *q.Query, ...int) error); ok {
		r1 = rf(ctx, projectID, query, roles...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, queryMember, query
func (_m *Manager) List(ctx context.Context, queryMember models.Member, query *q.Query) ([]*models.Member, error) {
	ret := _m.Called(ctx, queryMember, query)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*models.Member
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Member, *q.Query) ([]*models.Member, error)); ok {
		return rf(ctx, queryMember, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.Member, *q.Query) []*models.Member); ok {
		r0 = rf(ctx, queryMember, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Member)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.Member, *q.Query) error); ok {
		r1 = rf(ctx, queryMember, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRoles provides a mock function with given fields: ctx, user, projectID
func (_m *Manager) ListRoles(ctx context.Context, user *models.User, projectID int64) ([]int, error) {
	ret := _m.Called(ctx, user, projectID)

	if len(ret) == 0 {
		panic("no return value specified for ListRoles")
	}

	var r0 []int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.User, int64) ([]int, error)); ok {
		return rf(ctx, user, projectID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.User, int64) []int); ok {
		r0 = rf(ctx, user, projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.User, int64) error); ok {
		r1 = rf(ctx, user, projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchMemberByName provides a mock function with given fields: ctx, projectID, entityName
func (_m *Manager) SearchMemberByName(ctx context.Context, projectID int64, entityName string) ([]*models.Member, error) {
	ret := _m.Called(ctx, projectID, entityName)

	if len(ret) == 0 {
		panic("no return value specified for SearchMemberByName")
	}

	var r0 []*models.Member
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) ([]*models.Member, error)); ok {
		return rf(ctx, projectID, entityName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) []*models.Member); ok {
		r0 = rf(ctx, projectID, entityName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Member)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, projectID, entityName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRole provides a mock function with given fields: ctx, projectID, pmID, role
func (_m *Manager) UpdateRole(ctx context.Context, projectID int64, pmID int, role int) error {
	ret := _m.Called(ctx, projectID, pmID, role)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRole")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int, int) error); ok {
		r0 = rf(ctx, projectID, pmID, role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
