package main

import (
	"go-initial-project/config"
	"go-initial-project/controller"
	docs "go-initial-project/docs"
	"go-initial-project/repository"
	"go-initial-project/router"
	"go-initial-project/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title My API
// @version 1.0
// @description This is a sample server.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	config.LoadEnv()

	db := config.ConnectDB()

	userRepo := repository.NewUserRepository(db)
	activityRepo := repository.NewActivityRepository(db)

	userService := service.NewUserService(userRepo)
	activityService := service.NewActivityService(activityRepo)

	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(userService)

	// Router
	r := router.SetupRouter(activityService, userController, authController)

	docs.SwaggerInfo.BasePath = "/api"

	r.GET("/swagger", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + config.AppConfig.App.Port)
}
