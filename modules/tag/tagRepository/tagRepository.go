package tagrepository

import (
	"context"
	"go-gorm-gin/shared/model"

	"gorm.io/gorm"
)

type TagRepository interface {
	Create(ctx context.Context, data *model.Tag) (*model.Tag, error)

	Update(data *model.Tag) error
	Delete(tag *model.Tag) error
	FindById(id int) (model.Tag, error)
	FindAll() ([]model.Tag, error)
}

type TagRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagRepositoryImpl(db *gorm.DB) TagRepository {
	return &TagRepositoryImpl{
		Db: db,
	}
}

func (t TagRepositoryImpl) Create(ctx context.Context, data *model.Tag) (*model.Tag, error) {
	t.Db = nil
	result := t.Db.Create(data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

func (t *TagRepositoryImpl) Update(tag *model.Tag) error {
	result := t.Db.Save(tag)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func (t *TagRepositoryImpl) Delete(tag *model.Tag) error {
	result := t.Db.Delete(tag)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *TagRepositoryImpl) FindById(id int) (model.Tag, error) {
	var tag model.Tag
	result := t.Db.First(tag, id)
	if result.Error != nil {
		return tag, result.Error
	}
	return tag, nil
}

func (t *TagRepositoryImpl) FindAll() ([]model.Tag, error) {
	var tags []model.Tag
	result := t.Db.Find(&tags)
	if result.Error != nil {
		return tags, result.Error
	}
	return tags, nil
}
