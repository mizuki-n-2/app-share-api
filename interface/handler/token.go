package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt"
)

func GetUserIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	return userID
}