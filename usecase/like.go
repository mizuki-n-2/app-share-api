package usecase

import (
	"app-share-api/domain/model"
	"app-share-api/domain/repository"

	"errors"
)

type LikeUsecase interface {
	LikePost(userID, postID int) (*model.Like, error)
	UnlikePost(ID, userID, postID int) error
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

func (lu *likeUsecase) UnlikePost(ID, userID, postID int) error {
	like, err := lu.likeRepository.FindByID(ID)
	if err != nil {
		return err
	}
	
	if like.PostID != postID {
		return errors.New("対応する投稿が正しくありません")
	}

	if like.UserID != userID {
		return errors.New("権限がありません")
	}

	err = lu.likeRepository.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}