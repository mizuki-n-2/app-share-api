package repositoryImpl

import (
	"time"

	"gorm.io/gorm"

	"app-share-api/domain/model/post"
	"app-share-api/domain/repository"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) repository.PostRepository {
	return &postRepository{db: db}
}

func (pr *postRepository) Store(post *post.Post) (*post.Post, error) {
	if post.ID != 0 {
		post.UpdatedAt = time.Now()
		if err := pr.db.Save(&post).Error; err != nil {
			return nil, err
		}
	} else {
		post.CreatedAt = time.Now()
		post.UpdatedAt = time.Now()
		if err := pr.db.Create(&post).Error; err != nil {
			return nil, err
		}
	}

	return post, nil
}

func (pr *postRepository) FindByID(ID int) (*post.Post, error) {
	post := &post.Post{ID: ID}
	if err := pr.db.First(&post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (pr *postRepository) FindAll() ([]*post.Post, error) {
	var posts []*post.Post
	if err := pr.db.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (pr *postRepository) Delete(ID int) error {
	post := &post.Post{ID: ID}
	if err := pr.db.Delete(&post).Error; err != nil {
		return err
	}

	return nil
}