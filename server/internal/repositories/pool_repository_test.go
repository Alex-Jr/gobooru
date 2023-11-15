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

func (s *TestSuite) SetupSuite() {
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

func (s *TestSuite) TearDownSuite() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()

	s.Require().NoError(s.psqlContainer.Terminate(ctx))
}

func TestSuite_Run(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestCreatePool() {
	tests := []struct {
		name         string
		ctx          context.Context
		args         repositories.PoolCreateArgs
		expectedPool models.Pool
	}{
		{
			name: "all fields",
			ctx:  context.Background(),
			args: repositories.PoolCreateArgs{
				Custom:      []string{"custom 1", "custom 2"},
				Description: "description 1",
				Name:        "creation 1",
				Posts:       []int{3, 1, 2},
			},
			expectedPool: models.Pool{
				Custom:      pq.StringArray{"custom 1", "custom 2"},
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
		{
			name: "no custom",
			ctx:  context.Background(),
			args: repositories.PoolCreateArgs{
				Description: "description 2",
				Name:        "creation 2",
				Posts:       []int{2, 1, 3},
			},
			expectedPool: models.Pool{
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
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			createdPool, err := s.poolRepository.Create(tt.ctx, tt.args)

			require.Nil(t, err)
			assert.NotZero(t, createdPool.ID)
			assert.NotZero(t, createdPool.CreatedAt)
			assert.NotZero(t, createdPool.UpdatedAt)
			assert.EqualValues(t, tt.expectedPool.Custom, createdPool.Custom)
			for i, post := range createdPool.Posts {
				assert.Equal(t, tt.expectedPool.Posts[i].ID, post.ID)
			}
			assert.Equal(t, tt.expectedPool.PostCount, createdPool.PostCount)
			assert.Equal(t, tt.expectedPool.Name, createdPool.Name)
			assert.Equal(t, tt.expectedPool.Description, createdPool.Description)

			fetchedPool, err := s.poolRepository.GetFull(context.TODO(), createdPool.ID)
			require.Nil(t, err)
			assert.EqualValues(t, createdPool, fetchedPool)
		})
	}
}
