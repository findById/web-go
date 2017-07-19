package api

type Response struct {
	StatusCode string      `json:"statusCode"`
	Message    string      `json:"message"`
	Result     interface{} `json:"result"`
}
