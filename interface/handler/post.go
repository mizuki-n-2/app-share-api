package handler

import (
	"net/http"

	"app-share-api/usecase"

	"github.com/labstack/echo"
)

type PostHandler interface {
	CreatePost() echo.HandlerFunc
}

type postHandler struct {
	postUsecase usecase.PostUsecase
}

// postHandlerのコンストラクタ
func NewPostHandler(postUsecase usecase.PostUsecase) PostHandler {
	return &postHandler{
		postUsecase: postUsecase,
	}
}

type requestPost struct {
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

type responsePost struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (ph *postHandler) CreatePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestPost
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdPost, err := ph.postUsecase.CreatePost(req.UserID, req.Title, req.Content)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responsePost{
			ID:      createdPost.ID,
			UserID:  createdPost.UserID,
			Title:   createdPost.Title,
			Content: createdPost.Content,
			CreatedAt: createdPost.CreatedAt,
			UpdatedAt: createdPost.UpdatedAt,
		}

		return c.JSON(http.StatusCreated, res)
	}
}
