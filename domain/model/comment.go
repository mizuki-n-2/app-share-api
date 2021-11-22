package model

type Comment struct {
	ID        int   `json:"id"`
	UserID    string `json:"user_id"`
	PostID    string `json:"post_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
