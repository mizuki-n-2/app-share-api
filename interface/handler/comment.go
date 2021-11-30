package handler

import (
	"net/http"
	"time"

	"app-share-api/usecase"

	"github.com/labstack/echo/v4"
)

type CommentHandler interface {
	CreateComment() echo.HandlerFunc
	GetComments() echo.HandlerFunc
	UpdateComment() echo.HandlerFunc
	DeleteComment() echo.HandlerFunc
}

type commentHandler struct {
	commentUsecase usecase.CommentUsecase
}

func NewCommentHandler(commentUsecase usecase.CommentUsecase) CommentHandler {
	return &commentHandler{
		commentUsecase: commentUsecase,
	}
}

type requestComment struct {
	PostID  string `json:"post_id"`
	Content string `json:"content"`
}

type responseComment struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	PostID    string    `json:"post_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ch *commentHandler) CreateComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := GetUserIDFromToken(c)

		var req requestComment
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		comment, err := ch.commentUsecase.CreateComment(userID, req.PostID, req.Content)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseComment{
			ID:        comment.ID,
			UserID:    comment.UserID,
			PostID:    comment.PostID,
			Content:   string(comment.Content),
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

func (ch *commentHandler) GetComments() echo.HandlerFunc {
	return func(c echo.Context) error {
		postID := c.QueryParam("post_id")

		comments, err := ch.commentUsecase.GetComments(postID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, comments)
	}
}

func (ch *commentHandler) UpdateComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		// TODO: バリデーション
		content := c.FormValue("content")

		userID := GetUserIDFromToken(c)

		comment, err := ch.commentUsecase.UpdateComment(id, userID, content)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseComment{
			ID:        comment.ID,
			UserID:    comment.UserID,
			PostID:    comment.PostID,
			Content:   string(comment.Content),
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (ch *commentHandler) DeleteComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		userID := GetUserIDFromToken(c)

		err := ch.commentUsecase.DeleteComment(id, userID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}
