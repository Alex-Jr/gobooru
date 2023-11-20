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
	suite.Run(t, new(PostTestSuit))
}

func (s *PostTestSuit) TestPostCreate() {
	type args struct {
		description string
		rating      string
		tags        []string
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
			name: "create post with tags",
			args: args{
				description: "test description",
				rating:      "s",
				tags:        []string{"tag_one", "tag_two"},
			},
			want: want{
				post: models.Post{
					Description: "test description",
					Rating:      "s",
					TagIDs:      []string{"tag_one", "tag_two"},
					TagCount:    2,
					PoolCount:   0,
					Pools:       make(models.PoolList, 0),
					Tags: models.TagList{
						{
							ID:        "tag_one",
							PostCount: 2,
						},
						{
							ID:        "tag_two",
							PostCount: 1,
						},
					},
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
					Rating:      tc.args.rating,
					Tags:        tc.args.tags,
				},
			)

			s.Require().NoError(err)
			s.Assert().NotZero(post.ID)
			s.Assert().NotZero(post.CreatedAt)
			s.Assert().NotZero(post.UpdatedAt)
			s.Assert().Equal(tc.want.post.Description, post.Description)
			s.Assert().Equal(tc.want.post.Rating, post.Rating)
			s.Assert().Equal(tc.want.post.TagIDs, post.TagIDs)
			s.Assert().Equal(tc.want.post.TagCount, post.TagCount)
			s.Assert().Equal(tc.want.post.PoolCount, post.PoolCount)
			s.Assert().Equal(tc.want.post.Pools, post.Pools)

			for i, tag := range post.Tags {
				s.Assert().Equal(tc.want.post.Tags[i].ID, tag.ID)
				s.Assert().Equal(tc.want.post.Tags[i].PostCount, tag.PostCount)
				s.Assert().Zero(tag.Description)
				s.Assert().NotZero(tag.CreatedAt)
				s.Assert().NotZero(tag.UpdatedAt)
			}
		})
	}
}

func (s *PostTestSuit) TestPostDelete() {
	type args struct {
		postID int
	}

	testCases := []struct {
		name string
		args args
	}{
		{
			name: "delete post",
			args: args{
				postID: 1,
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			err := s.postRepository.Delete(
				context.TODO(),
				tc.args.postID,
			)

			s.Require().NoError(err)
		})
	}
}
