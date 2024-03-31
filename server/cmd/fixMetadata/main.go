package main

import (
	"context"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/ffmpeg"
	"gobooru/internal/repositories"
)

func main() {
	sqlx := database.MustGetSQLXConnection(database.DBConfig{
		Host:     "localhost",
		Port:     "5450",
		User:     "user",
		Password: "password",
		Database: "database",
	})

	db := database.NewSQLClient(
		sqlx,
	)

	ffmpeg := ffmpeg.NewFfmpegModule()
	postRepository := repositories.NewPostRepository(
		db,
	)

	i := 16
	for {
		posts, _, err := postRepository.List(context.TODO(), repositories.ListPostsArgs{
			Search:   "width:0..0",
			Page:     i,
			PageSize: 1000,
		})
		if err != nil {
			panic(err)
		}

		if len(posts) == 0 {
			break
		}

		for _, post := range posts {
			if post.FilePath == "/web/images/no-file.webp" {
				fmt.Printf("Skipped post %d\n", post.ID)
				continue
			}

			fileMetadata, err := ffmpeg.ExtractMetadata(fmt.Sprintf("../../gobooru/docker-data/server/%s", post.FilePath))
			if err != nil {
				panic(err)
			}

			_, err = postRepository.Update(context.TODO(), repositories.UpdatePostArgs{
				ID:       post.ID,
				Width:    &fileMetadata.Width,
				Height:   &fileMetadata.Height,
				Duration: &fileMetadata.Duration,
			})
			if err != nil {
				panic(err)
			}

			fmt.Printf("Updated post %d\n", post.ID)
		}

		i++
	}
}
