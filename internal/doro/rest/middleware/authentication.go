package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"moll-y.io/doro/internal/pkg/service"
	"strings"
)

type AuthenticationMiddleware struct {
	AuthenticationService *service.AuthenticationService
}

func (am *AuthenticationMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			log.Println(`header "Authorization" is missing.`)
			c.Set("actor", nil)
			c.Next()
			return
		}
		jwt := strings.TrimPrefix(header, "Bearer ")
		actor, err := am.AuthenticationService.Parse(jwt)
		if err != nil {
			c.Set("actor", nil)
			c.Next()
			return
		}
		c.Set("actor", actor)
		c.Next()
	}
}
