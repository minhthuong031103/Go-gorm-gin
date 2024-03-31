package tagrequest

import (
	"errors"
	"go-gorm-gin/common"
	"go-gorm-gin/shared/model"
)

type CreateTagRequest struct {
	Name string `json:"name" binding:"required"`
}

func (c *CreateTagRequest) Validate() error {
	if common.IsEmptyString(c.Name) {
		return common.NewCustomError(
			errors.New("name is required"),
			"Tên không được để trống",
			"ErrorValidationNameRequired",
		)
	}
	return nil
}

func (c *CreateTagRequest) ToModel() *model.Tag {
	return &model.Tag{
		Name: c.Name,
	}
}
