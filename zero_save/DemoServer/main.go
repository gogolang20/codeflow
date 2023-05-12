package main

import (
	"codeflow/zero_save/DemoServer/logics"
	"codeflow/zero_save/DemoServer/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	router := gin.Default()

	// 授权登录
	// 用户权限查询
	router.Use(middleware.Auth())

	// monitor: prometheus
	go func() {
		logrus.Info("start prometheus")
		middleware.Start()
	}()

	jobGroup := router.Group("app/v1/")
	{
		jobGroup.Use(middleware.Metric())

		jobGroup.POST("/job", logics.CreateJob)
		jobGroup.GET("/job/:job_id", logics.GetJob)
		// 显示未完成任务
		jobGroup.GET("/job/list", logics.ListJob)
		jobGroup.PATCH("/job/:job_id", logics.UpdateJob)
		jobGroup.DELETE("/job/:job_id", logics.DeleteJob)
	}

	router.Run("127.0.0.1:9000")
}
