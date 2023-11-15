package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/repositories"
)

func main() {
	sqlx := database.MustGetSQLXConnection(database.GetDevConfig())

	db := database.NewSQLClient(
		sqlx,
	)

	database.MustRunMigrations(
		sqlx,
	)

	poolRepository := repositories.NewPoolRepository(db)

	pool, count, err := poolRepository.ListFull(context.TODO(), repositories.PoolListFullArgs{
		Text:     "custom:test",
		Page:     1,
		PageSize: 10,
	})

	if err != nil {
		panic(err)
	}

	s, _ := json.MarshalIndent(pool, "", "\t")
	fmt.Print(string(s))
	fmt.Print(count)

	// fmt.Print(count)
}
