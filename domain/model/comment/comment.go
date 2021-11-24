package comment

import (
	"time"
	"errors"
)

type Comment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewComment(userID, postID int, content string) (*Comment, error) {
	if content == "" {
		return nil, errors.New("contentを入力してください")
	}

	comment := &Comment{
		UserID:  userID,
		PostID:  postID,
		Content: content,
	}

	return comment, nil
}

func (comment *Comment) Update(content string) error {
	if content == "" {
		return errors.New("contentを入力してください")
	}

	comment.Content = content

	return nil
}