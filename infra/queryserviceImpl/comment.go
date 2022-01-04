package queryserviceImpl

import (
	"gorm.io/gorm"

	"app-share-api/application/queryservice"
	"app-share-api/application/queryservice/dto"
)

type commentQueryService struct {
	db *gorm.DB
}

func NewCommentQueryService(db *gorm.DB) queryservice.CommentQueryService {
	return &commentQueryService{db: db}
}

func (cr *commentQueryService) GetCommentsByPostID(postID string) ([]*dto.Comment, error) {
	var comments []*dto.Comment

	query := "SELECT filtered_comments.id, filtered_comments.post_id, filtered_comments.user_id, filtered_comments.content, filtered_comments.updated_at, users.name AS user_name, users.avatar AS user_avatar, filtered_likes.likes_count FROM (SELECT * FROM comments WHERE post_id = ?) AS filtered_comments LEFT JOIN (SELECT target_id, count(*) AS likes_count FROM likes GROUP BY target_id) AS filtered_likes ON filtered_comments.id = filtered_likes.target_id LEFT JOIN users ON filtered_comments.user_id = users.id ORDER BY filtered_comments.updated_at DESC"

	if err := cr.db.Raw(query, postID).Error; err != nil {
		return nil, err
	}

	return comments, nil
}