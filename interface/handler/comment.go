package handler

import (
	"net/http"
	"strconv"
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
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}

type responseComment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ch *commentHandler) CreateComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		postID, err := strconv.Atoi(c.Param("post_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestComment
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		comment, err := ch.commentUsecase.CreateComment(req.UserID, postID, req.Content)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseComment{
			ID:        comment.ID,
			UserID:    comment.UserID,
			PostID: 	comment.PostID,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

func (ch *commentHandler) GetComments() echo.HandlerFunc {
	return func(c echo.Context) error {
		postID, err := strconv.Atoi(c.Param("post_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		comments, err := ch.commentUsecase.GetComments(postID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := []responseComment{}
		for _, comment := range comments {
			res = append(res, responseComment{
				ID:        comment.ID,
				UserID:    comment.UserID,
				PostID:    comment.PostID,
				Content:   comment.Content,
				CreatedAt: comment.CreatedAt,
				UpdatedAt: comment.UpdatedAt,
			})
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (ch *commentHandler) UpdateComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestComment
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		comment, err := ch.commentUsecase.UpdateComment(id, req.Content)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseComment{
			ID:        comment.ID,
			UserID:    comment.UserID,
			PostID:    comment.PostID,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (ch *commentHandler) DeleteComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = ch.commentUsecase.DeleteComment(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}
