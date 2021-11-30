package dto

type Like struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	TargetID   string `json:"target_id"`
	UserName   string `json:"user_name"`
	UserAvatar string `json:"user_avatar"`
}
