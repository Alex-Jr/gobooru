// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"
	database "gobooru/internal/database"

	mock "github.com/stretchr/testify/mock"
)

// MockTagAliasQuery is an autogenerated mock type for the TagAliasQuery type
type MockTagAliasQuery struct {
	mock.Mock
}

type MockTagAliasQuery_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTagAliasQuery) EXPECT() *MockTagAliasQuery_Expecter {
	return &MockTagAliasQuery_Expecter{mock: &_m.Mock}
}

// ResolveAlias provides a mock function with given fields: ctx, db, tags
func (_m *MockTagAliasQuery) ResolveAlias(ctx context.Context, db database.DBClient, tags []string) error {
	ret := _m.Called(ctx, db, tags)

	if len(ret) == 0 {
		panic("no return value specified for ResolveAlias")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.DBClient, []string) error); ok {
		r0 = rf(ctx, db, tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTagAliasQuery_ResolveAlias_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResolveAlias'
type MockTagAliasQuery_ResolveAlias_Call struct {
	*mock.Call
}

// ResolveAlias is a helper method to define mock.On call
//   - ctx context.Context
//   - db database.DBClient
//   - tags []string
func (_e *MockTagAliasQuery_Expecter) ResolveAlias(ctx interface{}, db interface{}, tags interface{}) *MockTagAliasQuery_ResolveAlias_Call {
	return &MockTagAliasQuery_ResolveAlias_Call{Call: _e.mock.On("ResolveAlias", ctx, db, tags)}
}

func (_c *MockTagAliasQuery_ResolveAlias_Call) Run(run func(ctx context.Context, db database.DBClient, tags []string)) *MockTagAliasQuery_ResolveAlias_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.DBClient), args[2].([]string))
	})
	return _c
}

func (_c *MockTagAliasQuery_ResolveAlias_Call) Return(_a0 error) *MockTagAliasQuery_ResolveAlias_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTagAliasQuery_ResolveAlias_Call) RunAndReturn(run func(context.Context, database.DBClient, []string) error) *MockTagAliasQuery_ResolveAlias_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTagAliasQuery creates a new instance of MockTagAliasQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTagAliasQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTagAliasQuery {
	mock := &MockTagAliasQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
