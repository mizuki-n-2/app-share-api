package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Entity: 投稿
type Post struct {
	ID        string      `json:"id"`
	UserID    string      `json:"user_id" gorm:"not null"`
	Title     PostTitle   `json:"title" gorm:"not null;type:varchar(100)"`
	Content   PostContent `json:"content" gorm:"size:255"`
	Image     string      `json:"image"`
	AppURL    PostAppURL  `json:"app_url"`
	CreatedAt time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}

func NewPost(userID, title, content, image, appURL string) (*Post, error) {
	postID := uuid.NewString()

	postTitle, err := NewPostTitle(title)
	if err != nil {
		return nil, err
	}

	postContent, err := NewPostContent(content)
	if err != nil {
		return nil, err
	}

	postAppURL, err := NewPostAppURL(appURL)
	if err != nil {
		return nil, err
	}

	post := &Post{
		ID:      postID,
		UserID:  userID,
		Title:   postTitle,
		Content: postContent,
		Image:   image,
		AppURL:  postAppURL,
	}

	return post, nil
}

func (post *Post) Update(userID, title, content, appURL string) error {
	if userID != post.UserID {
		return errors.New("権限がありません")
	}

	postTitle, err := NewPostTitle(title)
	if err != nil {
		return err
	}

	postContent, err := NewPostContent(content)
	if err != nil {
		return err
	}

	postAppURL, err := NewPostAppURL(appURL)
	if err != nil {
		return err
	}

	post.Title = postTitle
	post.Content = postContent
	post.AppURL = postAppURL

	return nil
}

func (post *Post) UpdateImage(userID, image string) error {
	if userID != post.UserID {
		return errors.New("権限がありません")
	}

	// TODO: バリデーション
	post.Image = image

	return nil
}
