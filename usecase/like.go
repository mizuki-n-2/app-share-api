package usecase

import (
	"app-share-api/domain/model"
	"app-share-api/domain/repository"

	"errors"
)

type LikeUsecase interface {
	Like(userID, targetID, targetType string) (*model.Like, error)
	Unlike(ID, userID string) error
}

type likeUsecase struct {
	likeRepository repository.LikeRepository
}

func NewLikeUsecase(likeRepository repository.LikeRepository) LikeUsecase {
	return &likeUsecase{
		likeRepository: likeRepository,
	}
}

func (lu *likeUsecase) Like(userID, targetID, targetType string) (*model.Like, error) {
	like, err := model.NewLike(userID, targetID, targetType)
	if err != nil {
		return nil, err
	}

	createdLike, err := lu.likeRepository.Store(like)
	if err != nil {
		return nil, err
	}

	return createdLike, nil
}

func (lu *likeUsecase) Unlike(ID, userID string) error {
	like, err := lu.likeRepository.FindByID(ID)
	if err != nil {
		return err
	}

	// このロジックはどこにおくべきか？
	if like.UserID != userID {
		return errors.New("権限がありません")
	}

	err = lu.likeRepository.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}