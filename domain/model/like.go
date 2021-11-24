package model

import "time"

type Like struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}

func NewLike(userID, postID int) (*Like, error) {
	like := &Like{
		UserID: userID,
		PostID: postID,
	}

	return like, nil
}
