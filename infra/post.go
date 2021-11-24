package infra

import (
	"time"

	"gorm.io/gorm"

	"app-share-api/domain/model"
	"app-share-api/domain/repository"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) repository.PostRepository {
	return &postRepository{db: db}
}

func (pr *postRepository) Store(post *model.Post) (*model.Post, error) {
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

func (pr *postRepository) FindByID(ID int) (*model.Post, error) {
	post := &model.Post{ID: ID}
	if err := pr.db.First(&post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (pr *postRepository) FindAll() ([]*model.Post, error) {
	var posts []*model.Post
	if err := pr.db.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (pr *postRepository) Delete(ID int) error {
	post := &model.Post{ID: ID}
	if err := pr.db.Delete(&post).Error; err != nil {
		return err
	}

	return nil
}