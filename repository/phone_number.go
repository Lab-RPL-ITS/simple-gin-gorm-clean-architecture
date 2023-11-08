package repository

import (
	"context"
	"rpl-simple-backend/entity"

	"gorm.io/gorm"
)

type PhoneNumberRepository interface {
	Create(ctx context.Context, phoneNumber entity.PhoneNumber) (entity.PhoneNumber, error)
	FindByID(ctx context.Context, id uint64) (entity.PhoneNumber, error)
	FindAll(ctx context.Context) ([]entity.PhoneNumber, error)
}

type phoneNumberRepository struct {
	db *gorm.DB
}

func NewPhoneNumberRepository(db *gorm.DB) PhoneNumberRepository {
	return &phoneNumberRepository{
		db: db,
	}	
}

func (r *phoneNumberRepository) Create(ctx context.Context, phoneNumber entity.PhoneNumber) (entity.PhoneNumber, error) {
	if err := r.db.Create(&phoneNumber).Error; err != nil {
		return entity.PhoneNumber{}, err
	}

	return phoneNumber, nil
}

func (r *phoneNumberRepository) FindByID(ctx context.Context, id uint64) (entity.PhoneNumber, error) {
	var phoneNumber entity.PhoneNumber

	if err := r.db.Where("id = ?", id).Find(&phoneNumber).Error; err != nil {
		return entity.PhoneNumber{}, err
	}

	return phoneNumber, nil
}

func (r *phoneNumberRepository) FindAll(ctx context.Context) ([]entity.PhoneNumber, error) {
	var phoneNumbers []entity.PhoneNumber

	if err := r.db.Find(&phoneNumbers).Error; err != nil {
		return []entity.PhoneNumber{}, err
	}

	return phoneNumbers, nil
}