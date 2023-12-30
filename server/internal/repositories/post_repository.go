package repositories

import (
	"context"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/models"
	"gobooru/internal/queries"
	"gobooru/internal/slice_utils"
	"time"
)

type PostRepository interface {
	Create(ctx context.Context, args CreatePostArgs) (models.Post, error)
	Delete(ctx context.Context, postID int) error
	GetFull(ctx context.Context, postID int) (models.Post, error)
	GetFullByHash(ctx context.Context, hash string) (models.Post, error)
	List(ctx context.Context, args ListPostsArgs) ([]models.Post, int, error)
	SaveRelations(ctx context.Context, post *models.Post, relations *[]models.PostRelation) error
	Update(ctx context.Context, args UpdatePostArgs) (models.Post, error)
}

type postRepository struct {
	sqlClient         database.SQLClient
	postQuery         queries.PostQuery
	poolQuery         queries.PoolQuery
	tagQuery          queries.TagQuery
	postTag           queries.PostTagQuery
	postRelationQuery queries.PostRelationQuery
	tagCategoryQuery  queries.TagCategoryQuery
	tagAliasQuery     queries.TagAliasQuery
	tagImplication    queries.TagImplicationQuery
}

type CreatePostArgs struct {
	Custom      []string
	Description string
	FileExt     string
	FilePath    string
	FileSize    int
	MD5         string
	Rating      string
	Sources     []string
	Tags        []string
	ThumbPath   string
}

func NewPostRepository(sqlClient database.SQLClient) PostRepository {
	return &postRepository{
		sqlClient:         sqlClient,
		postQuery:         queries.NewPostQuery(),
		tagQuery:          queries.NewTagQuery(),
		postTag:           queries.NewPostTagQuery(),
		postRelationQuery: queries.NewPostRelationQuery(),
		tagCategoryQuery:  queries.NewTagCategoryQuery(),
		tagAliasQuery:     queries.NewTagAliasQuery(),
		tagImplication:    queries.NewTagImplicationQuery(),
		poolQuery:         queries.NewPoolQuery(),
	}
}

func (r *postRepository) Create(ctx context.Context, args CreatePostArgs) (models.Post, error) {
	tx, err := r.sqlClient.BeginTxx(ctx, nil)
	if err != nil {
		return models.Post{}, fmt.Errorf("sqlClient.BeginTxx: %w", err)
	}
	defer tx.Rollback()

	now := time.Now()

	tagsDeduped := slice_utils.Deduplicate(args.Tags)

	err = r.tagAliasQuery.ResolveAlias(ctx, tx, tagsDeduped)
	if err != nil {
		return models.Post{}, fmt.Errorf("tagAliasQuery.ResolveAlias: %w", err)
	}

	err = r.tagImplication.ResolveImplications(ctx, tx, &tagsDeduped)
	if err != nil {
		return models.Post{}, fmt.Errorf("tagImplication.ResolveImplications: %w", err)
	}

	tags := make([]models.Tag, len(tagsDeduped))

	post := models.Post{
		Custom:      args.Custom,
		Rating:      args.Rating,
		Description: args.Description,
		TagIDs:      make([]string, len(tags)),
		TagCount:    len(tags),
		PoolCount:   0,
		CreatedAt:   now,
		UpdatedAt:   now,
		Pools:       make(models.PoolList, 0),
		Tags:        make(models.TagList, len(tags)),
		MD5:         args.MD5,
		FileExt:     args.FileExt,
		FilePath:    args.FilePath,
		FileSize:    args.FileSize,
		ThumbPath:   args.ThumbPath,
		Sources:     args.Sources,
	}

	// TODO: probably there's a better way to handle nil slices
	if post.Custom == nil {
		post.Custom = make([]string, 0)
	}

	if post.Sources == nil {
		post.Sources = make([]string, 0)
	}

	for i, tag := range tagsDeduped {
		tags[i] = models.Tag{
			ID:          tag,
			PostCount:   1,
			CreatedAt:   now,
			UpdatedAt:   now,
			Description: "",
			CategoryID:  "general",
		}

		post.TagIDs[i] = tag
	}

	post.Tags = tags

	err = r.postQuery.Create(ctx, tx, &post)

	if err != nil {
		return models.Post{}, fmt.Errorf("postQuery.Create: %w", err)
	}

	err = r.tagQuery.CreateMany(ctx, tx, &tags)

	if err != nil {
		return models.Post{}, fmt.Errorf("tagQuery.CreateMany: %w", err)
	}

	err = r.postTag.AssociatePosts(ctx, tx, post, tags)
	if err != nil {
		return models.Post{}, fmt.Errorf("postTag.AssociatePosts: %w", err)
	}

	createdTagsCount := 0
	for _, tag := range tags {
		if now.Before(tag.CreatedAt) {
			createdTagsCount++
		}
	}

	//	TODO: don't use magic string here
	err = r.tagCategoryQuery.UpdateTagCount(ctx, tx, "general", createdTagsCount)
	if err != nil {
		return models.Post{}, fmt.Errorf("tagCategoryQuery.UpdateTagCount: %w", err)
	}

	tx.Commit()

	return post, nil
}

