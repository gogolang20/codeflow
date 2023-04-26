package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// query postform
func main() {
	r := gin.Default()

	r.POST("/user/:id", PostForm)

	r.Run()
}

func PostForm(c *gin.Context) {
	id := c.Param("id")
	userName := c.PostForm("user_name")

	c.JSON(http.StatusOK, gin.H{
		"id":        id,
		"user_name": userName,
	})
}
