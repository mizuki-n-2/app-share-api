package dto

type User struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Avatar        string `json:"avatar"`
	Bio           string `json:"bio"`
}

type RankingLikeUser struct {
	User
	AllLikedCount int `json:"all_liked_count"`
}

type RankingPostUser struct {
	User
	AllPostsCount int `json:"all_posts_count"`
}
