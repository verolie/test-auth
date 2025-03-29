package model

type SuccessResponse struct {
	StatusCode int    `json:"StatusCode"`
	Message    string `json:"Message"`
	Data       any    `json:"data"`
}
