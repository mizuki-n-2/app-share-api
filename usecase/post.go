package usecase

import (
	"app-share-api/domain/model"
	"app-share-api/domain/repository"
)

type PostUsecase interface {
	CreatePost(userID int, title, content string) (*model.Post, error)
	GetPost(ID int) (*model.Post, error)
	GetAllPosts() ([]*model.Post, error)
	UpdatePost(ID int, title, content string) (*model.Post, error)
	DeletePost(ID int) error
}

type postUsecase struct {
	postRepository repository.PostRepository
}

// postUsecaseのコンストラクタ
func NewPostUsecase(postRepository repository.PostRepository) PostUsecase {
	return &postUsecase{
		postRepository: postRepository,
	}
}

func (pu *postUsecase) CreatePost(userID int, title, content string) (*model.Post, error) {
	post, err := model.NewPost(userID, title, content)
	if err != nil {
		return nil, err
	}

	createdPost, err := pu.postRepository.Store(post)
	if err != nil {
		return nil, err
	}

	return createdPost, nil
}

func (pu *postUsecase) GetPost(ID int) (*model.Post, error) {
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

func (pu *postUsecase) UpdatePost(ID int, title, content string) (*model.Post, error) {
	post, err := pu.postRepository.FindByID(ID)
	if err != nil {
		return nil, err
	}

	post.Update(title, content)

	updatedPost, err := pu.postRepository.Store(post)
	if err != nil {
		return nil, err
	}

	return updatedPost, nil
}

func (pu *postUsecase) DeletePost(ID int) error {
	err := pu.postRepository.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}