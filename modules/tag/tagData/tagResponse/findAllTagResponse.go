package tagresponse

import "go-gorm-gin/shared/model"

// PaginationResponse represents the paginated response structure
type FindAllTagResponse struct {
	Data       []model.Tag `json:"data"`
	TotalPages int         `json:"totalPages"`
	Page       int         `json:"page"`
	TotalItems int         `json:"totalItems"`
}
