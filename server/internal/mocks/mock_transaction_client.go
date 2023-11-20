// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"

	sqlx "github.com/go-sqlx/sqlx"
)

// MockTransactionClient is an autogenerated mock type for the TransactionClient type
type MockTransactionClient struct {
	mock.Mock
}

type MockTransactionClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTransactionClient) EXPECT() *MockTransactionClient_Expecter {
	return &MockTransactionClient_Expecter{mock: &_m.Mock}
}

// Commit provides a mock function with given fields:
func (_m *MockTransactionClient) Commit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTransactionClient_Commit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Commit'
type MockTransactionClient_Commit_Call struct {
	*mock.Call
}

// Commit is a helper method to define mock.On call
func (_e *MockTransactionClient_Expecter) Commit() *MockTransactionClient_Commit_Call {
	return &MockTransactionClient_Commit_Call{Call: _e.mock.On("Commit")}
}

func (_c *MockTransactionClient_Commit_Call) Run(run func()) *MockTransactionClient_Commit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTransactionClient_Commit_Call) Return(_a0 error) *MockTransactionClient_Commit_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTransactionClient_Commit_Call) RunAndReturn(run func() error) *MockTransactionClient_Commit_Call {
	_c.Call.Return(run)
	return _c
}

// ExecContext provides a mock function with given fields: ctx, query, args
func (_m *MockTransactionClient) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 sql.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (sql.Result, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) sql.Result); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTransactionClient_ExecContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecContext'
type MockTransactionClient_ExecContext_Call struct {
	*mock.Call
}

// ExecContext is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *MockTransactionClient_Expecter) ExecContext(ctx interface{}, query interface{}, args ...interface{}) *MockTransactionClient_ExecContext_Call {
	return &MockTransactionClient_ExecContext_Call{Call: _e.mock.On("ExecContext",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *MockTransactionClient_ExecContext_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *MockTransactionClient_ExecContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockTransactionClient_ExecContext_Call) Return(_a0 sql.Result, _a1 error) *MockTransactionClient_ExecContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTransactionClient_ExecContext_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (sql.Result, error)) *MockTransactionClient_ExecContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetContext provides a mock function with given fields: ctx, dest, query, args
func (_m *MockTransactionClient) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, dest, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, dest, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTransactionClient_GetContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetContext'
type MockTransactionClient_GetContext_Call struct {
	*mock.Call
}

// GetContext is a helper method to define mock.On call
//   - ctx context.Context
//   - dest interface{}
//   - query string
//   - args ...interface{}
func (_e *MockTransactionClient_Expecter) GetContext(ctx interface{}, dest interface{}, query interface{}, args ...interface{}) *MockTransactionClient_GetContext_Call {
	return &MockTransactionClient_GetContext_Call{Call: _e.mock.On("GetContext",
		append([]interface{}{ctx, dest, query}, args...)...)}
}

func (_c *MockTransactionClient_GetContext_Call) Run(run func(ctx context.Context, dest interface{}, query string, args ...interface{})) *MockTransactionClient_GetContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(interface{}), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockTransactionClient_GetContext_Call) Return(_a0 error) *MockTransactionClient_GetContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTransactionClient_GetContext_Call) RunAndReturn(run func(context.Context, interface{}, string, ...interface{}) error) *MockTransactionClient_GetContext_Call {
	_c.Call.Return(run)
	return _c
}

// NamedExecContext provides a mock function with given fields: ctx, query, arg
func (_m *MockTransactionClient) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	ret := _m.Called(ctx, query, arg)

	var r0 sql.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) (sql.Result, error)); ok {
		return rf(ctx, query, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) sql.Result); ok {
		r0 = rf(ctx, query, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}) error); ok {
		r1 = rf(ctx, query, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTransactionClient_NamedExecContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NamedExecContext'
type MockTransactionClient_NamedExecContext_Call struct {
	*mock.Call
}

// NamedExecContext is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - arg interface{}
func (_e *MockTransactionClient_Expecter) NamedExecContext(ctx interface{}, query interface{}, arg interface{}) *MockTransactionClient_NamedExecContext_Call {
	return &MockTransactionClient_NamedExecContext_Call{Call: _e.mock.On("NamedExecContext", ctx, query, arg)}
}

func (_c *MockTransactionClient_NamedExecContext_Call) Run(run func(ctx context.Context, query string, arg interface{})) *MockTransactionClient_NamedExecContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}))
	})
	return _c
}

