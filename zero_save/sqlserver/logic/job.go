package logic

import (
	"net/http"

	"codeflow/zero_save/sqlserver/dao"
	"codeflow/zero_save/sqlserver/model"

	"github.com/gin-gonic/gin"
)

func CreateJob(c *gin.Context) {
	var job model.Job
	if err := c.ShouldBind(&job); err != nil {
		c.String(http.StatusBadRequest, "error")
	}

	newJob, err := dao.CreateJob(job)
	if err != nil {
		c.String(http.StatusBadRequest, "error")
	}

	c.JSON(http.StatusOK, model.ToResponse(newJob))
}

func GetJob(c *gin.Context) {
	jobID := c.Param("job_id")

	job, err := dao.GetJob(jobID)
	if err != nil {
		c.String(http.StatusBadRequest, "error")
	}

	c.JSON(http.StatusOK, model.ToResponse(job))
}

func ListJob(c *gin.Context) {

	jobs, err := dao.ListJob()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, jobs)
}

func UpdateJob(c *gin.Context) {

	job, err := dao.UpdateJob()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, model.ToResponse(job))
}

func DeleteJob(c *gin.Context) {

	if err := dao.DeleteJob(); err != nil {
		c.String(http.StatusBadRequest, "ok")
	}

	c.String(http.StatusOK, "ok")
}
