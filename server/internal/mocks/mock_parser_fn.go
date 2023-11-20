// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MockParserFn is an autogenerated mock type for the ParserFn type
type MockParserFn struct {
	mock.Mock
}

type MockParserFn_Expecter struct {
	mock *mock.Mock
}

func (_m *MockParserFn) EXPECT() *MockParserFn_Expecter {
	return &MockParserFn_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *MockParserFn) Execute(_a0 interface{}) interface{} {
	ret := _m.Called(_a0)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// MockParserFn_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockParserFn_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 interface{}
func (_e *MockParserFn_Expecter) Execute(_a0 interface{}) *MockParserFn_Execute_Call {
	return &MockParserFn_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *MockParserFn_Execute_Call) Run(run func(_a0 interface{})) *MockParserFn_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockParserFn_Execute_Call) Return(_a0 interface{}) *MockParserFn_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockParserFn_Execute_Call) RunAndReturn(run func(interface{}) interface{}) *MockParserFn_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockParserFn creates a new instance of MockParserFn. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockParserFn(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockParserFn {
	mock := &MockParserFn{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
