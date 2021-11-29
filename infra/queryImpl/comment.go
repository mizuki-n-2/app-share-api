package queryImpl

import (
	"app-share-api/infra"
	"app-share-api/domain/model"
	"app-share-api/domain/query"
)

type commentQueryService struct {}

func NewCommentQueryService() query.CommentQueryService {
	return &commentQueryService{}
}

func (cr *commentQueryService) FindByPostID(postID string) ([]model.Comment, error) {
	db := infra.GetDB()

	var comments []model.Comment
	if err := db.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}