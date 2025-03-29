package model

type ErrorResponse struct {
	StatusCode int    `json:"StatusCode"`
	Message    string `json:"Message"`
}
