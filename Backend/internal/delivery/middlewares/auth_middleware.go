package middlewares

import (
	"addis-hiwot/internal/domain/schema"
	"addis-hiwot/internal/repository"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Middleware struct {
	sr repository.SessionRepository
}

func New(sr repository.SessionRepository) *Middleware {
	return &Middleware{sr}
}
func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &schema.AuthClaim{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_AUTH_SECRET")), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		//if user logged out before the token expired
		blacklisted, _ := m.sr.IsBlacklisted(tokenStr)
		if blacklisted {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "revoked token"})
		}
		claims, ok := token.Claims.(*schema.AuthClaim)
		if !ok {
			log.Printf("[middleware] type mismatch claim is not *schema.AuthClaim  got '%T'", token.Claims)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "unknown claims type, cannot proceed"})
			return
		}
		c.Set("claim", claims)
		c.Next()
	}
}

func CheckRoles(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claim, ok := ctx.MustGet("claim").(*schema.AuthClaim)
		if !ok {
			log.Printf("[middleware] type mismatch claim is not *schema.AuthClaim  got '%T'", ctx.MustGet("claim"))
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "unknown claims type, cannot proceed"})
			return
		}
		for _, role := range roles {
			if claim.Role == role {
				ctx.Next()
				return
			}
		}
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: You don't have permission to access this resource"})
		ctx.Abort()
	}

}
