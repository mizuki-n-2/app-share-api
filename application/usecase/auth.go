package usecase

import (
	"app-share-api/domain/repository"
	"app-share-api/domain/model"

	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthUsecase interface {
	Login(email, password string) (string, error)
}

type authUsecase struct {
	userRepository repository.UserRepository
}

func NewAuthUsecase(userRepository repository.UserRepository) AuthUsecase {
	return &authUsecase{
		userRepository: userRepository,
	}
}

type MyCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func (au *authUsecase) Login(email, password string) (string, error) {
	userEmail, err := model.NewUserEmail(email)
	if err != nil {
		return "", err
	}
	user, err := au.userRepository.FindByEmail(userEmail)
	if err != nil {
		return "", err
	}

	err = user.ComparePassword(password)
	if err != nil {
		return "", err
	}

	token, err := createToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func createToken(userID string) (string, error) {
	signingKey := []byte(os.Getenv("JWT_SIGNING_KEY"))

	claims := MyCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "app_share",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}