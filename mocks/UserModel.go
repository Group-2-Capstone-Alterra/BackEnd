// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	user "PetPalApp/features/user"

	mock "github.com/stretchr/testify/mock"
)

// UserModel is an autogenerated mock type for the UserModel type
type UserModel struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *UserModel) Delete(id uint) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Insert provides a mock function with given fields: input
func (_m *UserModel) Insert(input user.Core) error {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(user.Core) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutById provides a mock function with given fields: id, input
func (_m *UserModel) PutById(id uint, input user.Core) error {
	ret := _m.Called(id, input)

	if len(ret) == 0 {
		panic("no return value specified for PutById")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, user.Core) error); ok {
		r0 = rf(id, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectByEmail provides a mock function with given fields: email
func (_m *UserModel) SelectByEmail(email string) (*user.Core, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for SelectByEmail")
	}

	var r0 *user.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*user.Core, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *user.Core); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectById provides a mock function with given fields: id
func (_m *UserModel) SelectById(id uint) (*user.Core, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for SelectById")
	}

	var r0 *user.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*user.Core, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *user.Core); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserModel creates a new instance of UserModel. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserModel(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserModel {
	mock := &UserModel{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
