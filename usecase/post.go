package usecase

import (
	"app-share-api/domain/model"
	"app-share-api/domain/repository"
)

type PostUsecase interface {

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