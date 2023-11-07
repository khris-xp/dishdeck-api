package types

type MediaListResponse struct {
	Status  int         `json:"statusCode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
