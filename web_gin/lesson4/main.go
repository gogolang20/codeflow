package main

import "github.com/gin-gonic/gin"

// 泛绑定
func main() {
	r := gin.Default()

	r.GET("/user/:name/:action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.JSON(200, gin.H{
			"name":   name,
			"action": action,
		})
	})

	r.GET("/posts/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.JSON(200, gin.H{
			"name":   name,
			"action": action,
		})
	})

	r.Run()
}
