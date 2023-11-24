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

func TestPostServiceCreate(t *testing.T) {
	postRepository := mocks.NewMockPostRepository(t)

	postService := services.NewPostService(services.PostServiceConfig{
		PostRepository: postRepository,
	})

	mockedPost := fakes.Post1
	mockedPost.Pools = nil
	mockedPost.Tags = []models.Tag{
		fakes.Tag1,
	}

	postRepository.On(
		"Create",
		context.TODO(),
		repositories.CreatePostArgs{
			Description: "post 1 description",
			Rating:      "S",
			Tags:        []string{"tag_one"},
		},
	).Return(
		mockedPost,
		nil,
	)

	response, err := postService.Create(context.TODO(), dtos.CreatePostDTO{
		Description: "post 1 description",
		Rating:      "S",
		Tags:        []string{"tag_one"},
	})
	require.NoError(t, err)

	b, err := json.Marshal(response)
	require.NoError(t, err)

	require.JSONEq(t, `
		{
			"post":		{
				"created_at": "2020-01-01T00:00:00Z",
				"description": "post 1 description",
				"id": 1,
				"pool_count": 4,
				"pools": null,
				"rating": "S",
				"tag_count": 1,
				"tag_ids": ["tag_one"],
				"tags": [
					{
						"id": "tag_one",
						"post_count": 1,
						"created_at": "2020-01-01T00:00:00Z",
						"updated_at": "2020-01-01T00:00:00Z",
						"description": "tag one description"
					}
				],
				"updated_at": "2020-01-01T00:00:00Z"
			}
		}
		`,
		string(b),
	)
}

func TestPostServiceDelete(t *testing.T) {
	postRepository := mocks.NewMockPostRepository(t)

	postService := services.NewPostService(services.PostServiceConfig{
		PostRepository: postRepository,
	})

	postRepository.On(
		"GetFull",
		context.TODO(),
		1,
	).Return(
		fakes.LoadPostRelations(fakes.Post1),
		nil,
	)

	postRepository.On(
		"Delete",
		context.TODO(),
		1,
	).Return(
		nil,
	)

	response, err := postService.Delete(context.TODO(), dtos.DeletePostDTO{
		ID: 1,
	})
	require.NoError(t, err)

	b, err := json.Marshal(response)
	require.NoError(t, err)

	require.JSONEq(t, `
		{
			"post":		{
				"created_at": "2020-01-01T00:00:00Z",
				"description": "post 1 description",
				"id": 1,
				"pool_count": 4,
				"pools": [
					{
						"created_at": "2022-01-01T00:00:00Z",
						"custom": ["a"],
						"description": "pool 1 description",
						"id": 1,
						"name": "pool 1 name",
						"post_count": 1,
						"posts": null,
						"updated_at": "2022-01-01T00:00:00Z"
					},
					{
						"created_at": "2022-01-01T00:00:00Z",
						"custom": [],
						"description": "pool 2 description",
						"id": 2,
						"name": "pool 2 name",
						"post_count": 2,
						"posts": null,
						"updated_at": "2022-01-01T00:00:00Z"
					},
					{
						"created_at": "2020-01-01T00:00:00Z",
						"custom": [],
						"description": "pool 3 description",
						"id": 3,
						"name": "pool 3 name",
						"post_count": 3,
						"posts": null,
						"updated_at": "2020-01-01T00:00:00Z"
					},
					{
						"created_at": "2020-01-01T00:00:00Z",
						"custom": ["shared"],
						"description": "pool 4 description",
						"id": 4,
						"name": "pool 4 name",
						"post_count": 1,
						"posts": null,
						"updated_at": "2020-01-01T00:00:00Z"
					}
				],
				"rating": "S",
				"tag_count": 1,
				"tag_ids": ["tag_one"],
				"tags": [
					{
						"id": "tag_one",
						"post_count": 1,
						"created_at": "2020-01-01T00:00:00Z",
						"updated_at": "2020-01-01T00:00:00Z",
						"description": "tag one description"
					}
				],
				"updated_at": "2020-01-01T00:00:00Z"
			}
		}
		`,
		string(b),
	)
}

