package usecase

import (
	"app-share-api/domain/model/post"
	"app-share-api/domain/repository"

	"errors"
)

type PostUsecase interface {
	CreatePost(userID int, title, content string) (*post.Post, error)
	GetPost(ID int) (*post.Post, error)
	GetAllPosts() ([]*post.Post, error)
	UpdatePost(ID, userID int, title, content string) (*post.Post, error)
	DeletePost(ID, userID int) error
}

type postUsecase struct {
	postRepository repository.PostRepository
}

func NewPostUsecase(postRepository repository.PostRepository) PostUsecase {
	return &postUsecase{
		postRepository: postRepository,
	}
}

func (pu *postUsecase) CreatePost(userID int, title, content string) (*post.Post, error) {
	post, err := post.NewPost(userID, title, content)
	if err != nil {
		return nil, err
	}

	createdPost, err := pu.postRepository.Store(post)
	if err != nil {
		return nil, err
	}

	return createdPost, nil
}

func (pu *postUsecase) GetPost(ID int) (*post.Post, error) {
	post, err := pu.postRepository.FindByID(ID)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (pu *postUsecase) GetAllPosts() ([]*post.Post, error) {
	posts, err := pu.postRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (pu *postUsecase) UpdatePost(ID, userID int, title, content string) (*post.Post, error) {
	post, err := pu.postRepository.FindByID(ID)
	if err != nil {
		return nil, err
	}

	if post.UserID != userID {
		return nil, errors.New("権限がありません")
	}

	post.Update(title, content)

	updatedPost, err := pu.postRepository.Store(post)
	if err != nil {
		return nil, err
	}

	return updatedPost, nil
}

func (pu *postUsecase) DeletePost(ID, userID int) error {
	post, err := pu.postRepository.FindByID(ID)
	if err != nil {
		return err
	}

	if post.UserID != userID {
		return errors.New("権限がありません")
	}

	err = pu.postRepository.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}