package usecase

import (
	"app-share-api/domain/model"
	"app-share-api/domain/queryservice"
	"app-share-api/domain/queryservice/dto"
	"app-share-api/domain/repository"

	"errors"
)

type PostUsecase interface {
	CreatePost(userID, title, content, image, appURL string) (*model.Post, error)
	UpdatePost(ID, userID, title, content, appURL string) (*model.Post, error)
	UpdatePostImage(ID, userID, image string) (*model.Post, error)
	DeletePost(ID, userID string) error
	GetPost(ID string) (*dto.Post, error)
	GetPostsByUserID(userID string) ([]*dto.Post, error)
	GetLikePosts(userID string) ([]*dto.Post, error)
	GetAllPosts() ([]*dto.Post, error)
}

type postUsecase struct {
	postRepository repository.PostRepository
	postQueryService queryservice.PostQueryService
}

func NewPostUsecase(postRepository repository.PostRepository, postQueryService queryservice.PostQueryService) PostUsecase {
	return &postUsecase{
		postRepository: postRepository,
		postQueryService: postQueryService,
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

func (pu *postUsecase) GetAllPosts() ([]*dto.Post, error) {
	posts, err := pu.postQueryService.GetAllPosts()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (pu *postUsecase) GetPost(ID string) (*dto.Post, error) {
	post, err := pu.postQueryService.GetPostByID(ID)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (pu *postUsecase) GetPostsByUserID(userID string) ([]*dto.Post, error) {
	posts, err := pu.postQueryService.GetPostsByUserID(userID)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (pu *postUsecase) GetLikePosts(userID string) ([]*dto.Post, error) {
	posts, err := pu.postQueryService.GetLikePostsByUserID(userID)
	if err != nil {
		return nil, err
	}

	return posts, nil
}