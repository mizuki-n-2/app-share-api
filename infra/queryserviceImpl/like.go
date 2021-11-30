package queryserviceImpl

import (
	"gorm.io/gorm"

	"app-share-api/domain/queryservice"
	"app-share-api/domain/queryservice/dto"
)

type likeQueryService struct {
	db *gorm.DB
}

func NewLikeQueryService(db *gorm.DB) queryservice.LikeQueryService {
	return &likeQueryService{db: db}
}

func (lqs *likeQueryService) GetLikesByTargetID(targetID string) ([]*dto.Like, error) {
	var likes []*dto.Like

	query := "SELECT filtered_likes.id, filtered_likes.user_id, filtered_likes.target_id, users.name AS user_name, users.avatar AS user_avatar FROM (SELECT * FROM likes WHERE target_id = ?) AS filtered_likes LEFT JOIN users ON filtered_likes.user_id = users.id ORDER BY filtered_likes.created_at DESC"

	if err := lqs.db.Raw(query, targetID).Scan(&likes).Error; err != nil {
		return nil, err
	}

	return likes, nil
}