func (r *postRepository) Delete(ctx context.Context, postID int) error {
	tx, err := r.sqlClient.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("sqlClient.BeginTxx: %w", err)
	}
	defer tx.Rollback()

	post := models.Post{
		ID: postID,
	}

	err = r.postQuery.Delete(ctx, tx, &post)

	if err != nil {
		return fmt.Errorf("postQuery.Delete: %w", err)
	}

	err = r.tagQuery.UpdatePostCount(ctx, tx, post.TagIDs, -1)
	if err != nil {
		return fmt.Errorf("tagQuery.UpdatePostCount: %w", err)
	}

	err = r.poolQuery.RemovePost(ctx, tx, post.ID)
	if err != nil {
		return fmt.Errorf("poolQuery.RemovePost: %w", err)
	}

	tx.Commit()

	return nil
}

func (r *postRepository) GetFull(ctx context.Context, postID int) (models.Post, error) {
	post := models.Post{
		ID: postID,
	}

	err := r.postQuery.GetFull(ctx, r.sqlClient, &post)

	if err != nil {
		return models.Post{}, fmt.Errorf("postQuery.GetFull: %w", err)
	}

	return post, nil
}

func (r *postRepository) GetFullByHash(ctx context.Context, hash string) (models.Post, error) {
	post := models.Post{
		MD5: hash,
	}

	err := r.postQuery.GetFullByHash(ctx, r.sqlClient, &post)

	if err != nil {
		return models.Post{}, fmt.Errorf("postQuery.GetFullByHash: %w", err)
	}

	return post, nil
}

type ListPostsArgs struct {
	Search   string
	Page     int
	PageSize int
}

func (r *postRepository) List(ctx context.Context, args ListPostsArgs) ([]models.Post, int, error) {
	posts := make([]models.Post, 0)
	count := 0

	err := r.postQuery.List(
		ctx,
		r.sqlClient,
		models.Search{
			Text:     args.Search,
			Page:     args.Page,
			PageSize: args.PageSize,
		},
		&posts,
		&count,
	)

	if err != nil {
		return nil, 0, fmt.Errorf("postQuery.List: %w", err)
	}

	return posts, count, nil
}

func (r *postRepository) SaveRelations(ctx context.Context, post *models.Post, relations *[]models.PostRelation) error {
	relationsToCreate := make([]models.PostRelation, 0, len(*relations)*2)

	for _, relation := range *relations {
		relationsToCreate = append(relationsToCreate, models.PostRelation{
			PostID:      post.ID,
			OtherPostID: relation.OtherPostID,
			Similarity:  relation.Similarity,
			Type:        "SIMILAR",
		})

		relationsToCreate = append(relationsToCreate, models.PostRelation{
			PostID:      relation.OtherPostID,
			OtherPostID: post.ID,
			Similarity:  relation.Similarity,
			Type:        "SIMILAR",
		})
	}

	err := r.postRelationQuery.InsertRelations(ctx, r.sqlClient, *post, relationsToCreate)
	if err != nil {
		return fmt.Errorf("postQuery.InsertRelations: %w", err)
	}

	for i := range *relations {
		otherPost := models.Post{
			ID: (*relations)[i].OtherPostID,
		}

		r.postQuery.GetFull(ctx, r.sqlClient, &otherPost)

		(*relations)[i].OtherPost = otherPost
	}

	post.Relations = *relations

	return nil
}

