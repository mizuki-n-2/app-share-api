package repositoryImpl

import (
	"gorm.io/gorm"

	"app-share-api/domain/model/user"
	"app-share-api/domain/repository"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) Store(user *user.User) (*user.User, error) {
	if err := ur.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) FindByID(ID int) (*user.User, error) {
	user := &user.User{ID: ID}
	if err := ur.db.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) FindByEmail(email user.Email) (*user.User, error) {
	user := &user.User{Email: email}
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}