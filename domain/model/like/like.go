package model

import (
	"time"

	"github.com/google/uuid"
)

// Entity: いいね
type Like struct {
	ID         string         `json:"id"`
	UserID     string         `json:"user_id" gorm:"not null"`
	TargetID   string         `json:"target_id" gorm:"not null"`
	TargetType LikeTargetType `json:"target_type" gorm:"not null"`
	CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime"`
}

func NewLike(userID, targetID, targetType string) (*Like, error) {
	likeID := uuid.NewString()

	likeTargetType, err := NewLikeTargetType(targetType)
	if err != nil {
		return nil, err
	}

	like := &Like{
		ID:         likeID,
		UserID:     userID,
		TargetID:   targetID,
		TargetType: likeTargetType,
	}

	return like, nil
}
