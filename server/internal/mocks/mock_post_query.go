// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"
	database "gobooru/internal/database"

	mock "github.com/stretchr/testify/mock"

	models "gobooru/internal/models"
)

// MockPostQuery is an autogenerated mock type for the PostQuery type
type MockPostQuery struct {
	mock.Mock
}

type MockPostQuery_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPostQuery) EXPECT() *MockPostQuery_Expecter {
	return &MockPostQuery_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, db, post
func (_m *MockPostQuery) Create(ctx context.Context, db database.DBClient, post *models.Post) error {
	ret := _m.Called(ctx, db, post)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.DBClient, *models.Post) error); ok {
		r0 = rf(ctx, db, post)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPostQuery_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockPostQuery_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - db database.DBClient
//   - post *models.Post
func (_e *MockPostQuery_Expecter) Create(ctx interface{}, db interface{}, post interface{}) *MockPostQuery_Create_Call {
	return &MockPostQuery_Create_Call{Call: _e.mock.On("Create", ctx, db, post)}
}

func (_c *MockPostQuery_Create_Call) Run(run func(ctx context.Context, db database.DBClient, post *models.Post)) *MockPostQuery_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.DBClient), args[2].(*models.Post))
	})
	return _c
}

func (_c *MockPostQuery_Create_Call) Return(_a0 error) *MockPostQuery_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPostQuery_Create_Call) RunAndReturn(run func(context.Context, database.DBClient, *models.Post) error) *MockPostQuery_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, db, post
func (_m *MockPostQuery) Delete(ctx context.Context, db database.DBClient, post *models.Post) error {
	ret := _m.Called(ctx, db, post)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.DBClient, *models.Post) error); ok {
		r0 = rf(ctx, db, post)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPostQuery_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockPostQuery_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - db database.DBClient
//   - post *models.Post
func (_e *MockPostQuery_Expecter) Delete(ctx interface{}, db interface{}, post interface{}) *MockPostQuery_Delete_Call {
	return &MockPostQuery_Delete_Call{Call: _e.mock.On("Delete", ctx, db, post)}
}

func (_c *MockPostQuery_Delete_Call) Run(run func(ctx context.Context, db database.DBClient, post *models.Post)) *MockPostQuery_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.DBClient), args[2].(*models.Post))
	})
	return _c
}

func (_c *MockPostQuery_Delete_Call) Return(_a0 error) *MockPostQuery_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPostQuery_Delete_Call) RunAndReturn(run func(context.Context, database.DBClient, *models.Post) error) *MockPostQuery_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetFull provides a mock function with given fields: ctx, db, post
func (_m *MockPostQuery) GetFull(ctx context.Context, db database.DBClient, post *models.Post) error {
	ret := _m.Called(ctx, db, post)

	if len(ret) == 0 {
		panic("no return value specified for GetFull")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.DBClient, *models.Post) error); ok {
		r0 = rf(ctx, db, post)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPostQuery_GetFull_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFull'
type MockPostQuery_GetFull_Call struct {
	*mock.Call
}

// GetFull is a helper method to define mock.On call
//   - ctx context.Context
//   - db database.DBClient
//   - post *models.Post
func (_e *MockPostQuery_Expecter) GetFull(ctx interface{}, db interface{}, post interface{}) *MockPostQuery_GetFull_Call {
	return &MockPostQuery_GetFull_Call{Call: _e.mock.On("GetFull", ctx, db, post)}
}

func (_c *MockPostQuery_GetFull_Call) Run(run func(ctx context.Context, db database.DBClient, post *models.Post)) *MockPostQuery_GetFull_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.DBClient), args[2].(*models.Post))
	})
	return _c
}

func (_c *MockPostQuery_GetFull_Call) Return(_a0 error) *MockPostQuery_GetFull_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPostQuery_GetFull_Call) RunAndReturn(run func(context.Context, database.DBClient, *models.Post) error) *MockPostQuery_GetFull_Call {
	_c.Call.Return(run)
	return _c
}

// GetFullByHash provides a mock function with given fields: ctx, db, post
func (_m *MockPostQuery) GetFullByHash(ctx context.Context, db database.DBClient, post *models.Post) error {
	ret := _m.Called(ctx, db, post)

	if len(ret) == 0 {
		panic("no return value specified for GetFullByHash")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.DBClient, *models.Post) error); ok {
		r0 = rf(ctx, db, post)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPostQuery_GetFullByHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFullByHash'
type MockPostQuery_GetFullByHash_Call struct {
	*mock.Call
}

// GetFullByHash is a helper method to define mock.On call
//   - ctx context.Context
//   - db database.DBClient
//   - post *models.Post
func (_e *MockPostQuery_Expecter) GetFullByHash(ctx interface{}, db interface{}, post interface{}) *MockPostQuery_GetFullByHash_Call {
	return &MockPostQuery_GetFullByHash_Call{Call: _e.mock.On("GetFullByHash", ctx, db, post)}
}

func (_c *MockPostQuery_GetFullByHash_Call) Run(run func(ctx context.Context, db database.DBClient, post *models.Post)) *MockPostQuery_GetFullByHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.DBClient), args[2].(*models.Post))
	})
	return _c
}

