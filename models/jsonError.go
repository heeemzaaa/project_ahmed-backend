package models

type ErrorJson struct {
	Status int `json:"status"`
	Error  any `json:"error,omitempty"`
}
