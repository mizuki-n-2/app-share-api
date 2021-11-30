package usecase

import (
	"app-share-api/domain/model"
	"app-share-api/domain/repository"
	"app-share-api/domain/query"
	"app-share-api/domain/query/dto"
)

type UserUsecase interface {
	CreateUser(name, email, password string) (*model.User, error)
	UpdateUser(ID, name, bio string) (*model.User, error)
	UpdateUserAvatar(ID, avatar string) (*model.User, error)
	GetUser(ID string) (*dto.User, error)
	GetAllUsers() ([]*dto.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
	userQueryService query.UserQueryService
}

func NewUserUsecase(userRepository repository.UserRepository, userQueryService query.UserQueryService) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		userQueryService: userQueryService,
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

func (uu *userUsecase) UpdateUser(ID, name, bio string) (*model.User, error) {
	user, err := uu.userRepository.FindByID(ID)
	if err != nil {
		return nil, err
	}

	err = user.SetName(name)
	if err != nil {
		return nil, err
	}

	err = user.SetBio(bio)
	if err != nil {
		return nil, err
	}

	updatedUser, err := uu.userRepository.Update(user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (uu *userUsecase) UpdateUserAvatar(ID, avatar string) (*model.User, error) {
	user, err := uu.userRepository.FindByID(ID)
	if err != nil {
		return nil, err
	}

	err = user.SetAvatar(avatar)
	if err != nil {
		return nil, err
	}

	updatedUser, err := uu.userRepository.Update(user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (uu *userUsecase) GetUser(ID string) (*dto.User, error) {
	user, err := uu.userQueryService.GetUserByID(ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu *userUsecase) GetAllUsers() ([]*dto.User, error) {
	users, err := uu.userQueryService.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}