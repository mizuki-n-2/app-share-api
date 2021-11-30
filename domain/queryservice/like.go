package queryservice

import "app-share-api/domain/queryservice/dto"

type LikeQueryService interface {
	GetLikesByTargetID(targetID string) ([]*dto.Like, error)
}