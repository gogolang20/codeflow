package model

import "github.com/gin-gonic/gin"

type Job struct {
	JobID      string `json:"job_id"`
	Originator string `json:"originator"`
	Recipient  string `json:"recipient"`
	Describe   string `json:"describe"`
	Meta       Meta
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
