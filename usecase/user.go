package usecase

import (
	"app-share-api/domain/model"
	"app-share-api/domain/repository"
)

type UserUsecase interface {
	CreateUser(name, email, password string) (*model.User, error)
	GetUser(ID int) (*model.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (uu *userUsecase) CreateUser(name, email, password string) (*model.User, error) {
	user, err := model.NewUser(name, email, password)
	if err != nil {
		return nil, err
	}

	createdUser, err := uu.userRepository.Store(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (uu *userUsecase) GetUser(ID int) (*model.User, error) {
	user, err := uu.userRepository.FindByID(ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
