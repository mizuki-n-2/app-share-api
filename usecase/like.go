package usecase

import (
	"app-share-api/domain/model"
	"app-share-api/domain/repository"
)

type LikeUsecase interface {
	LikePost(userID, postID int) (*model.Like, error)
	UnlikePost(ID int) error
}

type likeUsecase struct {
	likeRepository repository.LikeRepository
}

func NewLikeUsecase(likeRepository repository.LikeRepository) LikeUsecase {
	return &likeUsecase{
		likeRepository: likeRepository,
	}
}

func (lu *likeUsecase) LikePost(userID, postID int) (*model.Like, error) {
	like, err := model.NewLike(userID, postID)
	if err != nil {
		return nil, err
	}

	createdLike, err := lu.likeRepository.Store(like)
	if err != nil {
		return nil, err
	}

	return createdLike, nil
}

func (lu *likeUsecase) UnlikePost(ID int) error {
	_, err := lu.likeRepository.FindByID(ID)
	if err != nil {
		return err
	}

	err = lu.likeRepository.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}