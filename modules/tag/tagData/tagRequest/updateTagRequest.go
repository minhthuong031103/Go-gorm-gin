package tagrequest

import (
	"errors"
	"go-gorm-gin/common"
)

type UpdateTagRequest struct {
	Name string `json:"name" binding:"required"`
}

func (u *UpdateTagRequest) Validate() error {
	if common.IsEmptyString(u.Name) {
		return common.NewCustomError(
			errors.New("name is required"),
			"Tên không được để trống",
			"ErrorValidationNameRequired",
		)
	}
	return nil
}
