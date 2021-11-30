package queryImpl

import (
	"gorm.io/gorm"

	"app-share-api/domain/query"
)

type commentQueryService struct {
	db *gorm.DB
}

func NewCommentQueryService(db *gorm.DB) query.CommentQueryService {
	return &commentQueryService{db: db}
}
