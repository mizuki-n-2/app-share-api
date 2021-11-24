package handler

import (
	"net/http"

	"app-share-api/usecase"

	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	Login() echo.HandlerFunc
}

type authHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) AuthHandler {
	return &authHandler{
		authUsecase: authUsecase,
	}
}

type requestAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type responseAuth struct {
	Token string `json:"token"`
}

func (ah *authHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestAuth
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		token, err := ah.authUsecase.Login(req.Email, req.Password)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseAuth{
			Token: token,
		}

		return c.JSON(http.StatusOK, res)
	}
}