package handler

import (
	"net/http"
	"strconv"
	"time"

	"app-share-api/usecase"

	"github.com/labstack/echo/v4"
)

type LikeHandler interface {
	Like() echo.HandlerFunc
	Unlike() echo.HandlerFunc
}

type likeHandler struct {
	likeUsecase usecase.LikeUsecase
}

func NewLikeHandler(likeUsecase usecase.LikeUsecase) LikeHandler {
	return &likeHandler{
		likeUsecase: likeUsecase,
	}
}

type requestLike struct {
	UserID int `json:"user_id"`
	PostID int `json:"post_id"`
}

type responseLike struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (lh *likeHandler) Like() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestLike
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		like, err := lh.likeUsecase.LikePost(req.UserID, req.PostID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseLike{
			ID:        like.ID,
			UserID:    like.UserID,
			PostID:    like.PostID,
			CreatedAt: like.CreatedAt,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

func (lh *likeHandler) Unlike() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = lh.likeUsecase.UnlikePost(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}
