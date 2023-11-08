package repository

import (
	"context"
	"rpl-simple-backend/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user entity.User) (entity.User, error)
	GetAllUser(ctx context.Context) ([]entity.User, error)
	GetUserById(ctx context.Context, id uint64) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}	
}

func (r *userRepository) Create(ctx context.Context, user entity.User) (entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetAllUser(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	if err := r.db.Find(&users).Error; err != nil {
		return []entity.User{}, err
	}

	return users, nil
}

func (r *userRepository) GetUserById(ctx context.Context, id uint64) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}