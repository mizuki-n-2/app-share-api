package repositoryImpl

import (
	"gorm.io/gorm"

	"app-share-api/domain/model"
	"app-share-api/domain/repository"
)

type likeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) repository.LikeRepository {
	return &likeRepository{db: db}
}

func (lr *likeRepository) Store(like *model.Like) (*model.Like, error) {
	if err := lr.db.Create(&like).Error; err != nil {
		return nil, err
	}

	return like, nil
}

func (lr *likeRepository) FindByID(ID string) (*model.Like, error) {
	like := &model.Like{ID: ID}
	if err := lr.db.First(&like).Error; err != nil {
		return nil, err
	}

	return like, nil
}

func (lr *likeRepository) Delete(ID string) error {
	like := &model.Like{ID: ID}
	if err := lr.db.Delete(&like).Error; err != nil {
		return err
	}

	return nil
}