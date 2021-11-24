package usecase

import (
	"app-share-api/domain/model/user"
	"app-share-api/domain/repository"
)

type UserUsecase interface {
	CreateUser(name, email, password string) (*user.User, error)
	GetUser(ID int) (*user.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (uu *userUsecase) CreateUser(name, email, password string) (*user.User, error) {
	user, err := user.NewUser(name, email, password)
	if err != nil {
		return nil, err
	}

	createdUser, err := uu.userRepository.Store(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (uu *userUsecase) GetUser(ID int) (*user.User, error) {
	user, err := uu.userRepository.FindByID(ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
