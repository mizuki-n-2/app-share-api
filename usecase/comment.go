package usecase

import (
	"app-share-api/domain/model"
	"app-share-api/domain/repository"

	"errors"
)

type CommentUsecase interface {
	CreateComment(userID, postID, content string) (*model.Comment, error)
	GetComments(postID string) ([]*model.Comment, error)
	UpdateComment(ID, userID, content string) (*model.Comment, error)
	DeleteComment(ID, userID string) error
}

type commentUsecase struct {
	commentRepository repository.CommentRepository
}

func NewCommentUsecase(commentRepository repository.CommentRepository) CommentUsecase {
	return &commentUsecase{
		commentRepository: commentRepository,
	}
}

func (cu *commentUsecase) CreateComment(userID, postID, content string) (*model.Comment, error) {
	comment, err := model.NewComment(userID, postID, content)
	if err != nil {
		return nil, err
	}

	createdComment, err := cu.commentRepository.Store(comment)
	if err != nil {
		return nil, err
	}

	return createdComment, nil
}

func (cu *commentUsecase) GetComments(postID string) ([]*model.Comment, error) {
	comments, err := cu.commentRepository.FindByPostID(postID)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (cu *commentUsecase) UpdateComment(ID, userID, content string) (*model.Comment, error) {
	comment, err := cu.commentRepository.FindByID(ID)
	if err != nil {
		return nil, err
	}

	if comment.UserID != userID {
		return nil, errors.New("権限がありません")
	}

	comment.Update(userID, content)

	updatedComment, err := cu.commentRepository.Update(comment)
	if err != nil {
		return nil, err
	}

	return updatedComment, nil
}

func (cu *commentUsecase) DeleteComment(ID, userID string) error {
	comment, err := cu.commentRepository.FindByID(ID)
	if err != nil {
		return err
	}

	// これはここでいいのか(?)
	if comment.UserID != userID {
		return errors.New("権限がありません")
	}

	err = cu.commentRepository.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}