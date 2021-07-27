package entity

import "time"

type DBRecord struct {
	Key        string    `json:"key,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	TotalCount int       `json:"totalCount,omitempty"`
	Counts     []int     `json:"counts,omitempty"`
}
