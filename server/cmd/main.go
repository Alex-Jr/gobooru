package main

import (
	"gobooru/internal/controllers"
	"gobooru/internal/database"
	"gobooru/internal/repositories"
	"gobooru/internal/routes"
	"gobooru/internal/services"
	"log"

	"github.com/labstack/echo/v4"
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

	poolService := services.NewPoolService(services.PoolRepositoryConfig{
		PoolRepository: poolRepository,
	})

	healthCheckController := controllers.NewHealthCheckController()
	poolController := controllers.NewPoolController(controllers.PoolControllerConfig{
		PoolService: poolService,
	})

	r := echo.New()

	routes.RegisterHealthCheckRoutes(r, healthCheckController)
	routes.RegisterPoolRoutes(r, poolController)

	for _, route := range r.Routes() {
		log.Printf("%s %s %s", route.Method, route.Path, route.Name)
	}

	r.Start(":8080")
}
