package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yangirxd/store-app/auth/api/dto"
	_ "github.com/yangirxd/store-app/auth/docs"
	"github.com/yangirxd/store-app/auth/service"
	"net/http"
)

// @Summary Register a new user
// @Description Register a new user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dto.RegisterRequest true "Register request"
// @Success 201 {object} domain.User
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/v1/register [post]
func registerHandler(authService *service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := authService.Register(req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, user)
	}
}

// @Summary Login user
// @Description Login with email and password to get JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dto.LoginRequest true "Login request"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/v1/login [post]
func loginHandler(authService *service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := authService.Login(req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
