package middleware

import (
	"github.com/gin-gonic/gin"
)

func (m *AuthMiddleware) RoleAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(403, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}
		if role == "admin" {
			c.Next()
			return
		} else {
			c.JSON(403, gin.H{"error": "Forbidden, only admin can access this resource"})
			c.Abort()
			return
		}
	}
}

func (m *AuthMiddleware) RoleUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(403, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}
		if role == "user" {
			c.Next()
			return
		} else {
			c.JSON(403, gin.H{"error": "Forbidden, only user can access this resource"})
			c.Abort()
			return
		}
	}
}
