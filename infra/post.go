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