package handler

import (
	"net/http"
	"strconv"
	"time"

	"app-share-api/usecase"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	CreateUser() echo.HandlerFunc
	GetUser() echo.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

type requestUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type responseUser struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (uh *userHandler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestUser
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user, err := uh.userUsecase.CreateUser(req.Name, req.Email, req.Password)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseUser{
			ID:        user.ID,
			Name:      string(user.Name),
			Email:     string(user.Email),
			CreatedAt: user.CreatedAt,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

func (uh *userHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user, err := uh.userUsecase.GetUser(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseUser{
			ID:        user.ID,
			Name:      string(user.Name),
			Email:     string(user.Email),
			CreatedAt: user.CreatedAt,
		}

		return c.JSON(http.StatusOK, res)
	}
}
