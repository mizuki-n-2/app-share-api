package queryservice

import "app-share-api/domain/queryservice/dto"

type LikeQueryService interface {
	GetLikesByTargetID(targetID, targetType string) ([]*dto.Like, error)
}