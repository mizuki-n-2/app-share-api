package handler

import (
	"net/http"
	"strconv"
	"time"

	"app-share-api/usecase"

	"github.com/labstack/echo/v4"
)

type PostHandler interface {
	CreatePost() echo.HandlerFunc
	GetPost() echo.HandlerFunc
	GetAllPosts() echo.HandlerFunc
	UpdatePost() echo.HandlerFunc
	DeletePost() echo.HandlerFunc
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
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type responsePost struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ph *postHandler) CreatePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestPost
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		post, err := ph.postUsecase.CreatePost(req.UserID, req.Title, req.Content)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responsePost{
			ID:        post.ID,
			UserID:    post.UserID,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

func (ph *postHandler) GetPost() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		post, err := ph.postUsecase.GetPost(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responsePost{
			ID:        post.ID,
			UserID:    post.UserID,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (ph *postHandler) GetAllPosts() echo.HandlerFunc {
	return func(c echo.Context) error {
		posts, err := ph.postUsecase.GetAllPosts()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := []responsePost{}
		for _, post := range posts {
			res = append(res, responsePost{
				ID:        post.ID,
				UserID:    post.UserID,
				Title:     post.Title,
				Content:   post.Content,
				CreatedAt: post.CreatedAt,
				UpdatedAt: post.UpdatedAt,
			})
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (ph *postHandler) UpdatePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestPost
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		post, err := ph.postUsecase.UpdatePost(id, req.Title, req.Content)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responsePost{
			ID:        post.ID,
			UserID:    post.UserID,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (ph *postHandler) DeletePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = ph.postUsecase.DeletePost(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}
