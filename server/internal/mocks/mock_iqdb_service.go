// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	models "gobooru/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// MockIQDBService is an autogenerated mock type for the IQDBService type
type MockIQDBService struct {
	mock.Mock
}

type MockIQDBService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIQDBService) EXPECT() *MockIQDBService_Expecter {
	return &MockIQDBService_Expecter{mock: &_m.Mock}
}

// HandlePost provides a mock function with given fields: post
func (_m *MockIQDBService) HandlePost(post models.Post) ([]models.PostRelation, error) {
	ret := _m.Called(post)

	if len(ret) == 0 {
		panic("no return value specified for HandlePost")
	}

	var r0 []models.PostRelation
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Post) ([]models.PostRelation, error)); ok {
		return rf(post)
	}
	if rf, ok := ret.Get(0).(func(models.Post) []models.PostRelation); ok {
		r0 = rf(post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.PostRelation)
		}
	}

	if rf, ok := ret.Get(1).(func(models.Post) error); ok {
		r1 = rf(post)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIQDBService_HandlePost_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandlePost'
type MockIQDBService_HandlePost_Call struct {
	*mock.Call
}

// HandlePost is a helper method to define mock.On call
//   - post models.Post
func (_e *MockIQDBService_Expecter) HandlePost(post interface{}) *MockIQDBService_HandlePost_Call {
	return &MockIQDBService_HandlePost_Call{Call: _e.mock.On("HandlePost", post)}
}

func (_c *MockIQDBService_HandlePost_Call) Run(run func(post models.Post)) *MockIQDBService_HandlePost_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.Post))
	})
	return _c
}

func (_c *MockIQDBService_HandlePost_Call) Return(_a0 []models.PostRelation, _a1 error) *MockIQDBService_HandlePost_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIQDBService_HandlePost_Call) RunAndReturn(run func(models.Post) ([]models.PostRelation, error)) *MockIQDBService_HandlePost_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIQDBService creates a new instance of MockIQDBService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIQDBService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIQDBService {
	mock := &MockIQDBService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
