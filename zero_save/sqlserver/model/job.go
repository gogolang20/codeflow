package model

import "github.com/gin-gonic/gin"

type Job struct {
	JobID            string `json:"job_id"`            // 任务 ID
	Originator       string `json:"originator"`        // 发起人
	OriginatorNumber string `json:"originator_number"` // 发起人工号
	Recipient        string `json:"recipient"`         // 接收人
	RecipientNumber  string `json:"recipient_number"`  // 接收人工号
	Describe         string `json:"describe"`          // 任务描述

	Status string `json:"status"` // job 完成状态
	Meta   Meta
}

func NewJob() *Job {
	return &Job{
		Meta: NewMeta(),
	}
}

func ToResponse(job *Job) gin.H {
	resp := gin.H{
		"ID": job.JobID,
	}

	return resp
}
