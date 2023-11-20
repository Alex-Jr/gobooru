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
				"pool_count": 0,
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
