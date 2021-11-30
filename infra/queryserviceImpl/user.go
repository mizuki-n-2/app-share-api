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
	query := "SELECT filtered_users.id, filtered_users.name, filtered_users.avatar, filtered_users.bio, filtered_posts.posts_count AS all_posts_count FROM (SELECT * FROM users WHERE id = ?) AS filtered_users LEFT JOIN (SELECT user_id, count(*) AS posts_count FROM posts GROUP BY user_id) AS filtered_posts ON filtered_users.id = filtered_posts.user_id"
	
	if err := uqs.db.Raw(query, ID).Scan(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (uqs *userQueryService) GetAllUsers() ([]*dto.User, error) {
	var users []*dto.User
	query := "SELECT users.id, users.name, users.avatar, users.bio, filtered_posts.posts_count AS all_posts_count FROM users LEFT JOIN (SELECT user_id, count(*) AS posts_count FROM posts GROUP BY user_id) AS filtered_posts ON users.id = filtered_posts.user_id"

	if err := uqs.db.Raw(query).Scan(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}