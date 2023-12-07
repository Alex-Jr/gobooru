// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MockFFMPEGModule is an autogenerated mock type for the FFMPEGModule type
type MockFFMPEGModule struct {
	mock.Mock
}

type MockFFMPEGModule_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFFMPEGModule) EXPECT() *MockFFMPEGModule_Expecter {
	return &MockFFMPEGModule_Expecter{mock: &_m.Mock}
}

// GenerateGifThumb provides a mock function with given fields: src, dst
func (_m *MockFFMPEGModule) GenerateGifThumb(src string, dst string) error {
	ret := _m.Called(src, dst)

	if len(ret) == 0 {
		panic("no return value specified for GenerateGifThumb")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(src, dst)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFFMPEGModule_GenerateGifThumb_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateGifThumb'
type MockFFMPEGModule_GenerateGifThumb_Call struct {
	*mock.Call
}

// GenerateGifThumb is a helper method to define mock.On call
//   - src string
//   - dst string
func (_e *MockFFMPEGModule_Expecter) GenerateGifThumb(src interface{}, dst interface{}) *MockFFMPEGModule_GenerateGifThumb_Call {
	return &MockFFMPEGModule_GenerateGifThumb_Call{Call: _e.mock.On("GenerateGifThumb", src, dst)}
}

func (_c *MockFFMPEGModule_GenerateGifThumb_Call) Run(run func(src string, dst string)) *MockFFMPEGModule_GenerateGifThumb_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockFFMPEGModule_GenerateGifThumb_Call) Return(_a0 error) *MockFFMPEGModule_GenerateGifThumb_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFFMPEGModule_GenerateGifThumb_Call) RunAndReturn(run func(string, string) error) *MockFFMPEGModule_GenerateGifThumb_Call {
	_c.Call.Return(run)
	return _c
}

// GenerateImageThumb provides a mock function with given fields: src, dst
func (_m *MockFFMPEGModule) GenerateImageThumb(src string, dst string) error {
	ret := _m.Called(src, dst)

	if len(ret) == 0 {
		panic("no return value specified for GenerateImageThumb")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(src, dst)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFFMPEGModule_GenerateImageThumb_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateImageThumb'
type MockFFMPEGModule_GenerateImageThumb_Call struct {
	*mock.Call
}

// GenerateImageThumb is a helper method to define mock.On call
//   - src string
//   - dst string
func (_e *MockFFMPEGModule_Expecter) GenerateImageThumb(src interface{}, dst interface{}) *MockFFMPEGModule_GenerateImageThumb_Call {
	return &MockFFMPEGModule_GenerateImageThumb_Call{Call: _e.mock.On("GenerateImageThumb", src, dst)}
}

func (_c *MockFFMPEGModule_GenerateImageThumb_Call) Run(run func(src string, dst string)) *MockFFMPEGModule_GenerateImageThumb_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockFFMPEGModule_GenerateImageThumb_Call) Return(_a0 error) *MockFFMPEGModule_GenerateImageThumb_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFFMPEGModule_GenerateImageThumb_Call) RunAndReturn(run func(string, string) error) *MockFFMPEGModule_GenerateImageThumb_Call {
	_c.Call.Return(run)
	return _c
}

// GenerateVideoThumb provides a mock function with given fields: src, dst
func (_m *MockFFMPEGModule) GenerateVideoThumb(src string, dst string) error {
	ret := _m.Called(src, dst)

	if len(ret) == 0 {
		panic("no return value specified for GenerateVideoThumb")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(src, dst)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFFMPEGModule_GenerateVideoThumb_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateVideoThumb'
type MockFFMPEGModule_GenerateVideoThumb_Call struct {
	*mock.Call
}

// GenerateVideoThumb is a helper method to define mock.On call
//   - src string
//   - dst string
func (_e *MockFFMPEGModule_Expecter) GenerateVideoThumb(src interface{}, dst interface{}) *MockFFMPEGModule_GenerateVideoThumb_Call {
	return &MockFFMPEGModule_GenerateVideoThumb_Call{Call: _e.mock.On("GenerateVideoThumb", src, dst)}
}

func (_c *MockFFMPEGModule_GenerateVideoThumb_Call) Run(run func(src string, dst string)) *MockFFMPEGModule_GenerateVideoThumb_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockFFMPEGModule_GenerateVideoThumb_Call) Return(_a0 error) *MockFFMPEGModule_GenerateVideoThumb_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFFMPEGModule_GenerateVideoThumb_Call) RunAndReturn(run func(string, string) error) *MockFFMPEGModule_GenerateVideoThumb_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFFMPEGModule creates a new instance of MockFFMPEGModule. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFFMPEGModule(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFFMPEGModule {
	mock := &MockFFMPEGModule{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
