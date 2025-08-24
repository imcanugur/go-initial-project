package controller

import (
	"go-initial-project/entity"
	"go-initial-project/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	*BaseController[entity.User]
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		BaseController: NewBaseController[entity.User](userService),
		userService:    userService,
	}
}

func (uc *UserController) RegisterRoutes(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.GET("", uc.GetAll)
		users.GET("/:id", uc.GetByID)
		users.POST("", uc.Create)
		users.PUT("/:id", uc.Update)
		users.DELETE("/:id", uc.Delete)
	}
}

// GetUsers godoc
// @Summary Get all users
// @Tags users
// @Produce json
// @Success 200 {array} entity.User
// @Router /users [get]
func (uc *UserController) GetUsers(ctx *gin.Context) {
	uc.BaseController.GetAll(ctx)
}

// GetUserByID godoc
// @Summary Get user by ID
// @Tags users
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} entity.User
// @Router /users/{id} [get]
func (uc *UserController) GetUserByID(ctx *gin.Context) {
	uc.BaseController.GetByID(ctx)
}

// CreateUser godoc
// @Summary Create user
// @Tags users
// @Accept json
// @Produce json
// @Param data body entity.User true "User"
// @Success 201 {object} entity.User
// @Router /users [post]
func (uc *UserController) CreateUser(ctx *gin.Context) {
	uc.BaseController.Create(ctx)
}

// UpdateUser godoc
// @Summary Update user
// @Tags users
// @Accept json
// @Produce json
// @Param data body entity.User true "User"
// @Success 200 {object} entity.User
// @Router /users [put]
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	uc.BaseController.Update(ctx)
}

// DeleteUser godoc
// @Summary Delete user
// @Tags users
// @Param id path int true "ID"
// @Success 204
// @Router /users/{id} [delete]
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	uc.BaseController.Delete(ctx)
}
