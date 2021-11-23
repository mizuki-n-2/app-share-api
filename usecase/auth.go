package usecase

import (
	"app-share-api/domain/repository"

	"os"
	"time"
	"errors"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Login(email, password string) (string, error)
}

type authUsecase struct {
	userRepository repository.UserRepository
}

// authUsecaseのコンストラクタ
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
	user, err := au.userRepository.FindByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("パスワードが違います")
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