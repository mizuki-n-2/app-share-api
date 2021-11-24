package usecase

import (
	"app-share-api/domain/model/comment"
	"app-share-api/domain/repository"

	"errors"
)

type CommentUsecase interface {
	CreateComment(userID, postID int, content string) (*comment.Comment, error)
	GetComments(postID int) ([]*comment.Comment, error)
	UpdateComment(ID, userID, postID int, content string) (*comment.Comment, error)
	DeleteComment(ID, userID, postID int) error
}

type commentUsecase struct {
	commentRepository repository.CommentRepository
}

func NewCommentUsecase(commentRepository repository.CommentRepository) CommentUsecase {
	return &commentUsecase{
		commentRepository: commentRepository,
	}
}

func (cu *commentUsecase) CreateComment(userID, postID int, content string) (*comment.Comment, error) {
	comment, err := comment.NewComment(userID, postID, content)
	if err != nil {
		return nil, err
	}

	createdComment, err := cu.commentRepository.Store(comment)
	if err != nil {
		return nil, err
	}

	return createdComment, nil
}

func (cu *commentUsecase) GetComments(postID int) ([]*comment.Comment, error) {
	comments, err := cu.commentRepository.FindByPostID(postID)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (cu *commentUsecase) UpdateComment(ID, userID, postID int, content string) (*comment.Comment, error) {
	comment, err := cu.commentRepository.FindByID(ID)
	if err != nil {
		return nil, err
	}

	if comment.PostID != postID {
		return nil, errors.New("対応する投稿が正しくありません")
	}

	if comment.UserID != userID {
		return nil, errors.New("権限がありません")
	}

	comment.Update(content)

	updatedComment, err := cu.commentRepository.Store(comment)
	if err != nil {
		return nil, err
	}

	return updatedComment, nil
}

func (cu *commentUsecase) DeleteComment(ID, userID, postID int) error {
	comment, err := cu.commentRepository.FindByID(ID)
	if err != nil {
		return err
	}

	if comment.PostID != postID {
		return errors.New("対応する投稿が正しくありません")
	}

	if comment.UserID != userID {
		return errors.New("権限がありません")
	}

	err = cu.commentRepository.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}