package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt"
)

func GetUserIDFromToken(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := int(claims["user_id"].(float64))
	return userID
}