func (_c *MockTransactionClient_NamedExecContext_Call) Return(_a0 sql.Result, _a1 error) *MockTransactionClient_NamedExecContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTransactionClient_NamedExecContext_Call) RunAndReturn(run func(context.Context, string, interface{}) (sql.Result, error)) *MockTransactionClient_NamedExecContext_Call {
	_c.Call.Return(run)
	return _c
}

// NamedQueryContext provides a mock function with given fields: ctx, query, arg
func (_m *MockTransactionClient) NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	ret := _m.Called(ctx, query, arg)

	var r0 *sqlx.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) (*sqlx.Rows, error)); ok {
		return rf(ctx, query, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) *sqlx.Rows); ok {
		r0 = rf(ctx, query, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}) error); ok {
		r1 = rf(ctx, query, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTransactionClient_NamedQueryContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NamedQueryContext'
type MockTransactionClient_NamedQueryContext_Call struct {
	*mock.Call
}

// NamedQueryContext is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - arg interface{}
func (_e *MockTransactionClient_Expecter) NamedQueryContext(ctx interface{}, query interface{}, arg interface{}) *MockTransactionClient_NamedQueryContext_Call {
	return &MockTransactionClient_NamedQueryContext_Call{Call: _e.mock.On("NamedQueryContext", ctx, query, arg)}
}

func (_c *MockTransactionClient_NamedQueryContext_Call) Run(run func(ctx context.Context, query string, arg interface{})) *MockTransactionClient_NamedQueryContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}))
	})
	return _c
}

func (_c *MockTransactionClient_NamedQueryContext_Call) Return(_a0 *sqlx.Rows, _a1 error) *MockTransactionClient_NamedQueryContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTransactionClient_NamedQueryContext_Call) RunAndReturn(run func(context.Context, string, interface{}) (*sqlx.Rows, error)) *MockTransactionClient_NamedQueryContext_Call {
	_c.Call.Return(run)
	return _c
}

// Rollback provides a mock function with given fields:
func (_m *MockTransactionClient) Rollback() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTransactionClient_Rollback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rollback'
type MockTransactionClient_Rollback_Call struct {
	*mock.Call
}

// Rollback is a helper method to define mock.On call
func (_e *MockTransactionClient_Expecter) Rollback() *MockTransactionClient_Rollback_Call {
	return &MockTransactionClient_Rollback_Call{Call: _e.mock.On("Rollback")}
}

func (_c *MockTransactionClient_Rollback_Call) Run(run func()) *MockTransactionClient_Rollback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTransactionClient_Rollback_Call) Return(_a0 error) *MockTransactionClient_Rollback_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTransactionClient_Rollback_Call) RunAndReturn(run func() error) *MockTransactionClient_Rollback_Call {
	_c.Call.Return(run)
	return _c
}

// SelectContext provides a mock function with given fields: ctx, dest, query, args
func (_m *MockTransactionClient) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, dest, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, dest, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTransactionClient_SelectContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectContext'
type MockTransactionClient_SelectContext_Call struct {
	*mock.Call
}

// SelectContext is a helper method to define mock.On call
//   - ctx context.Context
//   - dest interface{}
//   - query string
//   - args ...interface{}
func (_e *MockTransactionClient_Expecter) SelectContext(ctx interface{}, dest interface{}, query interface{}, args ...interface{}) *MockTransactionClient_SelectContext_Call {
	return &MockTransactionClient_SelectContext_Call{Call: _e.mock.On("SelectContext",
		append([]interface{}{ctx, dest, query}, args...)...)}
}

func (_c *MockTransactionClient_SelectContext_Call) Run(run func(ctx context.Context, dest interface{}, query string, args ...interface{})) *MockTransactionClient_SelectContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(interface{}), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockTransactionClient_SelectContext_Call) Return(_a0 error) *MockTransactionClient_SelectContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTransactionClient_SelectContext_Call) RunAndReturn(run func(context.Context, interface{}, string, ...interface{}) error) *MockTransactionClient_SelectContext_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTransactionClient creates a new instance of MockTransactionClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTransactionClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTransactionClient {
	mock := &MockTransactionClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
