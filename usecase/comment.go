package usecase

import (
	"app-share-api/domain/model"
	"app-share-api/domain/repository"
)

type CommentUsecase interface {
	CreateComment(userID, postID int, content string) (*model.Comment, error)
	GetComments(postID int) ([]*model.Comment, error)
	UpdateComment(ID int, content string) (*model.Comment, error)
	DeleteComment(ID int) error
}

type commentUsecase struct {
	commentRepository repository.CommentRepository
}

func NewCommentUsecase(commentRepository repository.CommentRepository) CommentUsecase {
	return &commentUsecase{
		commentRepository: commentRepository,
	}
}

func (cu *commentUsecase) CreateComment(userID, postID int, content string) (*model.Comment, error) {
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

func (cu *commentUsecase) GetComments(postID int) ([]*model.Comment, error) {
	comments, err := cu.commentRepository.FindByPostID(postID)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (cu *commentUsecase) UpdateComment(ID int, content string) (*model.Comment, error) {
	comment, err := cu.commentRepository.FindByID(ID)
	if err != nil {
		return nil, err
	}

	comment.Update(content)

	updatedComment, err := cu.commentRepository.Store(comment)
	if err != nil {
		return nil, err
	}

	return updatedComment, nil
}

func (cu *commentUsecase) DeleteComment(ID int) error {
	_, err := cu.commentRepository.FindByID(ID)
	if err != nil {
		return err
	}

	err = cu.commentRepository.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}