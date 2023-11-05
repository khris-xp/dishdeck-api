package responses

type MediaDto struct {
	StatusCode int     `json:"statusCode"`
	Message    string  `json:"message"`
	Data       *string `json:"data"`
}
