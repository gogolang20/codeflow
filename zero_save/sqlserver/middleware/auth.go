package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		return

		token := c.Request.Header.Get("token")
		if token == "" {
			c.Abort()
			return
		}
		fmt.Println("[middleware][Auth] ctx: ")

		c.Set("claims", "tmp")
		c.Next()
	}
}
