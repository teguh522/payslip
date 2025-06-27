package middleware

import (
	"github.com/gin-gonic/gin"
)

func (m *AuthMiddleware) RoleMiddleware() gin.HandlerFunc {
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
