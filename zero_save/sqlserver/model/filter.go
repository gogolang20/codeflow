package model

type JobFilter struct {
	JobID string `json:"job_id"`
}

func (j *JobFilter) MysqlFilter() {

}
