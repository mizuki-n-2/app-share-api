package model

import "time"

type Like struct {
	ID         int       `json:"id"`
	UserID     string    `json:"user_id"`
	TargetID   string    `json:"target_id"`
	TargetType string    `json:"target_type"`
	CreatedAt  time.Time `json:"created_at"`
}
