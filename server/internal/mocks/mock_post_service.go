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
