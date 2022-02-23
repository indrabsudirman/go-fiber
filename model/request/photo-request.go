package request

type PhotoCreateRequest struct {
	CategoryId uint `json:"category_id" validate:"required"`
}
