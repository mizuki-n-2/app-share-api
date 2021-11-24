package infra

import (
	"time"

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

func (cr *commentRepository) Store(comment *comment.Comment) (*comment.Comment, error) {
	if comment.ID != 0 {
		comment.UpdatedAt = time.Now()
		if err := cr.db.Save(&comment).Error; err != nil {
			return nil, err
		}
	} else {
		comment.CreatedAt = time.Now()
		comment.UpdatedAt = time.Now()
		if err := cr.db.Create(&comment).Error; err != nil {
			return nil, err
		}
	}

	return comment, nil
}

func (cr *commentRepository) FindByID(ID int) (*comment.Comment, error) {
	comment := &comment.Comment{ID: ID}
	if err := cr.db.First(&comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

func (cr *commentRepository) FindByPostID(postID int) ([]*comment.Comment, error) {
	var comments []*comment.Comment
	if err := cr.db.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (cr *commentRepository) Delete(ID int) error {
	comment := &comment.Comment{ID: ID}
	if err := cr.db.Delete(&comment).Error; err != nil {
		return err
	}

	return nil
}