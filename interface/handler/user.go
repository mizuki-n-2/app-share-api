package handler

import (
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"app-share-api/application/usecase"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	CreateUser() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	UploadUserAvatar() echo.HandlerFunc
	GetUser() echo.HandlerFunc
	GetRankingLikeUsers() echo.HandlerFunc
	GetRankingPostUsers() echo.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

type requestCreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type requestUpdateUser struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

type responseUser struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (uh *userHandler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestCreateUser
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
			UpdatedAt: user.UpdatedAt,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

func (uh *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		var req requestUpdateUser
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user, err := uh.userUsecase.UpdateUser(id, req.Name, req.Bio)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseUser{
			ID:        user.ID,
			Name:      string(user.Name),
			Email:     string(user.Email),
			Avatar:    user.Avatar,
			Bio:       string(user.Bio),
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uh *userHandler) UploadUserAvatar() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		fileModel := strings.Split(file.Filename, ".")
		fileName := "avatar_" + id + "." + fileModel[1]
		dst, err := os.Create("static/user/" + fileName)
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		avatar := "http://localhost:8080/static/user/" + fileName
		user, err := uh.userUsecase.UpdateUserAvatar(id, avatar)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseUser{
			ID:        user.ID,
			Name:      string(user.Name),
			Email:     string(user.Email),
			Avatar:    user.Avatar,
			Bio:       string(user.Bio),
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

func (uh *userHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		user, err := uh.userUsecase.GetUser(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, user)
	}
}

func (uh *userHandler) GetRankingLikeUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uh.userUsecase.GetRankingLikeUsers()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, users)
	}
}

func (uh *userHandler) GetRankingPostUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uh.userUsecase.GetRankingPostUsers()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, users)
	}
}