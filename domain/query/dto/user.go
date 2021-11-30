package dto

import (
	"app-share-api/domain/model"
)

type User struct {
	model.User
	AllLikeCount int `json:"all_like_count"`
	AllPostCount int `json:"all_post_count"`
}
