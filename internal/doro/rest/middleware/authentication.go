package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"strings"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			log.Println(`header "Authorization" is missing.`)
			c.Set("actor", nil)
			c.Next()
		}
		s := strings.TrimPrefix(header, "Bearer ")
		var claims struct {
			Actor string `json:"actor"`
			jwt.RegisteredClaims
		}
		t, err := jwt.ParseWithClaims(s, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil || !t.Valid {
			c.Set("actor", nil)
			c.Next()
		}
		c.Set("actor", claims.Actor)
		c.Next()
	}
}
