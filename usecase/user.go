package usecase

import (
	"app-share-api/domain/model"
	"app-share-api/domain/repository"
	"app-share-api/domain/queryservice"
	"app-share-api/domain/queryservice/dto"
)

type UserUsecase interface {
	CreateUser(name, email, password string) (*model.User, error)
	UpdateUser(ID, name, bio string) (*model.User, error)
	UpdateUserAvatar(ID, avatar string) (*model.User, error)
	GetUser(ID string) (*dto.User, error)
	GetRankingLikeUsers() ([]*dto.RankingLikeUser, error)
	GetRankingPostUsers() ([]*dto.RankingPostUser, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
	userQueryService queryservice.UserQueryService
}

func NewUserUsecase(userRepository repository.UserRepository, userQueryService queryservice.UserQueryService) UserUsecase {
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

func (uu *userUsecase) GetRankingLikeUsers() ([]*dto.RankingLikeUser, error) {
	users, err := uu.userQueryService.GetRankingLikeUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uu *userUsecase) GetRankingPostUsers() ([]*dto.RankingPostUser, error) {
	users, err := uu.userQueryService.GetRankingPostUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}