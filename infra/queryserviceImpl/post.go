package queryserviceImpl

import (
	"gorm.io/gorm"

	"app-share-api/domain/queryservice"
	"app-share-api/domain/queryservice/dto"
)

type postQueryService struct {
	db *gorm.DB
}

func NewPostQueryService(db *gorm.DB) queryservice.PostQueryService {
	return &postQueryService{db: db}
}

func (pqs *postQueryService) GetAllPosts() ([]*dto.Post, error) {
	var posts []*dto.Post
	query := "SELECT posts.id, posts.user_id, posts.title, posts.content, posts.image, posts.app_url, posts.updated_at, users.name AS user_name, users.avatar AS user_avatar, filtered_likes.likes_count, filtered_comments.comments_count FROM posts LEFT JOIN (SELECT target_id, count(*) AS likes_count FROM likes GROUP BY target_id) AS filtered_likes ON posts.id = filtered_likes.target_id LEFT JOIN (SELECT post_id, count(*) AS comments_count FROM comments GROUP BY post_id) AS filtered_comments ON posts.id = filtered_comments.post_id LEFT JOIN users ON posts.user_id = users.id ORDER BY posts.updated_at DESC"

	if err := pqs.db.Raw(query).Scan(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (pqs *postQueryService) GetPostByID(ID string) (*dto.Post, error) {
	var post *dto.Post
	query := "SELECT filtered_posts.id, filtered_posts.user_id, filtered_posts.title, filtered_posts.content, filtered_posts.image, filtered_posts.app_url, filtered_posts.updated_at, users.name AS user_name, users.avatar AS user_avatar, filtered_likes.likes_count, filtered_comments.comments_count FROM (SELECT * FROM posts WHERE id = ?) AS filtered_posts LEFT JOIN (SELECT target_id, count(*) AS likes_count FROM likes GROUP BY target_id) AS filtered_likes ON filtered_posts.id = filtered_likes.target_id LEFT JOIN (SELECT post_id, count(*) AS comments_count FROM comments GROUP BY post_id) AS filtered_comments ON filtered_posts.id = filtered_comments.post_id LEFT JOIN users ON filtered_posts.user_id = users.id ORDER BY filtered_posts.updated_at DESC"

	if err := pqs.db.Raw(query, ID).Scan(&post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (pqs *postQueryService) GetPostsByUserID(userID string) ([]*dto.Post, error) {
	var posts []*dto.Post
	query := "SELECT filtered_posts.id, filtered_posts.user_id, filtered_posts.title, filtered_posts.content, filtered_posts.image, filtered_posts.app_url, filtered_posts.updated_at, users.name AS user_name, users.avatar AS user_avatar, filtered_likes.likes_count, filtered_comments.comments_count FROM (SELECT * FROM posts WHERE user_id = ?) AS filtered_posts LEFT JOIN (SELECT target_id, count(*) AS likes_count FROM likes GROUP BY target_id) AS filtered_likes ON filtered_posts.id = filtered_likes.target_id LEFT JOIN (SELECT post_id, count(*) AS comments_count FROM comments GROUP BY post_id) AS filtered_comments ON filtered_posts.id = filtered_comments.post_id LEFT JOIN users ON filtered_posts.user_id = users.id ORDER BY filtered_posts.updated_at DESC"

	if err := pqs.db.Raw(query, userID).Scan(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (pqs *postQueryService) GetLikePostsByUserID(userID string) ([]*dto.Post, error) {
	var posts []*dto.Post
	query := "SELECT filtered_posts2.id, filtered_posts2.user_id, filtered_posts2.title, filtered_posts2.content, filtered_posts2.image, filtered_posts2.app_url, filtered_posts2.updated_at, filtered_posts2.user_name, filtered_posts2.user_avatar, filtered_posts2.likes_count, filtered_posts2.comments_count FROM (SELECT * FROM likes WHERE user_id = ? AND target_type = 'POST') AS filtered_likes2 LEFT JOIN (SELECT filtered_posts.id, filtered_posts.user_id, filtered_posts.title, filtered_posts.content, filtered_posts.image, filtered_posts.app_url, filtered_posts.updated_at, users.name AS user_name, users.avatar AS user_avatar, filtered_likes.likes_count, filtered_comments.comments_count FROM (SELECT * FROM posts WHERE user_id = ?) AS filtered_posts LEFT JOIN (SELECT target_id, count(*) AS likes_count FROM likes GROUP BY target_id) AS filtered_likes ON filtered_posts.id = filtered_likes.target_id LEFT JOIN (SELECT post_id, count(*) AS comments_count FROM comments GROUP BY post_id) AS filtered_comments ON filtered_posts.id = filtered_comments.post_id LEFT JOIN users ON filtered_posts.user_id = users.id ORDER BY filtered_posts.updated_at DESC) AS filtered_posts2 ON filtered_likes2.target_id = filtered_posts2.id ORDER BY filtered_likes2.created_at DESC"

	if err := pqs.db.Raw(query, userID).Scan(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}