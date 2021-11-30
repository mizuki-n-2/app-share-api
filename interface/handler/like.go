package handler

import (
	"net/http"
	"time"

	"app-share-api/usecase"

	"github.com/labstack/echo/v4"
)

type LikeHandler interface {
	Like() echo.HandlerFunc
	Unlike() echo.HandlerFunc
	GetLikes() echo.HandlerFunc
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
	TargetID string `json:"target_id"`
	TargetType string `json:"target_type"`
}

type responseLike struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	TargetID   string    `json:"target_id"`
	TargetType string    `json:"target_type"`
	CreatedAt  time.Time `json:"created_at"`
}

func (lh *likeHandler) Like() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestLike
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		userID := GetUserIDFromToken(c)

		like, err := lh.likeUsecase.Like(userID, req.TargetID, req.TargetType)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseLike{
			ID:        like.ID,
			UserID:    like.UserID,
			TargetID:  like.TargetID,
			TargetType: string(like.TargetType),
			CreatedAt: like.CreatedAt,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

func (lh *likeHandler) Unlike() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		userID := GetUserIDFromToken(c)

		err := lh.likeUsecase.Unlike(id, userID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}

func (lh *likeHandler) GetLikes() echo.HandlerFunc {
	return func(c echo.Context) error {
		targetID := c.Param("target_id")
		targetType := c.Param("target_type")

		likes, err := lh.likeUsecase.GetLikes(targetID, targetType)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, likes)
	}
}