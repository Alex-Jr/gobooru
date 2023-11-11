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

	newName := "new name"
	pool, err := poolRepository.Update(
		context.TODO(),
		repositories.PoolUpdateArgs{
			ID:    2,
			Name:  &newName,
			Posts: &[]int{1, 2},
		},
	)

	if err != nil {
		panic(err)
	}

	s, _ := json.MarshalIndent(pool, "", "\t")
	fmt.Print(string(s))

	// fmt.Print(count)
}
