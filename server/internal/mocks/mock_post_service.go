// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	context "context"
	dtos "gobooru/internal/dtos"

	mock "github.com/stretchr/testify/mock"
)

// MockPostService is an autogenerated mock type for the PostService type
type MockPostService struct {
	mock.Mock
}

type MockPostService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPostService) EXPECT() *MockPostService_Expecter {
	return &MockPostService_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, dto
func (_m *MockPostService) Create(ctx context.Context, dto dtos.CreatePostDTO) (dtos.CreatePostResponseDTO, error) {
	ret := _m.Called(ctx, dto)

	var r0 dtos.CreatePostResponseDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dtos.CreatePostDTO) (dtos.CreatePostResponseDTO, error)); ok {
		return rf(ctx, dto)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dtos.CreatePostDTO) dtos.CreatePostResponseDTO); ok {
		r0 = rf(ctx, dto)
	} else {
		r0 = ret.Get(0).(dtos.CreatePostResponseDTO)
	}

	if rf, ok := ret.Get(1).(func(context.Context, dtos.CreatePostDTO) error); ok {
		r1 = rf(ctx, dto)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPostService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockPostService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - dto dtos.CreatePostDTO
func (_e *MockPostService_Expecter) Create(ctx interface{}, dto interface{}) *MockPostService_Create_Call {
	return &MockPostService_Create_Call{Call: _e.mock.On("Create", ctx, dto)}
}

func (_c *MockPostService_Create_Call) Run(run func(ctx context.Context, dto dtos.CreatePostDTO)) *MockPostService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dtos.CreatePostDTO))
	})
	return _c
}

func (_c *MockPostService_Create_Call) Return(_a0 dtos.CreatePostResponseDTO, _a1 error) *MockPostService_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPostService_Create_Call) RunAndReturn(run func(context.Context, dtos.CreatePostDTO) (dtos.CreatePostResponseDTO, error)) *MockPostService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, dto
func (_m *MockPostService) Delete(ctx context.Context, dto dtos.DeletePostDTO) (dtos.DeletePostResponseDTO, error) {
	ret := _m.Called(ctx, dto)

	var r0 dtos.DeletePostResponseDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dtos.DeletePostDTO) (dtos.DeletePostResponseDTO, error)); ok {
		return rf(ctx, dto)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dtos.DeletePostDTO) dtos.DeletePostResponseDTO); ok {
		r0 = rf(ctx, dto)
	} else {
		r0 = ret.Get(0).(dtos.DeletePostResponseDTO)
	}

	if rf, ok := ret.Get(1).(func(context.Context, dtos.DeletePostDTO) error); ok {
		r1 = rf(ctx, dto)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPostService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockPostService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - dto dtos.DeletePostDTO
func (_e *MockPostService_Expecter) Delete(ctx interface{}, dto interface{}) *MockPostService_Delete_Call {
	return &MockPostService_Delete_Call{Call: _e.mock.On("Delete", ctx, dto)}
}

func (_c *MockPostService_Delete_Call) Run(run func(ctx context.Context, dto dtos.DeletePostDTO)) *MockPostService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dtos.DeletePostDTO))
	})
	return _c
}

func (_c *MockPostService_Delete_Call) Return(_a0 dtos.DeletePostResponseDTO, _a1 error) *MockPostService_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPostService_Delete_Call) RunAndReturn(run func(context.Context, dtos.DeletePostDTO) (dtos.DeletePostResponseDTO, error)) *MockPostService_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Fetch provides a mock function with given fields: ctx, dto
func (_m *MockPostService) Fetch(ctx context.Context, dto dtos.FetchPostDTO) (dtos.FetchPostResponseDTO, error) {
	ret := _m.Called(ctx, dto)

	var r0 dtos.FetchPostResponseDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dtos.FetchPostDTO) (dtos.FetchPostResponseDTO, error)); ok {
		return rf(ctx, dto)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dtos.FetchPostDTO) dtos.FetchPostResponseDTO); ok {
		r0 = rf(ctx, dto)
	} else {
		r0 = ret.Get(0).(dtos.FetchPostResponseDTO)
	}

	if rf, ok := ret.Get(1).(func(context.Context, dtos.FetchPostDTO) error); ok {
		r1 = rf(ctx, dto)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPostService_Fetch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fetch'
type MockPostService_Fetch_Call struct {
	*mock.Call
}

// Fetch is a helper method to define mock.On call
//   - ctx context.Context
//   - dto dtos.FetchPostDTO
func (_e *MockPostService_Expecter) Fetch(ctx interface{}, dto interface{}) *MockPostService_Fetch_Call {
	return &MockPostService_Fetch_Call{Call: _e.mock.On("Fetch", ctx, dto)}
}

func (_c *MockPostService_Fetch_Call) Run(run func(ctx context.Context, dto dtos.FetchPostDTO)) *MockPostService_Fetch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dtos.FetchPostDTO))
	})
	return _c
}

func (_c *MockPostService_Fetch_Call) Return(_a0 dtos.FetchPostResponseDTO, _a1 error) *MockPostService_Fetch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPostService_Fetch_Call) RunAndReturn(run func(context.Context, dtos.FetchPostDTO) (dtos.FetchPostResponseDTO, error)) *MockPostService_Fetch_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, dto
func (_m *MockPostService) List(ctx context.Context, dto dtos.ListPostDTO) (dtos.ListPostResponseDTO, error) {
	ret := _m.Called(ctx, dto)

	var r0 dtos.ListPostResponseDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dtos.ListPostDTO) (dtos.ListPostResponseDTO, error)); ok {
		return rf(ctx, dto)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dtos.ListPostDTO) dtos.ListPostResponseDTO); ok {
		r0 = rf(ctx, dto)
	} else {
		r0 = ret.Get(0).(dtos.ListPostResponseDTO)
	}

	if rf, ok := ret.Get(1).(func(context.Context, dtos.ListPostDTO) error); ok {
		r1 = rf(ctx, dto)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPostService_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockPostService_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - dto dtos.ListPostDTO
func (_e *MockPostService_Expecter) List(ctx interface{}, dto interface{}) *MockPostService_List_Call {
	return &MockPostService_List_Call{Call: _e.mock.On("List", ctx, dto)}
}

func (_c *MockPostService_List_Call) Run(run func(ctx context.Context, dto dtos.ListPostDTO)) *MockPostService_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dtos.ListPostDTO))
	})
	return _c
}

func (_c *MockPostService_List_Call) Return(_a0 dtos.ListPostResponseDTO, _a1 error) *MockPostService_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPostService_List_Call) RunAndReturn(run func(context.Context, dtos.ListPostDTO) (dtos.ListPostResponseDTO, error)) *MockPostService_List_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPostService creates a new instance of MockPostService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPostService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPostService {
	mock := &MockPostService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
