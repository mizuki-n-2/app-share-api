package dto

import "app-share-api/domain/model"

type Like struct {
	model.Like
	UserName string `json:"user_name"`
}
