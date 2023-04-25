package main

import (
	"codeflow/zero_save/sqlserver/logic"
	"codeflow/zero_save/sqlserver/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	router := gin.Default()
	// router.Use(Auth())

	go func() {
		logrus.Info("start prometheus")
		middleware.Start()
	}()

	gp := router.Group("app/v1")
	{
		gp.Use(middleware.Metric())
		gp.GET("/log", logic.Login)
	}

	router.Run("127.0.0.1:9000")
}
