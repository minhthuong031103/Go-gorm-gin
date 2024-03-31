package tagservice

import (
	"context"
	tagrepository "go-gorm-gin/modules/tag/tagRepository"
	"go-gorm-gin/shared/model"
)

type ITagService interface {
	CreateTag(ctx context.Context, data *model.Tag) error
	UpdateTag(data *model.Tag) error
	DeleteTag(tag model.Tag) error
	FindTagById(id int) (model.Tag, error)
	FindAllTags() ([]model.Tag, error)
}

type TagService struct {
	TagRepo tagrepository.TagRepository
}

func NewTagService(tagRepo tagrepository.TagRepository) *TagService {
	return &TagService{
		TagRepo: tagRepo,
	}
}

func (s *TagService) CreateTag(ctx context.Context, data *model.Tag) (*model.Tag, error) {
	return s.TagRepo.Create(ctx, data)
}

func (s *TagService) UpdateTag(data *model.Tag) error {
	return s.TagRepo.Update(data)
}

func (s *TagService) DeleteTag(tag *model.Tag) error {
	return s.TagRepo.Delete(tag)
}

func (s *TagService) FindTagById(id int) (model.Tag, error) {
	return s.TagRepo.FindById(id)
}

func (s *TagService) FindAllTags() ([]model.Tag, error) {
	return s.TagRepo.FindAll()
}
