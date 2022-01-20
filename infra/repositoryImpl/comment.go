package repositoryImpl

import (
	"gorm.io/gorm"

	"app-share-api/domain/model/comment"
	"app-share-api/domain/repository"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) repository.CommentRepository {
	return &commentRepository{db: db}
}

func (cr *commentRepository) Store(comment *model.Comment) (*model.Comment, error) {
	if err := cr.db.Create(&comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

func (cr *commentRepository) Update(comment *model.Comment) (*model.Comment, error) {
	if err := cr.db.Save(&comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

func (cr *commentRepository) FindByID(ID string) (*model.Comment, error) {
	comment := &model.Comment{ID: ID}
	if err := cr.db.First(&comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

func (cr *commentRepository) Delete(ID string) error {
	comment := &model.Comment{ID: ID}
	if err := cr.db.Delete(&comment).Error; err != nil {
		return err
	}

	return nil
}