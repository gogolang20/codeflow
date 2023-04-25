package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			c.Abort()
			return
		}
		fmt.Println("[main][Auth] ctx: ")
		// jwt.ParseWithClaims(token, jwt.Keyfunc())

		c.Set("claims", "123")
		c.Next()
	}
}
