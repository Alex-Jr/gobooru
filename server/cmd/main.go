package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/repositories"
)

func main() {
	db := database.NewSQLClient(
		database.GetSQLXConnection(database.GetDevConfig()),
	)

	database.MustRunMigrations(
		database.GetSQLConnection(database.GetDevConfig()),
	)

	poolRepository := repositories.NewPoolRepository(db)

	pool, _, err := poolRepository.ListFull(
		context.TODO(),
		repositories.PoolListFullArgs{
			Text:     "2",
			Page:     1,
			PageSize: 100,
		},
	)

	if err != nil {
		panic(err)
	}

	s, _ := json.MarshalIndent(pool, "", "\t")
	fmt.Print(string(s))

	// fmt.Print(count)
}
