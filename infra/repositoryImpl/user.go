package repositoryImpl

import (
	"gorm.io/gorm"

	"app-share-api/domain/model"
	"app-share-api/domain/repository"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) Store(user *model.User) (*model.User, error) {
	if err := ur.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) Update(user *model.User) (*model.User, error) {
	if err := ur.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) FindByID(ID string) (*model.User, error) {
	user := &model.User{ID: ID}
	if err := ur.db.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) FindByEmail(email model.UserEmail) (*model.User, error) {
	user := &model.User{Email: email}
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
