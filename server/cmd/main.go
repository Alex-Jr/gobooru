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

	createdPool, err := poolRepository.Create(context.TODO(), repositories.PoolCreateArgs{
		Custom:      []string{"asd"},
		Description: "",
		Name:        "test creation 1",
		Posts:       []int{2, 1},
	})

	if err != nil {
		panic(err)
	}

	s, _ := json.MarshalIndent(createdPool, "", "\t")
	fmt.Print(string(s))

	// fmt.Print(count)
}
