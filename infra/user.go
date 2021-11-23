package infra

import (
	"time"

	"gorm.io/gorm"

	"app-share-api/domain/model"
	"app-share-api/domain/repository"
)

type userRepository struct {
	db *gorm.DB
}

// userRepositoryのコンストラクタ
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) Store(user *model.User) (*model.User, error) {
	user.CreatedAt = time.Now()
	if err := ur.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) FindByID(id int) (*model.User, error) {
	user := &model.User{ID: id}
	if err := ur.db.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
