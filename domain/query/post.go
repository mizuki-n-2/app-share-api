package query

import (
	"app-share-api/domain/model"
)

type PostQueryService interface {
	FindAll() ([]model.Post, error)
}