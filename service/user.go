package service

import (
	"context"
	"rpl-simple-backend/dto"
	"rpl-simple-backend/entity"
	"rpl-simple-backend/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, req dto.UserRequest) (dto.UserResponse, error)
	GetAllUser(ctx context.Context) ([]entity.User, error)
	GetUserById(ctx context.Context, id uint64) (entity.User, error)
}

type userService struct {
	userRepo        repository.UserRepository
	phoneNumberRepo repository.PhoneNumberRepository
}

func NewUserService(userRepo repository.UserRepository, phoneNumberRepo repository.PhoneNumberRepository) UserService {
	return &userService{
		userRepo:        userRepo,
		phoneNumberRepo: phoneNumberRepo,
	}
}

func (s *userService) CreateUser(ctx context.Context, req dto.UserRequest) (dto.UserResponse, error) {
	user := entity.User{
		Name: req.Name,
	}

	user, err := s.userRepo.Create(ctx, user)
	if err != nil {
		return dto.UserResponse{}, err
	}

	var phoneNumbers []entity.PhoneNumber

	for _, data := range req.PhoneNumbers {
		phoneNumber := entity.PhoneNumber{
			PhoneNumber: data.PhoneNumber,
			CountryCode: data.CountryCode,
			UserID:      user.ID,
		}

		create, err := s.phoneNumberRepo.Create(ctx, phoneNumber)
		if err != nil {
			return dto.UserResponse{}, err
		}

		phoneNumber.ID = create.ID

		phoneNumbers = append(phoneNumbers, phoneNumber)
	}

	return dto.UserResponse{
		ID:           user.ID,
		Name:         user.Name,
		PhoneNumbers: phoneNumbers,
	}, nil
}

func (s *userService) GetAllUser(ctx context.Context) ([]entity.User, error) {
	users, err := s.userRepo.GetAllUser(ctx)
	if err != nil {
		return []entity.User{}, err
	}

	return users, nil
}

func (s *userService) GetUserById(ctx context.Context, id uint64) (entity.User, error) {
	user, err := s.userRepo.GetUserById(ctx, id)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
