package repositories_test

import (
	"context"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/fixtures/fakes"
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
				s.Run(fmt.Sprintf("post tag:%s", tag.ID), func() {
					s.Assert().Equal(tc.want.post.Tags[i].ID, tag.ID)
					s.Assert().Equal(tc.want.post.Tags[i].PostCount, tag.PostCount)
					s.Assert().Zero(tag.Description)
					s.Assert().NotZero(tag.CreatedAt)
					s.Assert().NotZero(tag.UpdatedAt)
				})
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

func (s *PostTestSuit) TestPostGetFull() {
	type args struct {
		postID int
	}

	type want struct {
		post models.Post
		err  error
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "existing post",
			args: args{
				postID: 1,
			},
			want: want{
				post: fakes.LoadPostRelations(fakes.Post1),
				err:  nil,
			},
		},
		{
			name: "non-existing post",
			args: args{
				postID: 999,
			},
			want: want{
				post: models.Post{},
				err:  database.ErrNotFound,
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			post, err := s.postRepository.GetFull(
				context.TODO(),
				tc.args.postID,
			)

			s.Require().ErrorIs(err, tc.want.err)
			s.Assert().Equal(tc.want.post.ID, post.ID)
			s.Assert().Equal(tc.want.post.Description, post.Description)
			s.Assert().Equal(tc.want.post.Rating, post.Rating)
			s.Assert().Equal(tc.want.post.TagIDs, post.TagIDs)
			s.Assert().Equal(tc.want.post.TagCount, post.TagCount)
			s.Assert().Equal(tc.want.post.PoolCount, post.PoolCount)
			s.Assert().Equal(tc.want.post.CreatedAt.Compare(post.CreatedAt), 0)
			s.Assert().Equal(tc.want.post.UpdatedAt.Compare(post.UpdatedAt), 0)

			for i, tag := range post.Tags {
				s.Run(fmt.Sprintf("post tag: %s", tag.ID), func() {
					s.Assert().Equal(tc.want.post.Tags[i].ID, tag.ID)
					s.Assert().Equal(tc.want.post.Tags[i].PostCount, tag.PostCount)
					s.Assert().Equal(tc.want.post.Tags[i].Description, tag.Description)
					s.Assert().Equal(tc.want.post.Tags[i].CreatedAt.Compare(tag.CreatedAt), 0)
					s.Assert().Equal(tc.want.post.Tags[i].UpdatedAt.Compare(tag.UpdatedAt), 0)
				})
			}

			for i, pool := range post.Pools {
				s.Run(fmt.Sprintf("post pool: %d", pool.ID), func() {
					s.Assert().Equal(tc.want.post.Pools[i].ID, pool.ID)
					s.Assert().Equal(tc.want.post.Pools[i].PostCount, pool.PostCount)
					s.Assert().Equal(tc.want.post.Pools[i].Description, pool.Description)
					s.Assert().Equal(tc.want.post.Pools[i].CreatedAt.Compare(pool.CreatedAt), 0)
					s.Assert().Equal(tc.want.post.Pools[i].UpdatedAt.Compare(pool.UpdatedAt), 0)
				})
			}
		})
	}
}

func (s *PostTestSuit) TestPostList() {
	type args struct {
		ctx      context.Context
		search   string
		page     int
		pageSize int
	}

	type want struct {
		posts []models.Post
		count int
		err   error
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "get all posts",
			args: args{
				ctx: context.TODO(),
			},
			want: want{
				posts: models.PostList{
					fakes.LoadPostNoRelations(fakes.Post5),
					fakes.LoadPostNoRelations(fakes.Post4),
					fakes.LoadPostNoRelations(fakes.Post3),
					fakes.LoadPostNoRelations(fakes.Post2),
					fakes.LoadPostNoRelations(fakes.Post1),
				},
				count: 5,
				err:   nil,
			},
		},
		{
			name: "get posts first page",
			args: args{
				ctx:      context.TODO(),
				page:     1,
				pageSize: 2,
			},
			want: want{
				posts: models.PostList{
					fakes.LoadPostNoRelations(fakes.Post5),
					fakes.LoadPostNoRelations(fakes.Post4),
				},
				count: 5,
				err:   nil,
			},
		},
		{
			name: "get posts second page",
			args: args{
				ctx:      context.TODO(),
				page:     2,
				pageSize: 2,
			},
			want: want{
				posts: models.PostList{
					fakes.LoadPostNoRelations(fakes.Post3),
					fakes.LoadPostNoRelations(fakes.Post2),
				},
				count: 5,
				err:   nil,
			},
		},
		{
			name: "get posts third page",
			args: args{
				ctx:      context.TODO(),
				page:     3,
				pageSize: 2,
			},
			want: want{
				posts: models.PostList{
					fakes.LoadPostNoRelations(fakes.Post1),
				},
				count: 5,
				err:   nil,
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			posts, count, err := s.postRepository.List(
				tc.args.ctx,
				repositories.ListPostsArgs{
					Search:   tc.args.search,
					Page:     tc.args.page,
					PageSize: tc.args.pageSize,
				},
			)

			s.Require().ErrorIs(err, tc.want.err)
			s.Require().Equal(len(tc.want.posts), len(posts))
			s.Assert().Equal(tc.want.count, count)

			for i, post := range posts {
				s.Run(fmt.Sprintf("post: %d", post.ID), func() {

					s.Assert().Equal(tc.want.posts[i].ID, post.ID)
					s.Assert().Equal(tc.want.posts[i].Description, post.Description)
					s.Assert().Equal(tc.want.posts[i].Rating, post.Rating)
					s.Assert().Equal(tc.want.posts[i].TagIDs, post.TagIDs)
					s.Assert().Equal(tc.want.posts[i].TagCount, post.TagCount)
				})
			}
		})
	}

}