func TestPostServiceFetch(t *testing.T) {
	postRepository := mocks.NewMockPostRepository(t)

	postService := services.NewPostService(services.PostServiceConfig{
		PostRepository: postRepository,
	})

	postRepository.On(
		"GetFull",
		context.TODO(),
		1,
	).Return(
		fakes.LoadPostRelations(fakes.Post1),
		nil,
	)

	response, err := postService.Fetch(context.TODO(), dtos.FetchPostDTO{
		ID: 1,
	})
	require.NoError(t, err)

	b, err := json.Marshal(response)
	require.NoError(t, err)

	require.JSONEq(t, `
		{
			"post":		{
				"created_at": "2020-01-01T00:00:00Z",
				"description": "post 1 description",
				"id": 1,
				"pool_count": 4,
				"pools": [
					{
						"created_at": "2022-01-01T00:00:00Z",
						"custom": ["a"],
						"description": "pool 1 description",
						"id": 1,
						"name": "pool 1 name",
						"post_count": 1,
						"posts": null,
						"updated_at": "2022-01-01T00:00:00Z"
					},
					{
						"created_at": "2022-01-01T00:00:00Z",
						"custom": [],
						"description": "pool 2 description",
						"id": 2,
						"name": "pool 2 name",
						"post_count": 2,
						"posts": null,
						"updated_at": "2022-01-01T00:00:00Z"
					},
					{
						"created_at": "2020-01-01T00:00:00Z",
						"custom": [],
						"description": "pool 3 description",
						"id": 3,
						"name": "pool 3 name",
						"post_count": 3,
						"posts": null,
						"updated_at": "2020-01-01T00:00:00Z"
					},
					{
						"created_at": "2020-01-01T00:00:00Z",
						"custom": ["shared"],
						"description": "pool 4 description",
						"id": 4,
						"name": "pool 4 name",
						"post_count": 1,
						"posts": null,
						"updated_at": "2020-01-01T00:00:00Z"
					}
				],
				"rating": "S",
				"tag_count": 1,
				"tag_ids": ["tag_one"],
				"tags": [
					{
						"id": "tag_one",
						"post_count": 1,
						"created_at": "2020-01-01T00:00:00Z",
						"updated_at": "2020-01-01T00:00:00Z",
						"description": "tag one description"
					}
				],
				"updated_at": "2020-01-01T00:00:00Z"
			}
		}
		`,
		string(b),
	)
}

func TestPostServiceList(t *testing.T) {
	postRepository := mocks.NewMockPostRepository(t)

	postService := services.NewPostService(services.PostServiceConfig{
		PostRepository: postRepository,
	})

	postRepository.On(
		"List",
		context.TODO(),
		repositories.ListPostsArgs{
			Search:   "post",
			Page:     1,
			PageSize: 1,
		},
	).Return(
		[]models.Post{
			fakes.LoadPostNoRelations(fakes.Post1),
		},
		1,
		nil,
	)

	response, err := postService.List(context.TODO(), dtos.ListPostDTO{
		Search:   "post",
		Page:     1,
		PageSize: 1,
	})
	require.NoError(t, err)

	b, err := json.Marshal(response)
	require.NoError(t, err)

	require.JSONEq(t, `
		{
			"count": 1,
			"posts": [
				{
					"created_at": "2020-01-01T00:00:00Z",
					"description": "post 1 description",
					"id": 1,
					"pool_count": 4,
					"pools": null,
					"rating": "S",
					"tag_count": 1,
					"tag_ids": ["tag_one"],
					"tags": null, 
					"updated_at": "2020-01-01T00:00:00Z"
				}
			]
		}
		`,
		string(b),
	)
}
