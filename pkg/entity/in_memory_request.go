package entity

type InMemoryRequest struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
