package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/teguh522/payslip/cmd/internal/pkg/config"
)

type AuthMiddleware struct {
	jwtSecretKey []byte
}

func NewAuthMiddleware(cfg *config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		jwtSecretKey: []byte(cfg.App.JWTSecret),
	}
}

func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Unexpected signing method")
			}
			return m.jwtSecretKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			exp := int64(claims["exp"].(float64))
			if time.Now().Unix() > exp {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is expired"})
				c.Abort()
				return
			}

			c.Set("userID", claims["userId"])
			c.Set("userName", claims["userName"])
			c.Set("role", claims["role"])
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}
	}
}

func GetUserIDFromContext(c *gin.Context) (string, bool) {
	userID, ok := c.Get("userID")
	if !ok {
		return "", false
	}
	return userID.(string), true
}

func GetUserNameFromContext(c *gin.Context) (string, bool) {
	userName, ok := c.Get("userName")
	if !ok {
		return "", false
	}
	return userName.(string), true
}

func GetUserRoleFromContext(c *gin.Context) (string, bool) {
	role, ok := c.Get("role")
	if !ok {
		return "", false
	}
	return role.(string), true
}
