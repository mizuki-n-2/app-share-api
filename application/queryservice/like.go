package queryservice

import "app-share-api/application/queryservice/dto"

type LikeQueryService interface {
	GetLikesByTargetID(targetID string) ([]*dto.Like, error)
}