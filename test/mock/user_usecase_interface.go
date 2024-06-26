// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/satriowisnugroho/book-store/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// UserUsecaseInterface is an autogenerated mock type for the UserUsecaseInterface type
type UserUsecaseInterface struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, payload
func (_m *UserUsecaseInterface) CreateUser(ctx context.Context, payload *entity.RegisterPayload) (*entity.User, error) {
	ret := _m.Called(ctx, payload)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.RegisterPayload) (*entity.User, error)); ok {
		return rf(ctx, payload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.RegisterPayload) *entity.User); ok {
		r0 = rf(ctx, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.RegisterPayload) error); ok {
		r1 = rf(ctx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, payload
func (_m *UserUsecaseInterface) Login(ctx context.Context, payload *entity.LoginPayload) (*entity.LoginResponse, error) {
	ret := _m.Called(ctx, payload)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 *entity.LoginResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.LoginPayload) (*entity.LoginResponse, error)); ok {
		return rf(ctx, payload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.LoginPayload) *entity.LoginResponse); ok {
		r0 = rf(ctx, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.LoginResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.LoginPayload) error); ok {
		r1 = rf(ctx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserUsecaseInterface creates a new instance of UserUsecaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserUsecaseInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserUsecaseInterface {
	mock := &UserUsecaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
