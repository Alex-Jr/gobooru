package repositories_test

import (
	"context"
	"gobooru/internal/database"
	"gobooru/internal/models"
	"gobooru/internal/repositories"
	"testing"
	"time"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	psqlContainer  *database.PostgresContainer
	poolRepository repositories.PoolRepository
}

func (s *TestSuite) SetupTest() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer ctxCancel()

	sqlContainer, err := database.NewPostgresContainer(ctx)
	s.Require().NoError(err)

	sqlx, err := database.GetSQLXConnection(sqlContainer.Config.DBConfig)
	s.Require().NoError(err)

	db := database.NewSQLClient(sqlx)

	err = database.RunMigrations(sqlx)
	s.Require().NoError(err)

	fixtures, err := testfixtures.New(
		testfixtures.Database(sqlx.DB),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory("../fixtures/storage"),
	)
	s.Require().NoError(err)

	err = fixtures.Load()
	s.Require().NoError(err)

	s.poolRepository = repositories.NewPoolRepository(db)
	s.psqlContainer = sqlContainer
}

func (s *TestSuite) TearDownTest() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()

	err := s.psqlContainer.Terminate(ctx)
	s.Require().NoError(err)
}

func TestSuite_Run(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestCreatePool() {
	type args struct {
		Custom      []string
		Description string
		Name        string
		PostIDs     []int
	}

	type want struct {
		pool models.Pool
		err  error
	}

	testCases := []struct {
		name string
		ctx  context.Context
		args args
		want want
	}{
		{
			name: "all fields",
			ctx:  context.Background(),
			args: args{
				Custom:      []string{"a", "b"},
				Description: "description 1",
				Name:        "creation 1",
				PostIDs:     []int{3, 1, 2},
			},
			want: want{
				err: nil,
				pool: models.Pool{
					Custom:      pq.StringArray{"a", "b"},
					Description: "description 1",
					Name:        "creation 1",
					PostCount:   3,
					Posts: []models.Post{
						{
							ID: 3,
						},
						{
							ID: 1,
						},
						{
							ID: 2,
						},
					},
				},
			},
		},
		{
			name: "no custom",
			ctx:  context.Background(),
			args: args{
				Description: "description 2",
				Name:        "creation 2",
				PostIDs:     []int{2, 1, 3},
			},
			want: want{
				err: nil,
				pool: models.Pool{
					Custom:      pq.StringArray{},
					Description: "description 2",
					Name:        "creation 2",
					PostCount:   3,
					Posts: []models.Post{
						{
							ID: 2,
						},
						{
							ID: 1,
						},
						{
							ID: 3,
						},
					},
				},
			},
		},
	}

	for _, tt := range testCases {
		s.T().Run(tt.name, func(t *testing.T) {
			createdPool, err := s.poolRepository.Create(tt.ctx, repositories.PoolCreateArgs{
				Custom:      tt.args.Custom,
				Description: tt.args.Description,
				Name:        tt.args.Name,
				PostIDs:     tt.args.PostIDs,
			})

			require.ErrorIs(t, err, tt.want.err)
			assert.NotZero(t, createdPool.ID)
			assert.NotZero(t, createdPool.CreatedAt)
			assert.NotZero(t, createdPool.UpdatedAt)
			assert.EqualValues(t, tt.want.pool.Custom, createdPool.Custom)
			assert.Equal(t, tt.want.pool.PostCount, createdPool.PostCount)
			assert.Equal(t, tt.want.pool.Name, createdPool.Name)
			assert.Equal(t, tt.want.pool.Description, createdPool.Description)

			// checks for post order
			for i, post := range createdPool.Posts {
				assert.Equal(t, tt.want.pool.Posts[i].ID, post.ID)
			}

			fetchedPool, err := s.poolRepository.GetFull(context.TODO(), createdPool.ID)
			require.Nil(t, err)
			assert.EqualValues(t, createdPool, fetchedPool)
		})
	}
}

func (s *TestSuite) TestDeletePool() {
	type args struct {
		poolID int
	}

	type want struct {
		err error
	}

	testCases := []struct {
		name string
		ctx  context.Context
		args args
		want want
	}{
		{
			name: "existing pool 1",
			ctx:  context.Background(),
			args: args{
				poolID: 1,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "non-existing pool",
			ctx:  context.Background(),
			args: args{
				poolID: 9999,
			},
			want: want{
				err: nil,
			},
		},
	}

	for _, tt := range testCases {
		s.T().Run(tt.name, func(t *testing.T) {
			err := s.poolRepository.Delete(tt.ctx, tt.args.poolID)
			require.ErrorIs(t, err, tt.want.err)

			_, err = s.poolRepository.GetFull(tt.ctx, tt.args.poolID)
			assert.ErrorIs(t, err, database.ErrNotFound)
		})
	}
}

func (s *TestSuite) TestGetPool() {
	type args struct {
		poolID int
	}

	type want struct {
		pool models.Pool
		err  error
	}

	tests := []struct {
		name string
		ctx  context.Context
		args args
		want want
	}{
		{
			name: "existing pool 3",
			ctx:  context.Background(),
			args: args{
				poolID: 3,
			},
			want: want{
				err: nil,
				pool: models.Pool{
					ID:          3,
					CreatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					Custom:      pq.StringArray{},
					Description: "pool 3 description",
					Name:        "pool 3 name",
					PostCount:   3,
					Posts: []models.Post{
						{
							CreatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
							Description: "post 3 description",
							ID:          3,
							Pools:       []models.Post{},
							UpdatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
						},
						{
							CreatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
							Description: "post 1 description",
							ID:          1,
							Pools:       []models.Post{},
							UpdatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
						},
						{
							CreatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
							Description: "post 2 description",
							ID:          2,
							Pools:       []models.Post{},
							UpdatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
						},
					},
				},
			},
		},
		{
			name: "non-existing pool",
			ctx:  context.Background(),
			args: args{
				poolID: 9999,
			},
			want: want{
				err: database.ErrNotFound,
				pool: models.Pool{
					ID:          0,
					CreatedAt:   time.Time{},
					UpdatedAt:   time.Time{},
					Custom:      nil,
					Description: "",
					Name:        "",
					PostCount:   0,
					Posts:       nil,
				},
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			pool, err := s.poolRepository.GetFull(tt.ctx, tt.args.poolID)

			require.ErrorIs(t, err, tt.want.err)

			assert.Equal(t, tt.want.pool.ID, pool.ID)
			assert.Equal(t, tt.want.pool.CreatedAt, pool.CreatedAt)
			assert.Equal(t, tt.want.pool.UpdatedAt, pool.UpdatedAt)
			assert.EqualValues(t, tt.want.pool.Custom, pool.Custom)
			assert.Equal(t, tt.want.pool.PostCount, pool.PostCount)
			assert.Equal(t, tt.want.pool.Name, pool.Name)
			assert.Equal(t, tt.want.pool.Description, pool.Description)

			for i, post := range pool.Posts {
				// checks for post order
				assert.Equal(t, tt.want.pool.Posts[i].ID, post.ID)
				assert.Equal(t, tt.want.pool.Posts[i].CreatedAt, post.CreatedAt)
				assert.Equal(t, tt.want.pool.Posts[i].UpdatedAt, post.UpdatedAt)
				assert.Equal(t, tt.want.pool.Posts[i].Description, post.Description)
			}
		})
	}
}

func (s *TestSuite) TestListPool() {
	tests := []struct {
		name          string
		ctx           context.Context
		args          repositories.PoolListFullArgs
		expectedPools []models.Pool
		expectedCount int
	}{
		{
			name: "filter by custom first page",
			ctx:  context.Background(),
			args: repositories.PoolListFullArgs{
				Text:     "custom:shared",
				Page:     1,
				PageSize: 2,
			},
			expectedPools: []models.Pool{
				{
					ID: 6,
				},
				{
					ID: 5,
				},
			},
			expectedCount: 3,
		},
		{
			name: "filter by custom second page",
			ctx:  context.Background(),
			args: repositories.PoolListFullArgs{
				Text:     "custom:shared",
				Page:     2,
				PageSize: 2,
			},
			expectedPools: []models.Pool{
				{
					ID: 4,
				},
			},
			expectedCount: 3,
		},
		{
			name: "filter by name",
			ctx:  context.Background(),
			args: repositories.PoolListFullArgs{
				Text:     "name:pool",
				Page:     1,
				PageSize: 1,
			},
			expectedPools: []models.Pool{
				{
					ID: 6,
				},
			},
			expectedCount: 6,
		},
		{
			name: "filter by less than created_at",
			ctx:  context.Background(),
			args: repositories.PoolListFullArgs{
				Text:     "createdAt:..2021-01-01",
				Page:     1,
				PageSize: 2,
			},
			expectedPools: []models.Pool{
				{
					ID: 6,
				},
				{
					ID: 5,
				},
			},
			expectedCount: 4,
		},
		{
			name: "filter by greater than created_at",
			ctx:  context.Background(),
			args: repositories.PoolListFullArgs{
				Text:     "createdAt:2021-01-01..",
				Page:     1,
				PageSize: 2,
			},
			expectedPools: []models.Pool{
				{
					ID: 2,
				},
				{
					ID: 1,
				},
			},
			expectedCount: 2,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			pools, count, err := s.poolRepository.ListFull(tt.ctx, tt.args)

			require.Nil(t, err)
			assert.Equal(t, len(tt.expectedPools), len(pools))
			assert.Equal(t, tt.expectedCount, count)
			for i, pool := range pools {
				assert.Equal(t, tt.expectedPools[i].ID, pool.ID)
			}
		})
	}

}
func (s *TestSuite) TestUpdatePool() {
	makeStringPointer := func(s string) *string {
		return &s
	}

	tests := []struct {
		name          string
		ctx           context.Context
		args          repositories.PoolUpdateArgs
		expectedError error
	}{
		{
			name: "description",
			ctx:  context.Background(),
			args: repositories.PoolUpdateArgs{
				ID:          1,
				Description: makeStringPointer("updated description 1"),
			},
			expectedError: nil,
		},
		{
			name: "name",
			ctx:  context.Background(),
			args: repositories.PoolUpdateArgs{
				ID:   1,
				Name: makeStringPointer("updated name 1"),
			},
			expectedError: nil,
		},
		{
			name: "custom",
			ctx:  context.Background(),
			args: repositories.PoolUpdateArgs{
				ID:     1,
				Custom: &[]string{"x", "y"},
			},
			expectedError: nil,
		},
		{
			name: "posts",
			ctx:  context.Background(),
			args: repositories.PoolUpdateArgs{
				ID:    1,
				Posts: &[]int{2, 3},
			},
			expectedError: nil,
		},
		{
			name: "posts 2",
			ctx:  context.Background(),
			args: repositories.PoolUpdateArgs{
				ID:    1,
				Posts: &[]int{1, 2, 3},
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			oldPool, _ := s.poolRepository.GetFull(tt.ctx, tt.args.ID)

			updatedPool, err := s.poolRepository.Update(tt.ctx, tt.args)
			require.ErrorIs(t, err, tt.expectedError)

			newPool, _ := s.poolRepository.GetFull(tt.ctx, tt.args.ID)

			if tt.args.Custom != nil {
				assert.NotEqual(t, updatedPool.Custom, oldPool.Custom)
				assert.Equal(t, updatedPool.Custom, newPool.Custom)
			} else {
				assert.Equal(t, updatedPool.Custom, oldPool.Custom)
			}

			if tt.args.Description != nil {
				assert.NotEqual(t, updatedPool.Description, oldPool.Description)
				assert.Equal(t, updatedPool.Description, newPool.Description)
			} else {
				assert.Equal(t, updatedPool.Description, oldPool.Description)
			}

			if tt.args.Name != nil {
				assert.NotEqual(t, updatedPool.Name, oldPool.Name)
				assert.Equal(t, updatedPool.Name, newPool.Name)
			} else {
				assert.Equal(t, updatedPool.Name, oldPool.Name)
			}

			if tt.args.Posts != nil {
				assert.NotEqual(t, updatedPool.Posts, oldPool.Posts)
				assert.Equal(t, updatedPool.Posts, newPool.Posts)
				// checks for post order
				for i, post := range updatedPool.Posts {
					assert.Equal(t, (*tt.args.Posts)[i], post.ID)
					assert.Equal(t, (*tt.args.Posts)[i], newPool.Posts[i].ID)
				}
			} else {
				assert.Equal(t, updatedPool.Posts, oldPool.Posts)
			}

			if tt.args.Custom != nil {
				assert.NotEqual(t, updatedPool.Custom, oldPool.Custom)
				assert.Equal(t, updatedPool.Custom, newPool.Custom)
			} else {
				assert.Equal(t, updatedPool.Custom, oldPool.Custom)
			}
		})
	}
}
