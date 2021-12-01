package queryserviceImpl

import (
	"gorm.io/gorm"

	"app-share-api/domain/queryservice"
	"app-share-api/domain/queryservice/dto"
)

type userQueryService struct {
	db *gorm.DB
}

func NewUserQueryService(db *gorm.DB) queryservice.UserQueryService {
	return &userQueryService{db: db}
}

func (uqs *userQueryService) GetUserByID(ID string) (*dto.User, error) {
	var user *dto.User
	query := "SELECT id, name, avatar, bio FROM users WHERE id = ?"
	
	if err := uqs.db.Raw(query, ID).Scan(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (uqs *userQueryService) GetRankingLikeUsers() ([]*dto.RankingLikeUser, error) {
	var users []*dto.RankingLikeUser
	query := "SELECT users.id, users.name, users.avatar, users.bio, filtered_posts.all_liked_count FROM users LEFT JOIN (SELECT posts.user_id, SUM(filtered_likes.likes_count) AS all_liked_count FROM posts LEFT JOIN (SELECT target_id, count(*) AS likes_count FROM likes GROUP BY target_id) AS filtered_likes ON posts.id = filtered_likes.target_id) AS filtered_posts ON users.id = filtered_posts.user_id"

	if err := uqs.db.Raw(query).Scan(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (uqs *userQueryService) GetRankingPostUsers() ([]*dto.RankingPostUser, error) {
	var users []*dto.RankingPostUser
	query := "SELECT users.id, users.name, users.avatar, users.bio, filtered_posts.posts_count AS all_posts_count FROM users LEFT JOIN (SELECT user_id, count(*) AS posts_count FROM posts GROUP BY user_id) AS filtered_posts ON users.id = filtered_posts.user_id"

	if err := uqs.db.Raw(query).Scan(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}