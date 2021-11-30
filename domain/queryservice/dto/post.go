package dto

import (
	"app-share-api/domain/model"
)

type Post struct {
	model.Post
	LikeCount    int       `json:"like_count"`
	CommentCount int       `json:"comment_count"`
}
