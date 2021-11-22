package usecase

import (
	"app-share-api/domain/model"
	"app-share-api/domain/repository"
)

type PostUsecase interface {
	CreatePost(userID int, title, content string) (*model.Post, error)
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