type UpdatePostArgs struct {
	ID          int
	Description *string
	Rating      *string
	Tags        *[]string
	Sources     *[]string
	Custom      *[]string
}

func (r *postRepository) Update(ctx context.Context, args UpdatePostArgs) (models.Post, error) {
	tx, err := r.sqlClient.BeginTxx(ctx, nil)
	if err != nil {
		return models.Post{}, fmt.Errorf("sqlClient.BeginTxx: %w", err)
	}

	post := models.Post{
		ID: args.ID,
	}

	err = r.postQuery.GetFull(ctx, tx, &post)
	if err != nil {
		return models.Post{}, fmt.Errorf("postQuery.GetFull: %w", err)
	}

	if args.Description != nil {
		post.Description = *args.Description
	}

	if args.Rating != nil {
		post.Rating = *args.Rating
	}

	if args.Tags != nil {
		tagsDeduped := slice_utils.Deduplicate(*args.Tags)

		err = r.tagAliasQuery.ResolveAlias(ctx, tx, tagsDeduped)
		if err != nil {
			return models.Post{}, fmt.Errorf("tagAliasQuery.ResolveAlias: %w", err)
		}

		err = r.tagImplication.ResolveImplications(ctx, tx, &tagsDeduped)
		if err != nil {
			return models.Post{}, fmt.Errorf("tagImplication.ResolveImplications: %w", err)
		}

		toRemove := slice_utils.Difference(post.TagIDs, tagsDeduped)
		toAdd := slice_utils.Difference(tagsDeduped, post.TagIDs)

		if len(toRemove) > 0 {
			err = r.postTag.DisassociatePostsByID(ctx, tx, post, toRemove)
			if err != nil {
				return models.Post{}, fmt.Errorf("postTag.DisassociatePosts: %w", err)
			}

			err = r.tagQuery.UpdatePostCount(ctx, tx, toRemove, -1)
			if err != nil {
				return models.Post{}, fmt.Errorf("tagQuery.UpdatePostCount: %w", err)
			}
		}

		if len(toAdd) > 0 {
			tags := make([]models.Tag, len(toAdd))

			now := time.Now()
			for i, tag := range toAdd {
				tags[i] = models.Tag{
					ID:          tag,
					PostCount:   1,
					Description: "",
					CategoryID:  "general",
					CreatedAt:   now,
					UpdatedAt:   now,
				}
			}

			err = r.tagQuery.CreateMany(ctx, tx, &tags)
			if err != nil {
				return models.Post{}, fmt.Errorf("tagQuery.CreateMany: %w", err)
			}

			err = r.postTag.AssociatePosts(ctx, tx, post, tags)
			if err != nil {
				return models.Post{}, fmt.Errorf("postTag.AssociatePosts: %w", err)
			}

			createdTagsCount := 0
			for _, tag := range tags {
				if now.Before(tag.CreatedAt) {
					createdTagsCount++
				}
			}

			//	TODO: don't use magic string here
			err = r.tagCategoryQuery.UpdateTagCount(ctx, tx, "general", createdTagsCount)
			if err != nil {
				return models.Post{}, fmt.Errorf("tagCategoryQuery.UpdateTagCount: %w", err)
			}
		}

		post.TagIDs = *args.Tags

		post.TagCount = len(*args.Tags)
	}

	if args.Sources != nil {
		post.Sources = *args.Sources
	}

	if args.Custom != nil {
		post.Custom = *args.Custom
	}

	post.UpdatedAt = time.Now()

	err = r.postQuery.Update(ctx, tx, post)
	if err != nil {
		return models.Post{}, fmt.Errorf("postQuery.Update: %w", err)
	}

	err = r.postQuery.GetFull(ctx, tx, &post)
	if err != nil {
		return models.Post{}, fmt.Errorf("postQuery.GetFull: %w", err)
	}

	tx.Commit()

	return post, nil
}
