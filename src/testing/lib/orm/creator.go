// Code generated by mockery v2.14.0. DO NOT EDIT.

package orm

import (
	orm "github.com/beego/beego/orm"
	mock "github.com/stretchr/testify/mock"
)

// Creator is an autogenerated mock type for the Creator type
type Creator struct {
	mock.Mock
}

// Create provides a mock function with given fields:
func (_m *Creator) Create() orm.Ormer {
	ret := _m.Called()

	var r0 orm.Ormer
	if rf, ok := ret.Get(0).(func() orm.Ormer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Ormer)
		}
	}

	return r0
}

type mockConstructorTestingTNewCreator interface {
	mock.TestingT
	Cleanup(func())
}

// NewCreator creates a new instance of Creator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCreator(t mockConstructorTestingTNewCreator) *Creator {
	mock := &Creator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
