package services_test

import (
	"context"
	"encoding/json"
	"gobooru/internal/dtos"
	"gobooru/internal/fixtures/fakes"
	"gobooru/internal/mocks"
	"gobooru/internal/models"
	"gobooru/internal/repositories"
	"gobooru/internal/services"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPoolServiceCreate(t *testing.T) {
	poolRepository := mocks.NewMockPoolRepository(t)

	poolService := services.NewPoolService(services.PoolRepositoryConfig{
		PoolRepository: poolRepository,
	})

	poolRepository.On(
		"Create",
		context.TODO(),
		repositories.PoolCreateArgs{
			Name:        "pool 1 description",
			Description: "pool 1 name",
			Custom:      []string{"a"},
			PostIDs:     []int{1},
		},
	).Return(
		fakes.LoadPool(fakes.Pool1),
		nil,
	)

	response, err := poolService.Create(context.TODO(), dtos.CreatePoolDTO{
		Name:        "pool 1 description",
		Description: "pool 1 name",
		Custom:      []string{"a"},
		PostIDs:     []int{1},
	})
	require.NoError(t, err)

	b, err := json.Marshal(response)
	require.NoError(t, err)

	require.JSONEq(t, `
		{
			"pool": {
				"created_at": "2022-01-01T00:00:00Z",
				"custom": ["a"],
				"description": "pool 1 description",
				"id": 1,
				"name": "pool 1 name",
				"post_count": 1,
				"updated_at": "2022-01-01T00:00:00Z",
				"posts": [
					{
						"created_at": "2020-01-01T00:00:00Z",
						"description": "post 1 description",
						"id": 1,
						"pools": null,
						"updated_at": "2020-01-01T00:00:00Z"
					}
				]
			}
		}
	`, string(b))

	poolRepository.AssertExpectations(t)
}

func TestPoolServiceDelete(t *testing.T) {
	poolRepository := mocks.NewMockPoolRepository(t)

	poolService := services.NewPoolService(services.PoolRepositoryConfig{
		PoolRepository: poolRepository,
	})

	poolRepository.On(
		"GetFull",
		context.TODO(),
		1,
	).Return(
		fakes.LoadPool(fakes.Pool1),
		nil,
	)

	poolRepository.On(
		"Delete",
		context.TODO(),
		1,
	).Return(
		nil,
	)

	response, err := poolService.Delete(
		context.TODO(),
		dtos.DeletePoolDTO{
			ID: 1,
		},
	)
	require.NoError(t, err)

	b, err := json.Marshal(response)
	require.NoError(t, err)

	require.JSONEq(t,
		`{
			"pool": {
				"created_at": "2022-01-01T00:00:00Z",
				"custom": ["a"],
				"description": "pool 1 description",
				"id": 1,
				"name": "pool 1 name",
				"post_count": 1,
				"updated_at": "2022-01-01T00:00:00Z",
				"posts": [
					{
						"created_at": "2020-01-01T00:00:00Z",
						"description": "post 1 description",
						"id": 1,
						"pools": null,
						"updated_at": "2020-01-01T00:00:00Z"
					}
				]
			}
		}`,
		string(b))
}

func TestPoolServiceFetch(t *testing.T) {
	poolRepository := mocks.NewMockPoolRepository(t)

	poolService := services.NewPoolService(services.PoolRepositoryConfig{
		PoolRepository: poolRepository,
	})

	poolRepository.On(
		"GetFull",
		context.TODO(),
		1,
	).Return(
		fakes.LoadPool(fakes.Pool1),
		nil,
	)

	response, err := poolService.Fetch(
		context.TODO(),
		dtos.FetchPoolDTO{
			ID: 1,
		},
	)
	require.NoError(t, err)

	b, err := json.Marshal(response)
	require.NoError(t, err)

	require.JSONEq(t,
		`{
			"pool": {
				"created_at": "2022-01-01T00:00:00Z",
				"custom": ["a"],
				"description": "pool 1 description",
				"id": 1,
				"name": "pool 1 name",
				"post_count": 1,
				"updated_at": "2022-01-01T00:00:00Z",
				"posts": [
					{
						"created_at": "2020-01-01T00:00:00Z",
						"description": "post 1 description",
						"id": 1,
						"pools": null,
						"updated_at": "2020-01-01T00:00:00Z"
					}
				]
			}
		}`,
		string(b))
}

func TestPoolServiceList(t *testing.T) {
	poolRepository := mocks.NewMockPoolRepository(t)

	poolService := services.NewPoolService(services.PoolRepositoryConfig{
		PoolRepository: poolRepository,
	})

	poolRepository.On(
		"ListFull",
		context.TODO(),
		repositories.PoolListFullArgs{
			Text:     "text",
			Page:     1,
			PageSize: 100,
		},
	).Return(
		[]models.Pool{
			fakes.LoadPool(fakes.Pool1),
		},
		1,
		nil,
	)

	response, err := poolService.List(
		context.TODO(),
		dtos.ListPoolDTO{
			Search:   "text",
			Page:     1,
			PageSize: 100,
		},
	)
	require.NoError(t, err)

	b, err := json.Marshal(response)
	require.NoError(t, err)

	require.JSONEq(t,
		`{
			"count": 1,
			"pools": [
				{
					"created_at": "2022-01-01T00:00:00Z",
					"custom": ["a"],
					"description": "pool 1 description",
					"id": 1,
					"name": "pool 1 name",
					"post_count": 1,
					"updated_at": "2022-01-01T00:00:00Z",
					"posts": [
						{
							"created_at": "2020-01-01T00:00:00Z",
							"description": "post 1 description",
							"id": 1,
							"pools": null,
							"updated_at": "2020-01-01T00:00:00Z"
						}
					]
				}
			]
		}`,
		string(b),
	)
}