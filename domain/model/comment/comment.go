package model

import (
	"time"
	"errors"

	"github.com/google/uuid"
)

// Entity: コメント
type Comment struct {
	ID        string         `json:"id"`
	UserID    string         `json:"user_id" gorm:"not null"`
	PostID    string         `json:"post_id" gorm:"not null"`
	Content   CommentContent `json:"content" gorm:"not null;size:255"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
}

func NewComment(userID, postID, content string) (*Comment, error) {
	commentID := uuid.NewString()

	commentContent, err := NewCommentContent(content)
	if err != nil {
		return nil, err
	}

	comment := &Comment{
		ID:      commentID,
		UserID:  userID,
		PostID:  postID,
		Content: commentContent,
	}

	return comment, nil
}

func (comment *Comment) Update(userID, content string) error {
	if userID != comment.UserID {
		return errors.New("権限がありません")
	}

	commentContent, err := NewCommentContent(content)
	if err != nil {
		return err
	}

	comment.Content = commentContent

	return nil
}