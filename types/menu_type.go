package types

type MenuResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type MenuListResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type MenuUpdateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageUrl    string `json:"imageUrl" validate:"required"`
	Category    string `json:"category" validate:"required"`
}

type MenuCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageUrl    string `json:"imageUrl" validate:"required"`
	Category    string `json:"category" validate:"required"`
	CreatedBy   string `json:"createdBy" validate:"required"`
}

type MenuRatingRequest struct {
	Rate float64 `json:"rate" validate:"required"`
}

type MenuReviewRequest struct {
	Review float64 `json:"review" validate:"required"`
}
