package dto

import "app-share-api/domain/model"

type Comment struct {
	model.Comment
	UserName  string `json:"user_name"`
	LikeCount int    `json:"like_count"`
}
