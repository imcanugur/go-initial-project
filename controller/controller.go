package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	RegisterRoutes(r *gin.RouterGroup)
}
