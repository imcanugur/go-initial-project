package controller

import (
	"go-initial-project/config"
	"go-initial-project/entity"
	"go-initial-project/middleware"
	authreq "go-initial-project/requests/auth"
	authres "go-initial-project/responses/auth"
	userres "go-initial-project/responses/user"
	"go-initial-project/service"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	userService *service.UserService
}

func NewAuthController(userService *service.UserService) *AuthController {
	return &AuthController{userService: userService}
}

func (ac *AuthController) RegisterRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", ac.Login)
		auth.POST("/register", ac.Register)
		auth.GET("/me", middleware.AuthRequired(), ac.Me)
	}
}

// Login godoc
// @Summary Login user
// @Description Authenticate user with email and password, return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param data body auth.LoginRequest true "Login data"
// @Success 200 {object} auth.AuthResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /auth/login [post]
func (ac *AuthController) Login(ctx *gin.Context) {
	var req authreq.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	user, err := ac.userService.FindByEmail(req.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString(config.JWTSecret())

	res := authres.AuthResponse{
		Token: tokenString,
		User: userres.UserResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		},
	}
	ctx.JSON(http.StatusOK, res)
}

// Register godoc
// @Summary Register user
// @Description Create a new user account and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param data body auth.RegisterRequest true "Register data"
// @Success 201 {object} auth.AuthResponse
// @Failure 400 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/register [post]
func (ac *AuthController) Register(ctx *gin.Context) {
	var req authreq.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	user := entity.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}

	createdUser, err := ac.userService.Create(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString(config.JWTSecret())

	res := authres.AuthResponse{
		Token: tokenString,
		User: userres.UserResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		},
	}
	ctx.JSON(http.StatusCreated, res)
}

// Me godoc
// @Summary Get current user
// @Description Return current authenticated user info
// @Tags auth
// @Security BearerAuth
// @Produce json
// @Success 200 {object} user.UserResponse
// @Failure 401 {object} map[string]string
// @Router /auth/me [get]
func (ac *AuthController) Me(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user, err := ac.userService.First(map[string]interface{}{"id": userID.(string)})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
	})
}
