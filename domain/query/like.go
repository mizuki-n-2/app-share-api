package query

import "app-share-api/domain/query/dto"

type LikeQueryService interface {
	GetLikesByTargetID(targetID, targetType string) ([]*dto.Like, error)
}