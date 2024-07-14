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
	"github.com/labstack/echo/v4/middleware"
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

	tagCategoryRepository := repositories.NewTagCategoryRepository(repositories.TagCategoryRepositoryConfig{
		SQLClient: db,
	})
	tagRepository := repositories.NewTagRepository(repositories.TagRepositoryConfig{
		SQLClient: db,
	})
	poolRepository := repositories.NewPoolRepository(db)
	postNotesRepository := repositories.NewPostNotesRepository(db)
	postRepository := repositories.NewPostRepository(db)

	IQDBService := services.NewIQDBService(services.IQDBServiceConfig{
		IQDB_URL:  os.Getenv("IQDB_URL"),
		BASE_PATH: os.Getenv("STATIC_PATH"),
	})
	fileService := services.NewFileService(services.FileServiceConfig{
		FFMPEGModule: ffmpegModule,
		BASE_PATH:    os.Getenv("STATIC_PATH"),
	})
	tagCategoryService := services.NewTagCategoryService(services.TagCategoryServiceConfig{
		TagCategoryRepository: tagCategoryRepository,
	})
	tagService := services.NewTagService(services.TagServiceConfig{
		TagRepository: tagRepository,
	})
	poolService := services.NewPoolService(services.PoolServiceConfig{
		PoolRepository: poolRepository,
	})
	postService := services.NewPostService(services.PostServiceConfig{
		PostRepository:      postRepository,
		PostNotesRepository: postNotesRepository,
		FileService:         fileService,
		IQDBService:         IQDBService,
	})

	healthCheckController := controllers.NewHealthCheckController()
	tagController := controllers.NewTagController(controllers.TagControllerConfig{
		TagService: tagService,
	})
	tagCategoryController := controllers.NewTagCategoryController(controllers.TagCategoryControllerConfig{
		TagCategoryService: tagCategoryService,
	})
	poolController := controllers.NewPoolController(controllers.PoolControllerConfig{
		PoolService: poolService,
	})
	postController := controllers.NewPostController(controllers.PostControllerConfig{
		PostService: postService,
	})

	e := echo.New()

	// TODO: Move to a middleware and should not print all errors
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		c.JSON(500, map[string]interface{}{
			"error": err.Error(),
		})
	}

	e.Use(middlewares.NewLoggerMiddleware())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{os.Getenv("CLIENT_URL")},
		AllowCredentials: true,
	}))

	e.Static("/static", os.Getenv("STATIC_PATH"))

	routes.RegisterHealthCheckRoutes(e, healthCheckController)
	routes.RegisterPoolRoutes(e, poolController)
	routes.RegisterPostRoutes(e, postController)
	routes.RegisterTagCategoryRoutes(e, tagCategoryController)
	routes.RegisterTagRoutes(e, tagController)

	for _, route := range e.Routes() {
		log.Printf("%s %s %s", route.Method, route.Path, route.Name)
	}

	e.Start(":8080")
}
