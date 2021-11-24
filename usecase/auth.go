package usecase

import (
	"app-share-api/domain/repository"
	"app-share-api/domain/model/user"

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
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func (au *authUsecase) Login(email, password string) (string, error) {
	newEmail, err := user.NewEmail(email)
	if err != nil {
		return "", err
	}
	user, err := au.userRepository.FindByEmail(*newEmail)
	if err != nil {
		return "", err
	}

	err = user.Password.Compare(password)
	if err != nil {
		return "", err
	}

	token, err := CreateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func CreateToken(userID int) (string, error) {
	signingKey := []byte(os.Getenv("JWT_SIGNING_KEY"))

	claims := MyCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
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