// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// MockPostController is an autogenerated mock type for the PostController type
type MockPostController struct {
	mock.Mock
}

type MockPostController_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPostController) EXPECT() *MockPostController_Expecter {
	return &MockPostController_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: c
func (_m *MockPostController) Create(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPostController_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockPostController_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - c echo.Context
func (_e *MockPostController_Expecter) Create(c interface{}) *MockPostController_Create_Call {
	return &MockPostController_Create_Call{Call: _e.mock.On("Create", c)}
}

func (_c *MockPostController_Create_Call) Run(run func(c echo.Context)) *MockPostController_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *MockPostController_Create_Call) Return(_a0 error) *MockPostController_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPostController_Create_Call) RunAndReturn(run func(echo.Context) error) *MockPostController_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: c
func (_m *MockPostController) Delete(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPostController_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockPostController_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - c echo.Context
func (_e *MockPostController_Expecter) Delete(c interface{}) *MockPostController_Delete_Call {
	return &MockPostController_Delete_Call{Call: _e.mock.On("Delete", c)}
}

func (_c *MockPostController_Delete_Call) Run(run func(c echo.Context)) *MockPostController_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *MockPostController_Delete_Call) Return(_a0 error) *MockPostController_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPostController_Delete_Call) RunAndReturn(run func(echo.Context) error) *MockPostController_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Fetch provides a mock function with given fields: c
func (_m *MockPostController) Fetch(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPostController_Fetch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fetch'
type MockPostController_Fetch_Call struct {
	*mock.Call
}

// Fetch is a helper method to define mock.On call
//   - c echo.Context
func (_e *MockPostController_Expecter) Fetch(c interface{}) *MockPostController_Fetch_Call {
	return &MockPostController_Fetch_Call{Call: _e.mock.On("Fetch", c)}
}

func (_c *MockPostController_Fetch_Call) Run(run func(c echo.Context)) *MockPostController_Fetch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *MockPostController_Fetch_Call) Return(_a0 error) *MockPostController_Fetch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPostController_Fetch_Call) RunAndReturn(run func(echo.Context) error) *MockPostController_Fetch_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPostController creates a new instance of MockPostController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPostController(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPostController {
	mock := &MockPostController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
