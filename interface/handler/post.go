package handler

import (
	"app-share-api/usecase"

	"github.com/labstack/echo"
)

type PostHandler interface {
	
}

type postHandler struct {
	postUsecase usecase.PostUsecase
}

// postHandlerのコンストラクタ
func NewPostHandler(postUsecase usecase.PostUsecase) PostHandler {
	return &postHandler{
		postUsecase: postUsecase,
	}
}