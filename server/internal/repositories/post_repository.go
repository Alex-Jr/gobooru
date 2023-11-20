package repositories

import (
	"context"
	"gobooru/internal/database"
	"gobooru/internal/models"
	"gobooru/internal/queries"
)

type PostRepository interface {
	Create(ctx context.Context, args CreatePostArgs) (models.Post, error)
}

type postRepository struct {
	sqlClient database.SQLClient
	postQuery queries.PostQuery
}

type CreatePostArgs struct {
	Description string
}

func NewPostRepository(sqlClient database.SQLClient) PostRepository {
	return &postRepository{
		sqlClient: sqlClient,
		postQuery: queries.NewPostQuery(),
	}
}

func (r *postRepository) Create(ctx context.Context, args CreatePostArgs) (models.Post, error) {
	post := models.Post{
		Description: args.Description,
	}

	err := r.postQuery.Create(ctx, r.sqlClient, &post)

	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}
