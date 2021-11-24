package infra

import (
	"time"

	"gorm.io/gorm"

	"app-share-api/domain/model/like"
	"app-share-api/domain/repository"
)

type likeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) repository.LikeRepository {
	return &likeRepository{db: db}
}

func (lr *likeRepository) Store(like *like.Like) (*like.Like, error) {
	like.CreatedAt = time.Now()
	if err := lr.db.Create(&like).Error; err != nil {
		return nil, err
	}

	return like, nil
}

func (lr *likeRepository) FindByID(ID int) (*like.Like, error) {
	like := &like.Like{ID: ID}
	if err := lr.db.First(&like).Error; err != nil {
		return nil, err
	}

	return like, nil
}

func (lr *likeRepository) Delete(ID int) error {
	like := &like.Like{ID: ID}
	if err := lr.db.Delete(&like).Error; err != nil {
		return err
	}

	return nil
}