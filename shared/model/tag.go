package model

type Tag struct {
	Id   int    `json:"id" example:"1" gorm:"column:id; primaryKey"`
	Name string `json:"name" example:"tag name" gorm:"column:name"`
}
