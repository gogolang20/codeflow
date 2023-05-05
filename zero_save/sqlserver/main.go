package main

import (
	"codeflow/zero_save/sqlserver/logic"
	"codeflow/zero_save/sqlserver/middleware"

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

		jobGroup.POST("/job", logic.CreateJob)
		jobGroup.GET("/job/:job_id", logic.GetJob)
		// 显示未完成任务
		jobGroup.GET("/job/list", logic.ListJob)
		jobGroup.PATCH("/job/:job_id", logic.UpdateJob)
		jobGroup.DELETE("/job/:job_id", logic.DeleteJob)
	}

	router.Run("127.0.0.1:9000")
}
