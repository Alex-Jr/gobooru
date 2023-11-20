package main

import (
	"gobooru/internal/controllers"
	"gobooru/internal/database"
	"gobooru/internal/middlewares"
	"gobooru/internal/repositories"
	"gobooru/internal/routes"
	"gobooru/internal/services"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	sqlx := database.MustGetSQLXConnection(database.GetConfig())

	db := database.NewSQLClient(
		sqlx,
	)

	database.MustRunMigrations(
		sqlx,
	)

	poolRepository := repositories.NewPoolRepository(db)
	postRepository := repositories.NewPostRepository(db)

	poolService := services.NewPoolService(services.PoolServiceConfig{
		PoolRepository: poolRepository,
	})
	postService := services.NewPostService(services.PostServiceConfig{
		PostRepository: postRepository,
	})

	healthCheckController := controllers.NewHealthCheckController()
	poolController := controllers.NewPoolController(controllers.PoolControllerConfig{
		PoolService: poolService,
	})
	postController := controllers.NewPostController(controllers.PostControllerConfig{
		PostService: postService,
	})

	r := echo.New()

	r.Use(middlewares.NewLoggerMiddleware())

	routes.RegisterHealthCheckRoutes(r, healthCheckController)
	routes.RegisterPoolRoutes(r, poolController)
	routes.RegisterPostRoutes(r, postController)

	for _, route := range r.Routes() {
		log.Printf("%s %s %s", route.Method, route.Path, route.Name)
	}

	r.Start(":8080")
}
