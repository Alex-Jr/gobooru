package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gobooru/internal/controllers"
	"gobooru/internal/dtos"
	"gobooru/internal/fixtures/fakes"
	"gobooru/internal/mocks"
	"gobooru/internal/models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPoolControllerCreate(t *testing.T) {
	poolService := mocks.NewMockPoolService(t)

	poolController := controllers.NewPoolController(controllers.PoolControllerConfig{
		PoolService: poolService,
	})

	args := struct {
		Description string   `json:"description"`
		Name        string   `json:"name"`
		Posts       []int    `json:"posts"`
		Custom      []string `json:"custom"`
	}{
		Description: "test description",
		Name:        "test name",
		Posts:       []int{3, 1},
		Custom:      []string{},
	}

	want := struct {
		statusCode int
		dto        dtos.CreatePoolResponseDTO
	}{
		statusCode: http.StatusOK,
		dto: dtos.CreatePoolResponseDTO{
			Pool: fakes.LoadPool(fakes.Pool1),
		},
	}

	poolService.On(
		"Create",
		context.TODO(),
		dtos.CreatePoolDTO{
			Description: args.Description,
			Name:        args.Name,
			PostIDs:     args.Posts,
			Custom:      args.Custom,
		},
	).Return(
		want.dto,
		nil,
	)

	e := echo.New()
	requestData, err := json.Marshal(args)
	require.NoError(t, err)

	req, err := http.NewRequest(
		http.MethodPost,
		"/post",
		bytes.NewBuffer(requestData),
	)
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	require.NoError(t, poolController.Create(c))

	var responseDTO dtos.CreatePoolResponseDTO

	err = json.Unmarshal(rec.Body.Bytes(), &responseDTO)
	require.NoError(t, err)

	assert.Equal(t, want.statusCode, rec.Code)
	assert.EqualValues(t, want.dto, responseDTO)
}

func TestPoolControllerDelete(t *testing.T) {
	args := struct {
		ID int
	}{
		ID: 1,
	}

	want := struct {
		statusCode int
		dto        dtos.DeletePoolResponseDTO
	}{
		statusCode: http.StatusOK,
		dto: dtos.DeletePoolResponseDTO{
			Pool: fakes.LoadPool(fakes.Pool1),
		},
	}

	poolService := mocks.NewMockPoolService(t)

	poolController := controllers.NewPoolController(controllers.PoolControllerConfig{
		PoolService: poolService,
	})

	poolService.On(
		"Delete",
		context.TODO(),
		dtos.DeletePoolDTO{
			ID: args.ID,
		},
	).Return(
		want.dto,
		nil,
	)

	e := echo.New()
	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("/pool/%d", 1),
		nil,
	)
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", 1))

	require.NoError(t, poolController.Delete(c))
	require.Equal(t, want.statusCode, rec.Code)

	var responseDTO dtos.DeletePoolResponseDTO

	err = json.Unmarshal(rec.Body.Bytes(), &responseDTO)
	require.NoError(t, err)

	assert.EqualValues(t, want.dto, responseDTO)
}

func TestPoolControllerFetch(t *testing.T) {
	args := struct {
		ID int
	}{
		ID: 1,
	}

	want := struct {
		statusCode int
		dto        dtos.FetchPoolResponseDTO
	}{
		statusCode: http.StatusOK,
		dto: dtos.FetchPoolResponseDTO{
			Pool: fakes.LoadPool(fakes.Pool1),
		},
	}

	poolService := mocks.NewMockPoolService(t)

	poolController := controllers.NewPoolController(controllers.PoolControllerConfig{
		PoolService: poolService,
	})

	poolService.On(
		"Fetch",
		context.TODO(),
		dtos.FetchPoolDTO{
			ID: args.ID,
		},
	).Return(
		want.dto,
		nil,
	)

	e := echo.New()
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/pool/%d", 1),
		nil,
	)
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", 1))

	require.NoError(t, poolController.Fetch(c))
	require.Equal(t, want.statusCode, rec.Code)

	var responseDTO dtos.FetchPoolResponseDTO

	err = json.Unmarshal(rec.Body.Bytes(), &responseDTO)
	require.NoError(t, err)

	assert.EqualValues(t, want.dto, responseDTO)
}

func TestPoolControllerList(t *testing.T) {
	args := struct {
		Search   string
		Page     int
		PageSize int
	}{
		Search:   "createdAt:..2021-01-01",
		Page:     1,
		PageSize: 10,
	}

	want := struct {
		statusCode int
		dto        dtos.ListPoolResponseDTO
	}{
		statusCode: http.StatusOK,
		dto: dtos.ListPoolResponseDTO{
			Pools: []models.Pool{
				fakes.LoadPool(fakes.Pool1),
				fakes.LoadPool(fakes.Pool2),
			},
		},
	}

	poolService := mocks.NewMockPoolService(t)

	poolController := controllers.NewPoolController(controllers.PoolControllerConfig{
		PoolService: poolService,
	})

	poolService.On(
		"List",
		context.TODO(),
		dtos.ListPoolDTO{
			Search:   args.Search,
			Page:     args.Page,
			PageSize: args.PageSize,
		},
	).Return(
		want.dto,
		nil,
	)

	q := make(url.Values)
	q.Set("search", args.Search)
	q.Set("page", fmt.Sprintf("%d", args.Page))
	q.Set("page_size", fmt.Sprintf("%d", args.PageSize))

	e := echo.New()
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/pool?%s", q.Encode()),
		nil,
	)

	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	require.NoError(t, poolController.List(c))

	var responseDTO dtos.ListPoolResponseDTO

	err = json.Unmarshal(rec.Body.Bytes(), &responseDTO)
	require.NoError(t, err)

	assert.Equal(t, want.statusCode, rec.Code)
	assert.EqualValues(t, want.dto, responseDTO)
}

func TestPoolControllerUpdate(t *testing.T) {}
