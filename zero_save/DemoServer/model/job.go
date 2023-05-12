package model

import "github.com/gin-gonic/gin"

type Job struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Meta   Meta
}

func NewJob() *Job {
	return &Job{
		Meta: NewMeta(),
	}
}

func ToResponse(job *Job) gin.H {
	resp := gin.H{
		"ID": 123,
	}

	return resp
}
