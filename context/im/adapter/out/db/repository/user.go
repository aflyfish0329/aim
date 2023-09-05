package repository

import (
	"context"
	"test/context/im/adapter/out/db/mapper"
	"test/context/im/adapter/out/db/model"
	model1 "test/context/im/domain/model"

	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

type UserRepository struct {
	db *gorm.DB
}

func (r UserRepository) GetUserById(ctx context.Context, id string) (model1.User, error) {
	var modelUser model.User
	if err := r.db.Where("id = ?", id).Take(&modelUser).Error; err != nil {
		return model1.User{}, err
	}

	return mapper.ModelUserToDomainUser(modelUser), nil
}

func (r UserRepository) GetUserByUsername(ctx context.Context, username string) (model1.User, error) {
	var modelUser model.User
	if err := r.db.Where("username = ?", username).Take(&modelUser).Error; err != nil {
		return model1.User{}, err
	}

	return mapper.ModelUserToDomainUser(modelUser), nil
}

func (r UserRepository) CreateUser(ctx context.Context, user model1.User) error {
	modelUser := mapper.DomainUserToModelUser(user)

	return r.db.Create(&modelUser).Error
}

func (r UserRepository) UpdateUser(ctx context.Context, user model1.User) error {
	modelUser := mapper.DomainUserToModelUser(user)

	err := r.db.Model(&modelUser).Select("*").Updates(&modelUser).Error
	if err != nil {
		return err
	}

	return nil
}
