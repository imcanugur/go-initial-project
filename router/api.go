package router

import (
	"go-initial-project/controller"
	"go-initial-project/middleware"
	"go-initial-project/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(activityService *service.ActivityService, controllers ...controller.Controller) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")

	// Tek middleware burada
	r.Use(middleware.ActivityLogger(activityService))
	api.Use(middleware.ActivityLogger(activityService)) // sadece burada

	for _, c := range controllers {
		c.RegisterRoutes(api)

	}

	return r
}