func (_c *MockPostQuery_GetFullByHash_Call) Return(_a0 error) *MockPostQuery_GetFullByHash_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPostQuery_GetFullByHash_Call) RunAndReturn(run func(context.Context, database.DBClient, *models.Post) error) *MockPostQuery_GetFullByHash_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, db, search, posts, count
func (_m *MockPostQuery) List(ctx context.Context, db database.DBClient, search models.Search, posts *[]models.Post, count *int) error {
	ret := _m.Called(ctx, db, search, posts, count)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.DBClient, models.Search, *[]models.Post, *int) error); ok {
		r0 = rf(ctx, db, search, posts, count)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPostQuery_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockPostQuery_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - db database.DBClient
//   - search models.Search
//   - posts *[]models.Post
//   - count *int
func (_e *MockPostQuery_Expecter) List(ctx interface{}, db interface{}, search interface{}, posts interface{}, count interface{}) *MockPostQuery_List_Call {
	return &MockPostQuery_List_Call{Call: _e.mock.On("List", ctx, db, search, posts, count)}
}

func (_c *MockPostQuery_List_Call) Run(run func(ctx context.Context, db database.DBClient, search models.Search, posts *[]models.Post, count *int)) *MockPostQuery_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.DBClient), args[2].(models.Search), args[3].(*[]models.Post), args[4].(*int))
	})
	return _c
}

func (_c *MockPostQuery_List_Call) Return(_a0 error) *MockPostQuery_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPostQuery_List_Call) RunAndReturn(run func(context.Context, database.DBClient, models.Search, *[]models.Post, *int) error) *MockPostQuery_List_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveTag provides a mock function with given fields: ctx, db, tag
func (_m *MockPostQuery) RemoveTag(ctx context.Context, db database.DBClient, tag string) error {
	ret := _m.Called(ctx, db, tag)

	if len(ret) == 0 {
		panic("no return value specified for RemoveTag")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.DBClient, string) error); ok {
		r0 = rf(ctx, db, tag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPostQuery_RemoveTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveTag'
type MockPostQuery_RemoveTag_Call struct {
	*mock.Call
}

// RemoveTag is a helper method to define mock.On call
//   - ctx context.Context
//   - db database.DBClient
//   - tag string
func (_e *MockPostQuery_Expecter) RemoveTag(ctx interface{}, db interface{}, tag interface{}) *MockPostQuery_RemoveTag_Call {
	return &MockPostQuery_RemoveTag_Call{Call: _e.mock.On("RemoveTag", ctx, db, tag)}
}

func (_c *MockPostQuery_RemoveTag_Call) Run(run func(ctx context.Context, db database.DBClient, tag string)) *MockPostQuery_RemoveTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.DBClient), args[2].(string))
	})
	return _c
}

func (_c *MockPostQuery_RemoveTag_Call) Return(_a0 error) *MockPostQuery_RemoveTag_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPostQuery_RemoveTag_Call) RunAndReturn(run func(context.Context, database.DBClient, string) error) *MockPostQuery_RemoveTag_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, db, post
func (_m *MockPostQuery) Update(ctx context.Context, db database.DBClient, post models.Post) error {
	ret := _m.Called(ctx, db, post)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.DBClient, models.Post) error); ok {
		r0 = rf(ctx, db, post)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPostQuery_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockPostQuery_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - db database.DBClient
//   - post models.Post
func (_e *MockPostQuery_Expecter) Update(ctx interface{}, db interface{}, post interface{}) *MockPostQuery_Update_Call {
	return &MockPostQuery_Update_Call{Call: _e.mock.On("Update", ctx, db, post)}
}

func (_c *MockPostQuery_Update_Call) Run(run func(ctx context.Context, db database.DBClient, post models.Post)) *MockPostQuery_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.DBClient), args[2].(models.Post))
	})
	return _c
}

func (_c *MockPostQuery_Update_Call) Return(_a0 error) *MockPostQuery_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPostQuery_Update_Call) RunAndReturn(run func(context.Context, database.DBClient, models.Post) error) *MockPostQuery_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdatePoolCount provides a mock function with given fields: ctx, db, post, increment
func (_m *MockPostQuery) UpdatePoolCount(ctx context.Context, db database.DBClient, post []models.Post, increment int) error {
	ret := _m.Called(ctx, db, post, increment)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePoolCount")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.DBClient, []models.Post, int) error); ok {
		r0 = rf(ctx, db, post, increment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPostQuery_UpdatePoolCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdatePoolCount'
type MockPostQuery_UpdatePoolCount_Call struct {
	*mock.Call
}

// UpdatePoolCount is a helper method to define mock.On call
//   - ctx context.Context
//   - db database.DBClient
//   - post []models.Post
//   - increment int
func (_e *MockPostQuery_Expecter) UpdatePoolCount(ctx interface{}, db interface{}, post interface{}, increment interface{}) *MockPostQuery_UpdatePoolCount_Call {
	return &MockPostQuery_UpdatePoolCount_Call{Call: _e.mock.On("UpdatePoolCount", ctx, db, post, increment)}
}

func (_c *MockPostQuery_UpdatePoolCount_Call) Run(run func(ctx context.Context, db database.DBClient, post []models.Post, increment int)) *MockPostQuery_UpdatePoolCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.DBClient), args[2].([]models.Post), args[3].(int))
	})
	return _c
}

func (_c *MockPostQuery_UpdatePoolCount_Call) Return(_a0 error) *MockPostQuery_UpdatePoolCount_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPostQuery_UpdatePoolCount_Call) RunAndReturn(run func(context.Context, database.DBClient, []models.Post, int) error) *MockPostQuery_UpdatePoolCount_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPostQuery creates a new instance of MockPostQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPostQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPostQuery {
	mock := &MockPostQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
