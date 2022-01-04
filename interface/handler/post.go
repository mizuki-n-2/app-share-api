package handler

import (
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"app-share-api/application/usecase"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"mime/multipart"
)

type PostHandler interface {
	CreatePost() echo.HandlerFunc
	GetPost() echo.HandlerFunc
	GetPosts() echo.HandlerFunc
	GetLikePosts() echo.HandlerFunc
	UpdatePost() echo.HandlerFunc
	UploadPostImage() echo.HandlerFunc
	DeletePost() echo.HandlerFunc
}

type postHandler struct {
	postUsecase usecase.PostUsecase
}

func NewPostHandler(postUsecase usecase.PostUsecase) PostHandler {
	return &postHandler{
		postUsecase: postUsecase,
	}
}

type requestPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	AppURL  string `json:"app_url"`
}

type responsePost struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Image     string    `json:"image"`
	AppURL    string    `json:"app_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ph *postHandler) CreatePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := GetUserIDFromToken(c)

		title := c.FormValue("title")
		content := c.FormValue("content")
		appURL := c.FormValue("app_url")

		// Multipart form
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}
		files := form.File["files"]

		file := files[0]

		uuid := uuid.NewString()
		image, err := uploadImage(file, uuid)
		if err != nil {
			return err
		}
		
		post, err := ph.postUsecase.CreatePost(userID, title, content, image, appURL)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responsePost{
			ID:        post.ID,
			UserID:    post.UserID,
			Title:     string(post.Title),
			Content:   string(post.Content),
			Image:     post.Image,
			AppURL:    string(post.AppURL),
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

func (ph *postHandler) GetPost() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		post, err := ph.postUsecase.GetPost(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, post)
	}
}

func (ph *postHandler) GetPosts() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.QueryParam("user_id")
		if userID != "" {
			posts, err := ph.postUsecase.GetPostsByUserID(userID)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err.Error())
			}

			return c.JSON(http.StatusOK, posts)
		} else {
			posts, err := ph.postUsecase.GetAllPosts()
			if err != nil {
				return c.JSON(http.StatusBadRequest, err.Error())
			}

			return c.JSON(http.StatusOK, posts)
		}
	}
}

func (ph *postHandler) GetLikePosts() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.QueryParam("user_id")

		posts, err := ph.postUsecase.GetLikePosts(userID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, posts)
	}
}

func (ph *postHandler) UpdatePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		userID := GetUserIDFromToken(c)

		var req requestPost
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		post, err := ph.postUsecase.UpdatePost(id, userID, req.Title, req.Content, req.AppURL)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responsePost{
			ID:        post.ID,
			UserID:    post.UserID,
			Title:     string(post.Title),
			Content:   string(post.Content),
			Image:     post.Image,
			AppURL:    string(post.AppURL),
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (ph *postHandler) UploadPostImage() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		userID := GetUserIDFromToken(c)

		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		
		image, err := uploadImage(file, id)
		if err != nil {
			return err
		}
		
		post, err := ph.postUsecase.UpdatePostImage(id, userID, image)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responsePost{
			ID:        post.ID,
			UserID:    post.UserID,
			Title:     string(post.Title),
			Content:   string(post.Content),
			Image:     post.Image,
			AppURL:    string(post.AppURL),
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (ph *postHandler) DeletePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		userID := GetUserIDFromToken(c)

		err := ph.postUsecase.DeletePost(id, userID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}

func uploadImage(file *multipart.FileHeader, id string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	fileModel := strings.Split(file.Filename, ".")
	fileName := "post_" + id + "." + fileModel[1]
	dst, err := os.Create("static/post/" + fileName)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	image := "http://localhost:8080/static/post/" + fileName

	return image, nil
}
