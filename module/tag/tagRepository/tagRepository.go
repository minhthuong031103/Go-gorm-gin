package tagrepository

import (
	"go-gorm-gin/shared/model"

	"gorm.io/gorm"
)

type TagRepository interface {
	Save(tag model.Tag)
	Update(tag model.Tag)
	Delete(tag model.Tag)
	FindById(id int) (model.Tag, error)
	FindAll() []model.Tag
}

type TagRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &TagRepositoryImpl{
		Db: db,
	}
}

func (t *TagRepositoryImpl) Save(tag model.Tag) {
	result := t.Db.Create(&tag)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (t* TagRepositoryImpl) Update(tag model.Tag){
	var updatedTag=
}