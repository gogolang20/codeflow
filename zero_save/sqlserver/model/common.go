package model

import "time"

type Meta struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func NewMeta() Meta {
	return Meta{
		CreatedAt: time.Now(),
	}
}

type Pagination struct {
	Limit  int64 `json:"limit,omitempty"`
	Offset int64 `json:"offset,omitempty"`
	Count  int64 `json:"count,omitempty"`
}

func NewPagination() *Pagination {
	return &Pagination{
		Limit:  15,
		Offset: 0,
		Count:  0,
	}
}
