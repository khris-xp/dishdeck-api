package types

type BlogResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BlogListResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BlogUpdateRequest struct {
	Title     string `json:"title" validate:"required"`
	Content   string `json:"content" validate:"required"`
	Author    string `json:"author" validate:"required"`
	UpdatedBy string `json:"updatedBy" validate:"required"`
}
