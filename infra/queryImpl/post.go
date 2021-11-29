package queryImpl

import (
	"app-share-api/infra"
	"app-share-api/domain/model"
	"app-share-api/domain/query"
)

type postQueryService struct {}

func NewPostQueryService() query.PostQueryService {
	return &postQueryService{}
}

func (pq *postQueryService) FindAll() ([]model.Post, error) {
	db := infra.GetDB()

	var posts []model.Post
	if err := db.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}