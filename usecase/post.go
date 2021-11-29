package usecase

import (
	"app-share-api/domain/model"
	"app-share-api/domain/repository"

	"errors"
)

type PostUsecase interface {
	CreatePost(userID, title, content, image, appURL string) (*model.Post, error)
	GetPost(ID string) (*model.Post, error)
	GetAllPosts() ([]*model.Post, error)
	UpdatePost(ID, userID, title, content, appURL string) (*model.Post, error)
	UpdatePostImage(ID, userID, image string) (*model.Post, error)
	DeletePost(ID, userID string) error
}

type postUsecase struct {
	postRepository repository.PostRepository
}

func NewPostUsecase(postRepository repository.PostRepository) PostUsecase {
	return &postUsecase{
		postRepository: postRepository,
	}
}

func (pu *postUsecase) CreatePost(userID, title, content, image, appURL string) (*model.Post, error) {
	post, err := model.NewPost(userID, title, content, image, appURL)
	if err != nil {
		return nil, err
	}

	createdPost, err := pu.postRepository.Store(post)
	if err != nil {
		return nil, err
	}

	return createdPost, nil
}

func (pu *postUsecase) GetPost(ID string) (*model.Post, error) {
	post, err := pu.postRepository.FindByID(ID)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (pu *postUsecase) GetAllPosts() ([]*model.Post, error) {
	posts, err := pu.postRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (pu *postUsecase) UpdatePost(ID, userID, title, content, appURL string) (*model.Post, error) {
	post, err := pu.postRepository.FindByID(ID)
	if err != nil {
		return nil, err
	}

	post.Update(userID, title, content, appURL)

	updatedPost, err := pu.postRepository.Update(post)
	if err != nil {
		return nil, err
	}

	return updatedPost, nil
}

func (pu *postUsecase) UpdatePostImage(ID, userID, image string) (*model.Post, error) {
	post, err := pu.postRepository.FindByID(ID)
	if err != nil {
		return nil, err
	}

	post.UpdateImage(userID, image)

	updatedPost, err := pu.postRepository.Update(post)
	if err != nil {
		return nil, err
	}

	return updatedPost, nil
}

func (pu *postUsecase) DeletePost(ID, userID string) error {
	post, err := pu.postRepository.FindByID(ID)
	if err != nil {
		return err
	}

	// これはここでいいのか(？)
	if post.UserID != userID {
		return errors.New("権限がありません")
	}

	err = pu.postRepository.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}