package repositories_test

import (
	"context"
	"gobooru/internal/database"
	"gobooru/internal/fixtures/fakes"
	"gobooru/internal/models"
	"gobooru/internal/repositories"
	"testing"
	"time"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/stretchr/testify/suite"
)

type TagTestSuit struct {
	suite.Suite
	psqlContainer *database.PostgresContainer
	tagRepository repositories.TagRepository
}

func (s *TagTestSuit) SetupTest() {
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

	s.tagRepository = repositories.NewTagRepository(repositories.TagRepositoryConfig{
		SQLClient: db,
	})
	s.psqlContainer = sqlContainer
}

func (s *TagTestSuit) TearDownTest() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()

	err := s.psqlContainer.Terminate(ctx)
	s.Require().NoError(err)
}

func TestTagSuite_Run(t *testing.T) {
	suite.Run(t, new(PostTestSuit))
}

func (s *TagTestSuit) TestTagDelete() {
	type args struct {
		tagID string
	}

	type want struct {
		err error
		tag models.Tag
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "delete tag",
			args: args{
				tagID: "tag_one",
			},
			want: want{
				err: nil,
				tag: fakes.LoadTagRelations(fakes.Tag1),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tag, err := s.tagRepository.Delete(
				context.Background(),
				tc.args.tagID,
			)
			s.Require().Equal(tc.want.err, err)
			s.Require().Equal(tc.want.tag, tag)

			_, err = s.tagRepository.Get(context.Background(), tc.args.tagID)
			s.Require().Equal(database.ErrNotFound, err)
		})
	}
}
