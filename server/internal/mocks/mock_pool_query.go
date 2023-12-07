// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"
	database "gobooru/internal/database"

	mock "github.com/stretchr/testify/mock"

	models "gobooru/internal/models"
)

// MockPoolQuery is an autogenerated mock type for the PoolQuery type
type MockPoolQuery struct {
	mock.Mock
}

type MockPoolQuery_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPoolQuery) EXPECT() *MockPoolQuery_Expecter {
	return &MockPoolQuery_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, db, pool
func (_m *MockPoolQuery) Create(ctx context.Context, db database.DBClient, pool *models.Pool) error {
	ret := _m.Called(ctx, db, pool)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.DBClient, *models.Pool) error); ok {
		r0 = rf(ctx, db, pool)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPoolQuery_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockPoolQuery_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - db database.DBClient
//   - pool *models.Pool
func (_e *MockPoolQuery_Expecter) Create(ctx interface{}, db interface{}, pool interface{}) *MockPoolQuery_Create_Call {
	return &MockPoolQuery_Create_Call{Call: _e.mock.On("Create", ctx, db, pool)}
}

func (_c *MockPoolQuery_Create_Call) Run(run func(ctx context.Context, db database.DBClient, pool *models.Pool)) *MockPoolQuery_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.DBClient), args[2].(*models.Pool))
	})
	return _c
}

func (_c *MockPoolQuery_Create_Call) Return(_a0 error) *MockPoolQuery_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPoolQuery_Create_Call) RunAndReturn(run func(context.Context, database.DBClient, *models.Pool) error) *MockPoolQuery_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, db, pool
func (_m *MockPoolQuery) Delete(ctx context.Context, db database.DBClient, pool *models.Pool) error {
	ret := _m.Called(ctx, db, pool)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.DBClient, *models.Pool) error); ok {
		r0 = rf(ctx, db, pool)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPoolQuery_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockPoolQuery_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - db database.DBClient
//   - pool *models.Pool
func (_e *MockPoolQuery_Expecter) Delete(ctx interface{}, db interface{}, pool interface{}) *MockPoolQuery_Delete_Call {
	return &MockPoolQuery_Delete_Call{Call: _e.mock.On("Delete", ctx, db, pool)}
}

func (_c *MockPoolQuery_Delete_Call) Run(run func(ctx context.Context, db database.DBClient, pool *models.Pool)) *MockPoolQuery_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.DBClient), args[2].(*models.Pool))
	})
	return _c
}

func (_c *MockPoolQuery_Delete_Call) Return(_a0 error) *MockPoolQuery_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPoolQuery_Delete_Call) RunAndReturn(run func(context.Context, database.DBClient, *models.Pool) error) *MockPoolQuery_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetFull provides a mock function with given fields: ctx, db, pool
func (_m *MockPoolQuery) GetFull(ctx context.Context, db database.DBClient, pool *models.Pool) error {
	ret := _m.Called(ctx, db, pool)

	if len(ret) == 0 {
		panic("no return value specified for GetFull")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.DBClient, *models.Pool) error); ok {
		r0 = rf(ctx, db, pool)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPoolQuery_GetFull_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFull'
type MockPoolQuery_GetFull_Call struct {
	*mock.Call
}

// GetFull is a helper method to define mock.On call
//   - ctx context.Context
//   - db database.DBClient
//   - pool *models.Pool
func (_e *MockPoolQuery_Expecter) GetFull(ctx interface{}, db interface{}, pool interface{}) *MockPoolQuery_GetFull_Call {
	return &MockPoolQuery_GetFull_Call{Call: _e.mock.On("GetFull", ctx, db, pool)}
}

func (_c *MockPoolQuery_GetFull_Call) Run(run func(ctx context.Context, db database.DBClient, pool *models.Pool)) *MockPoolQuery_GetFull_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.DBClient), args[2].(*models.Pool))
	})
	return _c
}

func (_c *MockPoolQuery_GetFull_Call) Return(_a0 error) *MockPoolQuery_GetFull_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPoolQuery_GetFull_Call) RunAndReturn(run func(context.Context, database.DBClient, *models.Pool) error) *MockPoolQuery_GetFull_Call {
	_c.Call.Return(run)
	return _c
}

// ListFull provides a mock function with given fields: ctx, db, search, pools, count
func (_m *MockPoolQuery) ListFull(ctx context.Context, db database.DBClient, search models.Search, pools *[]models.Pool, count *int) error {
	ret := _m.Called(ctx, db, search, pools, count)

	if len(ret) == 0 {
		panic("no return value specified for ListFull")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.DBClient, models.Search, *[]models.Pool, *int) error); ok {
		r0 = rf(ctx, db, search, pools, count)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPoolQuery_ListFull_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListFull'
type MockPoolQuery_ListFull_Call struct {
	*mock.Call
}

// ListFull is a helper method to define mock.On call
//   - ctx context.Context
//   - db database.DBClient
//   - search models.Search
//   - pools *[]models.Pool
//   - count *int
func (_e *MockPoolQuery_Expecter) ListFull(ctx interface{}, db interface{}, search interface{}, pools interface{}, count interface{}) *MockPoolQuery_ListFull_Call {
	return &MockPoolQuery_ListFull_Call{Call: _e.mock.On("ListFull", ctx, db, search, pools, count)}
}

func (_c *MockPoolQuery_ListFull_Call) Run(run func(ctx context.Context, db database.DBClient, search models.Search, pools *[]models.Pool, count *int)) *MockPoolQuery_ListFull_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.DBClient), args[2].(models.Search), args[3].(*[]models.Pool), args[4].(*int))
	})
	return _c
}

func (_c *MockPoolQuery_ListFull_Call) Return(_a0 error) *MockPoolQuery_ListFull_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPoolQuery_ListFull_Call) RunAndReturn(run func(context.Context, database.DBClient, models.Search, *[]models.Pool, *int) error) *MockPoolQuery_ListFull_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, db, pool
func (_m *MockPoolQuery) Update(ctx context.Context, db database.DBClient, pool *models.Pool) error {
	ret := _m.Called(ctx, db, pool)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.DBClient, *models.Pool) error); ok {
		r0 = rf(ctx, db, pool)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPoolQuery_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockPoolQuery_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - db database.DBClient
//   - pool *models.Pool
func (_e *MockPoolQuery_Expecter) Update(ctx interface{}, db interface{}, pool interface{}) *MockPoolQuery_Update_Call {
	return &MockPoolQuery_Update_Call{Call: _e.mock.On("Update", ctx, db, pool)}
}

func (_c *MockPoolQuery_Update_Call) Run(run func(ctx context.Context, db database.DBClient, pool *models.Pool)) *MockPoolQuery_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.DBClient), args[2].(*models.Pool))
	})
	return _c
}

func (_c *MockPoolQuery_Update_Call) Return(_a0 error) *MockPoolQuery_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPoolQuery_Update_Call) RunAndReturn(run func(context.Context, database.DBClient, *models.Pool) error) *MockPoolQuery_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPoolQuery creates a new instance of MockPoolQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPoolQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPoolQuery {
	mock := &MockPoolQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
