package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/federus1105/socialmedia/internals/pkg"
	"github.com/gin-gonic/gin"
)

type ctxKey string

const UserIDKey ctxKey = "user_id"

// fungsi untuk mengetahui user yang login sekarang
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Authorization header missing or invalid",
			})
			return
		}

		// Ambil token-nya
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Buat instance Claims
		claims := &pkg.Claims{}

		// Verifikasi token
		if err := claims.VerifyToken(tokenString); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Invalid token: " + err.Error(),
			})
			return
		}

		// Simpan user_id dan role ke context
		c.Set("user_id", claims.UserId)
		c.Set("role", claims.Email)

		ctx := context.WithValue(c.Request.Context(), UserIDKey, claims.UserId)
		c.Request = c.Request.WithContext(ctx)
		// Lanjut ke handler
		c.Next()
	}
}
