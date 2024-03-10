package repositories

import (
	"context"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/models"
	"gobooru/internal/queries"
)

type PostNotesRepository interface {
	// Create creates a post note entry.
	Create(ctx context.Context, args CreatePostNotesArgs) (models.PostNote, error)
}

type postNotesRepository struct {
	sqlClient      database.SQLClient
	postNotesQuery queries.PostNotesQuery
}

type CreatePostNotesArgs struct {
	PostID int
	Body   string
	X      int
	Y      int
	Width  int
	Height int
}

func NewPostNotesRepository(db database.SQLClient) PostNotesRepository {
	return &postNotesRepository{
		sqlClient:      db,
		postNotesQuery: queries.NewPostNotesQuery(),
	}
}

func (r *postNotesRepository) Create(ctx context.Context, args CreatePostNotesArgs) (models.PostNote, error) {
	postNote := models.PostNote{
		PostID: args.PostID,
		Body:   args.Body,
		X:      args.X,
		Y:      args.Y,
		Width:  args.Width,
		Height: args.Height,
	}

	err := r.postNotesQuery.Create(ctx, r.sqlClient, &postNote)
	if err != nil {
		return models.PostNote{}, fmt.Errorf("postNotesQuery.Create: %w", err)
	}

	return postNote, nil
}
