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
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPostServiceCreate(t *testing.T) {
	postRepository := mocks.NewMockPostRepository(t)
	fileService := mocks.NewMockFileService(t)
	iqdbService := mocks.NewMockIQDBService(t)

	postService := services.NewPostService(services.PostServiceConfig{
		PostRepository: postRepository,
		FileService:    fileService,
		IQDBService:    iqdbService,
	})

	mockedPost := fakes.LoadPostRelations(fakes.Post1)
	mockedPost.Pools = nil

	mockedFileHeader := &multipart.FileHeader{}

	fileService.On(
		"HandleUpload",
		mockedFileHeader,
	).Return(
		models.File{
			MD5:       "1",
			FileSize:  100,
			FileExt:   "jpg",
			FilePath:  "1.jpg",
			ThumbPath: "1-thumb.webp",
		},
		nil,
	)

	postRepository.On(
		"Create",
		context.TODO(),
		repositories.CreatePostArgs{
			Description: "post 1 description",
			Rating:      "S",
			Tags:        []string{"tag_one"},
			FileExt:     "jpg",
			FileSize:    100,
			FilePath:    "1.jpg",
			MD5:         "1",
			ThumbPath:   "1-thumb.webp",
		},
	).Return(
		mockedPost,
		nil,
	)

	var mockedPostRelation []models.PostRelation

	copy(mockedPostRelation, mockedPost.Relations)

	iqdbService.On(
		"HandlePost",
		mockedPost,
	).Return(
		mockedPostRelation,
		nil,
	)

	postRepository.On(
		"SaveRelations",
		context.TODO(),
		&mockedPost,
		&mockedPostRelation,
	).Return(
		nil,
	)

	response, err := postService.Create(context.TODO(), dtos.CreatePostDTO{
		Description: "post 1 description",
		Rating:      "S",
		Tags:        []string{"tag_one"},
		File:        mockedFileHeader,
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
				"md5": "1",
				"file_size": 100,
				"file_ext": "jpg",
				"file_path": "1.jpg",
				"thumb_path": "1-thumb.webp",
				"relations": [
					{
						"post_id": 1,
						"other_post_id": 2,
						"type": "SIMILAR",
						"similarity": 9999,
						"created_at": "2020-01-01T00:00:00Z",
						"other_post": {
							"created_at": "2020-01-01T00:00:00Z",
							"description": "post 2 description",
							"id": 2,
							"pool_count": 3,
							"rating": "S",
							"tag_count": 0,
							"tag_ids": [],
							"md5": "2",
							"file_size": 100,
							"file_ext": "jpg",
							"file_path": "2.jpg",
							"thumb_path": "2-thumb.webp",
							"relations": null,
							"tags": null,
							"updated_at": "2020-01-01T00:00:00Z",
							"pools": null
						}
					}
				],
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
				"rating": "S",
				"tag_count": 1,
				"tag_ids": ["tag_one"],
				"md5": "1",
				"file_size": 100,
				"file_ext": "jpg",
				"file_path": "1.jpg",
				"thumb_path": "1-thumb.webp",
				"relations":[
					{
						"post_id": 1,
						"other_post_id": 2,
						"type": "SIMILAR",
						"similarity": 9999,
						"created_at": "2020-01-01T00:00:00Z",
						"other_post": {
							"created_at": "2020-01-01T00:00:00Z",
							"description": "post 2 description",
							"id": 2,
							"pool_count": 3,
							"rating": "S",
							"tag_count": 0,
							"tag_ids": [],
							"md5": "2",
							"file_size": 100,
							"file_ext": "jpg",
							"file_path": "2.jpg",
							"thumb_path": "2-thumb.webp",
							"relations": null,
							"tags": null,
							"updated_at": "2020-01-01T00:00:00Z",
							"pools": null
						}
					}
				],
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
				"rating": "S",
				"tag_count": 1,
				"tag_ids": ["tag_one"],
				"md5": "1",
				"file_size": 100,
				"file_ext": "jpg",
				"file_path": "1.jpg",
				"thumb_path": "1-thumb.webp",
				"relations": [
					{
						"post_id": 1,
						"other_post_id": 2,
						"type": "SIMILAR",
						"similarity": 9999,
						"created_at": "2020-01-01T00:00:00Z",
						"other_post": {
							"created_at": "2020-01-01T00:00:00Z",
							"description": "post 2 description",
							"id": 2,
							"pool_count": 3,
							"rating": "S",
							"tag_count": 0,
							"tag_ids": [],
							"md5": "2",
							"file_size": 100,
							"file_ext": "jpg",
							"file_path": "2.jpg",
							"thumb_path": "2-thumb.webp",
							"relations": null,
							"tags": null,
							"updated_at": "2020-01-01T00:00:00Z",
							"pools": null
						}
					}
				],
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
					"md5": "1",
					"file_size": 100,
					"file_ext": "jpg",
					"file_path": "1.jpg",
					"thumb_path": "1-thumb.webp",
					"tags": null, 
					"updated_at": "2020-01-01T00:00:00Z",
					"relations": null
				}
			]
		}
		`,
		string(b),
	)
}
