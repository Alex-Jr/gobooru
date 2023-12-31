// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// MockPoolController is an autogenerated mock type for the PoolController type
type MockPoolController struct {
	mock.Mock
}

type MockPoolController_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPoolController) EXPECT() *MockPoolController_Expecter {
	return &MockPoolController_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: c
func (_m *MockPoolController) Create(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPoolController_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockPoolController_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - c echo.Context
func (_e *MockPoolController_Expecter) Create(c interface{}) *MockPoolController_Create_Call {
	return &MockPoolController_Create_Call{Call: _e.mock.On("Create", c)}
}

func (_c *MockPoolController_Create_Call) Run(run func(c echo.Context)) *MockPoolController_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *MockPoolController_Create_Call) Return(_a0 error) *MockPoolController_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPoolController_Create_Call) RunAndReturn(run func(echo.Context) error) *MockPoolController_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: c
func (_m *MockPoolController) Delete(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPoolController_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockPoolController_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - c echo.Context
func (_e *MockPoolController_Expecter) Delete(c interface{}) *MockPoolController_Delete_Call {
	return &MockPoolController_Delete_Call{Call: _e.mock.On("Delete", c)}
}

func (_c *MockPoolController_Delete_Call) Run(run func(c echo.Context)) *MockPoolController_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *MockPoolController_Delete_Call) Return(_a0 error) *MockPoolController_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPoolController_Delete_Call) RunAndReturn(run func(echo.Context) error) *MockPoolController_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Fetch provides a mock function with given fields: c
func (_m *MockPoolController) Fetch(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPoolController_Fetch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fetch'
type MockPoolController_Fetch_Call struct {
	*mock.Call
}

// Fetch is a helper method to define mock.On call
//   - c echo.Context
func (_e *MockPoolController_Expecter) Fetch(c interface{}) *MockPoolController_Fetch_Call {
	return &MockPoolController_Fetch_Call{Call: _e.mock.On("Fetch", c)}
}

func (_c *MockPoolController_Fetch_Call) Run(run func(c echo.Context)) *MockPoolController_Fetch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *MockPoolController_Fetch_Call) Return(_a0 error) *MockPoolController_Fetch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPoolController_Fetch_Call) RunAndReturn(run func(echo.Context) error) *MockPoolController_Fetch_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: c
func (_m *MockPoolController) List(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPoolController_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockPoolController_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - c echo.Context
func (_e *MockPoolController_Expecter) List(c interface{}) *MockPoolController_List_Call {
	return &MockPoolController_List_Call{Call: _e.mock.On("List", c)}
}

func (_c *MockPoolController_List_Call) Run(run func(c echo.Context)) *MockPoolController_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *MockPoolController_List_Call) Return(_a0 error) *MockPoolController_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPoolController_List_Call) RunAndReturn(run func(echo.Context) error) *MockPoolController_List_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: c
func (_m *MockPoolController) Update(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPoolController_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockPoolController_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - c echo.Context
func (_e *MockPoolController_Expecter) Update(c interface{}) *MockPoolController_Update_Call {
	return &MockPoolController_Update_Call{Call: _e.mock.On("Update", c)}
}

func (_c *MockPoolController_Update_Call) Run(run func(c echo.Context)) *MockPoolController_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *MockPoolController_Update_Call) Return(_a0 error) *MockPoolController_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPoolController_Update_Call) RunAndReturn(run func(echo.Context) error) *MockPoolController_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPoolController creates a new instance of MockPoolController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPoolController(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPoolController {
	mock := &MockPoolController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
