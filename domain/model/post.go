package model

type Post struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	ImagePath string `json:"image_path"`
	AppURL    string `json:"app_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}