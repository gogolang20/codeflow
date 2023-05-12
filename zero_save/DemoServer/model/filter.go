package model

type JobFilter struct {
	ID string `json:"id"`
}

func (j *JobFilter) MysqlFilter() {

}
