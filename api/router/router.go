package router

import (
	middlewares2 "eden/internal/pkg/middlewares"
	"os"

	"eden/api/v1/controller"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	app := gin.New()

	// logging to file
	// gin.ForceConsoleColor()
	// log.SetFormatter(&log.TextFormatter{ForceColors: true})
	// log.SetOutput(os.Stdout)

	// Logging
	gin.ForceConsoleColor()
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(os.Stdout)

	// Middlewares
	app.Use(middlewares2.HttpLogging())
	app.Use(gin.Recovery())
	app.Use(middlewares2.CORS())
	app.NoRoute(middlewares2.NoRouteHandler())
	// TODO: come here and use library that I'm creating for wanna taste
	// authMiddleware := middlewares.JwtAuth()

	// Routes
	// Docs Routes
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// setting api groupings
	api := app.Group("/api")
	versionOne := api.Group("/v1")

	// Test Route
	versionOne.GET("/ping", controller.Ping)

	// User Route
	//user := versionOne.Group("/user")
	// user.Use(authMiddleware.MiddlewareFunc())
	// {
	//user.GET("/:id", controller.GetUserById)
	//user.POST("", controller.CreateUser)
	// }

	//workout := versionOne.Group("/workout")
	//workout.Use(middlewares.UserContext())
	//{
	//	workout.GET("/:id", controller.GetWorkoutById)
	//	workout.POST("", controller.CreateWorkout)
	//	workout.GET("/:id/details", controller.GetWorkoutDetailsById)
	//}

	// Auth / login
	// auth := app.Group("/auth")
	// auth.POST("/login", authMiddleware.LoginHandler)
	// auth.GET("/refresh", authMiddleware.RefreshHandler)
	// auth.GET("/logout", authMiddleware.LogoutHandler)

	return app
}
