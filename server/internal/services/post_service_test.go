package services_test

import (
	"context"
	"encoding/json"
	"gobooru/internal/dtos"
	"gobooru/internal/fixtures/fakes"
	"gobooru/internal/mocks"
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

	postRepository.On(
		"Create",
		context.TODO(),
		repositories.CreatePostArgs{
			Description: "post 1 description",
		},
	).Return(
		fakes.Post1,
		nil,
	)

	response, err := postService.Create(context.TODO(), dtos.CreatePostDTO{
		Description: "post 1 description",
	})
	require.NoError(t, err)

	b, err := json.Marshal(response)
	require.NoError(t, err)

	require.JSONEq(t, `
		{
			"post": {
				"created_at": "2020-01-01T00:00:00Z",
				"description": "post 1 description",
				"id": 1,
				"pools": null,
				"updated_at": "2020-01-01T00:00:00Z"
			}
		}
		`,
		string(b),
	)
}
