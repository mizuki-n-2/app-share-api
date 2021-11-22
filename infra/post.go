package infra

import (
	"gorm.io/gorm"

	"app-share-api/domain/model"
	"app-share-api/domain/repository"
)

type postRepository struct {
	db *gorm.DB
}

// postRepositoryのコンストラクタ
func NewPostRepository(db *gorm.DB) repository.PostRepository {
	return &postRepository{db: db}
}

func (pr *postRepository) Store(post *model.Post) (*model.Post, error) {
	if err := pr.db.Create(&post).Error; err != nil {
		return nil, err
	}

	return post, nil
}