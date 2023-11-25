package main

import (
	"gobooru/internal/controllers"
	"gobooru/internal/database"
	"gobooru/internal/ffmpeg"
	"gobooru/internal/middlewares"
	"gobooru/internal/repositories"
	"gobooru/internal/routes"
	"gobooru/internal/services"
	"log"
	"os"

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

	ffmpegModule := ffmpeg.NewFfmpegModule()

	poolRepository := repositories.NewPoolRepository(db)
	postRepository := repositories.NewPostRepository(db)

	fileService := services.NewFileService(services.FileServiceConfig{
		FFMPEGModule: ffmpegModule,
		BASE_PATH:    os.Getenv("STATIC_PATH"),
	})
	poolService := services.NewPoolService(services.PoolServiceConfig{
		PoolRepository: poolRepository,
	})
	postService := services.NewPostService(services.PostServiceConfig{
		PostRepository: postRepository,
		FileService:    fileService,
	})

	healthCheckController := controllers.NewHealthCheckController()
	poolController := controllers.NewPoolController(controllers.PoolControllerConfig{
		PoolService: poolService,
	})
	postController := controllers.NewPostController(controllers.PostControllerConfig{
		PostService: postService,
	})

	e := echo.New()

	e.Use(middlewares.NewLoggerMiddleware())

	e.Static("/static", os.Getenv("STATIC_PATH"))

	routes.RegisterHealthCheckRoutes(e, healthCheckController)
	routes.RegisterPoolRoutes(e, poolController)
	routes.RegisterPostRoutes(e, postController)

	for _, route := range e.Routes() {
		log.Printf("%s %s %s", route.Method, route.Path, route.Name)
	}

	e.Start(":8080")
}
