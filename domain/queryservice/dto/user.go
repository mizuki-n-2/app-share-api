package dto

type User struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Avatar        string `json:"avatar"`
	Bio           string `json:"bio"`
	// AllLikedCount int    `json:"all_liked_count"`
	AllPostCount  int    `json:"all_post_count"`
}
