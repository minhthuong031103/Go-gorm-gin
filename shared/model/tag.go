package model

import "go-gorm-gin/common"

type Tag struct {
	Id   int    `json:"id" example:"1" gorm:"column:id; primaryKey"`
	Name string `json:"name" example:"tag name" gorm:"column:name"`
}

func (*Tag) TableName() string {
	return common.TableTag
}
