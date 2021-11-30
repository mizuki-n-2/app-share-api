package queryserviceImpl

import (
	"gorm.io/gorm"

	"app-share-api/domain/queryservice"
)

type commentQueryService struct {
	db *gorm.DB
}

func NewCommentQueryService(db *gorm.DB) queryservice.CommentQueryService {
	return &commentQueryService{db: db}
}
