package repositories_test

import (
	"context"
	"gobooru/internal/database"
	"gobooru/internal/models"
	"gobooru/internal/repositories"
	"testing"
	"time"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/stretchr/testify/suite"
)

type PostTestSuit struct {
	suite.Suite
	psqlContainer  *database.PostgresContainer
	postRepository repositories.PostRepository
}

func (s *PostTestSuit) SetupTest() {
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

	s.postRepository = repositories.NewPostRepository(db)
	s.psqlContainer = sqlContainer
}

func (s *PostTestSuit) TearDownTest() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()

	err := s.psqlContainer.Terminate(ctx)
	s.Require().NoError(err)
}

func TestPostSuite_Run(t *testing.T) {
	suite.Run(t, new(PoolTestSuite))
}

func (s *PostTestSuit) TestCreate() {
	type args struct {
		description string
	}

	type want struct {
		post models.Post
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "create post",
			args: args{
				description: "test description",
			},
			want: want{
				post: models.Post{
					Description: "test description",
				},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			post, err := s.postRepository.Create(
				context.TODO(),
				repositories.CreatePostArgs{
					Description: tc.args.description,
				},
			)

			s.Require().NoError(err)
			s.Assert().NotZero(post.ID)
			s.Assert().NotZero(post.CreatedAt)
			s.Assert().NotZero(post.UpdatedAt)
			s.Assert().Equal(tc.want.post.Description, post.Description)
		})
	}
}
