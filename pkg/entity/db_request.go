package entity

type DBRequest struct {
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
	MinCount  int    `json:"minCount,omitempty"`
	MaxCount  int    `json:"maxCount,omitempty"`
}
