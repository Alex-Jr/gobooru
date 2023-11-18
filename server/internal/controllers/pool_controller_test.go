package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"gobooru/internal/controllers"
	"gobooru/internal/database"
	"gobooru/internal/repositories"
	"gobooru/internal/services"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	psqlContainer  *database.PostgresContainer
	poolController controllers.PoolController
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

	s.psqlContainer = sqlContainer
	poolRepository := repositories.NewPoolRepository(db)
	poolService := services.NewPoolService(services.PoolRepositoryConfig{
		PoolRepository: poolRepository,
	})
	s.poolController = controllers.NewPoolController(controllers.PoolControllerConfig{
		PoolService: poolService,
	})
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

// func TestPoolControllerSuccess(t *testing.T) {
// 	poolService := mocks.NewMockPoolService(t)
//
// 	poolService.On(
// 		"Create",
// 		context.TODO(),
// 		dtos.CreatePoolDTO{
// 			Description: "test description",
// 			Name:        "test name",
// 			PostIDs:     []int{3, 1},
// 			Custom:      []string{},
// 		},
// 	).Return(
// 		dtos.CreatePoolResponseDTO{
// 			Pool: models.Pool{
// 				CreatedAt:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
// 				Custom:      []string{},
// 				Description: "test description",
// 				ID:          1,
// 				Name:        "test name",
// 				PostCount:   2,
// 				Posts: []models.Post{{
// 					CreatedAt:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
// 					Description: "",
// 					ID:          3,
// 					Pools:       nil,
// 					UpdatedAt:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
// 				}, {
// 					CreatedAt:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
// 					Description: "",
// 					ID:          1,
// 					Pools:       nil,
// 					UpdatedAt:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
// 				}},
// 				UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
// 			},
// 		},
// 		nil,
// 	)
//
// 	poolController := controllers.NewPoolController(controllers.PoolControllerConfig{
// 		PoolService: poolService,
// 	})
//
// 	e := echo.New()
// 	req, err := http.NewRequest(
// 		http.MethodPost,
// 		"/post",
// 		bytes.NewBuffer([]byte(`{
// 			"description": "test description",
// 			"name": "test name",
// 			"posts": [3, 1],
// 			"custom": []
// 		}`)),
// 	)
//
// 	req.Header.Set("Content-Type", "application/json")
//
// 	if err != nil {
// 		t.Errorf("The request could not be created because of: %v", err)
// 	}
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
//
// 	res := rec.Result()
// 	defer res.Body.Close()
//
// 	if assert.NoError(t, poolController.CreatePool(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Code)
//
// 		response := struct {
// 			Pool struct {
// 				CreatedAt   time.Time `json:"created_at"`
// 				Custom      []string  `json:"custom"`
// 				Description string    `json:"description"`
// 				ID          int       `json:"id"`
// 				Name        string    `json:"name"`
// 				PostCount   int       `json:"post_count"`
// 				Posts       []struct {
// 					CreatedAt   time.Time     `json:"created_at"`
// 					Description string        `json:"description"`
// 					ID          int           `json:"id"`
// 					Pool        []interface{} `json:"pool"`
// 					UpdatedAt   time.Time     `json:"updated_at"`
// 				} `json:"posts"`
// 				UpdatedAt time.Time `json:"updated_at"`
// 			} `json:"pool"`
// 		}{}
//
// 		json.Unmarshal(rec.Body.Bytes(), &response)
//
// 		assert.Equal(t, 1, response.Pool.ID)
// 		assert.Equal(t, "test name", response.Pool.Name)
// 		assert.Equal(t, "test description", response.Pool.Description)
// 		assert.Equal(t, 2, response.Pool.PostCount)
// 		assert.Equal(t, 2, len(response.Pool.Posts))
// 		assert.Equal(t, 0, len(response.Pool.Custom))
// 	}
// }

func (s *TestSuite) TestControllerPoolCreate() {
	e := echo.New()
	req, err := http.NewRequest(
		http.MethodPost,
		"/post",
		bytes.NewBuffer([]byte(`{
			"description": "test description",
			"name": "test name",
			"posts": [3, 1],
			"custom": []
		}`)),
	)
	s.Require().NoError(err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	res := rec.Result()
	defer res.Body.Close()

	if s.Assert().NoError(s.poolController.CreatePool(c)) {
		s.Assert().Equal(http.StatusOK, rec.Code)

		response := struct {
			Pool struct {
				CreatedAt   time.Time `json:"created_at"`
				Custom      []string  `json:"custom"`
				Description string    `json:"description"`
				ID          int       `json:"id"`
				Name        string    `json:"name"`
				PostCount   int       `json:"post_count"`
				Posts       []struct {
					CreatedAt   time.Time     `json:"created_at"`
					Description string        `json:"description"`
					ID          int           `json:"id"`
					Pool        []interface{} `json:"pool"`
					UpdatedAt   time.Time     `json:"updated_at"`
				} `json:"posts"`
				UpdatedAt time.Time `json:"updated_at"`
			} `json:"pool"`
		}{}

		json.Unmarshal(rec.Body.Bytes(), &response)

		s.Assert().NotZero(response.Pool.ID)
		s.Assert().Equal("test name", response.Pool.Name)
		s.Assert().Equal("test description", response.Pool.Description)
		s.Assert().Equal(2, response.Pool.PostCount)
		s.Assert().Equal(2, len(response.Pool.Posts))
		s.Assert().Equal([]string{}, response.Pool.Custom)
	}
}
