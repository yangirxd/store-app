package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yangirxd/store-app/orders/domain"
	"net/http"
	"strings"
)

func OrderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			c.Abort()
			return
		}

		claims, err := domain.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		fmt.Printf("Extracted email from token: %s\n", claims.Email)

		if claims.Email == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "email is empty in token"})
			c.Abort()
			return
		}

		c.Set("userEmail", claims.Email)
		c.Next()
	}
}
