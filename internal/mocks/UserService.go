// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	exception "books-api/pkg/exception"
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "books-api/internal/model"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, req
func (_m *UserService) Login(ctx context.Context, req *model.CreateUserReq) (*model.LoginUserRes, *exception.Exception) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 *model.LoginUserRes
	var r1 *exception.Exception
	if rf, ok := ret.Get(0).(func(context.Context, *model.CreateUserReq) (*model.LoginUserRes, *exception.Exception)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.CreateUserReq) *model.LoginUserRes); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.LoginUserRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.CreateUserReq) *exception.Exception); ok {
		r1 = rf(ctx, req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*exception.Exception)
		}
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, req
func (_m *UserService) Register(ctx context.Context, req *model.CreateUserReq) (*model.CreateUserRes, *exception.Exception) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 *model.CreateUserRes
	var r1 *exception.Exception
	if rf, ok := ret.Get(0).(func(context.Context, *model.CreateUserReq) (*model.CreateUserRes, *exception.Exception)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.CreateUserReq) *model.CreateUserRes); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.CreateUserRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.CreateUserReq) *exception.Exception); ok {
		r1 = rf(ctx, req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*exception.Exception)
		}
	}

	return r0, r1
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
