package service

import (
	"context"
	"errors"
	"rpl-simple-backend/dto"
	"rpl-simple-backend/entity"
	"rpl-simple-backend/repository"
)

type PhoneNumberService interface {
	CreatePhoneNumber(ctx context.Context, req dto.PhoneNumberRequest) (dto.PhoneNumberResponse, error)
	GetAllPhoneNumber(ctx context.Context) ([]dto.PhoneNumberResponse, error)
	GetPhoneNumberById(ctx context.Context, id uint64) (dto.PhoneNumberResponse, error)
}

type phoneNumberService struct {
	phoneNumberRepo repository.PhoneNumberRepository
	userRepo repository.UserRepository
}

func NewPhoneNumberService(phoneNumberRepo repository.PhoneNumberRepository, userRepo repository.UserRepository) PhoneNumberService {
	return &phoneNumberService{
		phoneNumberRepo: phoneNumberRepo,
		userRepo: userRepo,
	}
}

func (s *phoneNumberService) CreatePhoneNumber(ctx context.Context, req dto.PhoneNumberRequest) (dto.PhoneNumberResponse, error) {
	user, err := s.userRepo.GetUserById(ctx, req.UserID)
	if err != nil {
		return dto.PhoneNumberResponse{}, errors.New("user not found")
	}

	phoneNumber := entity.PhoneNumber{
		PhoneNumber: req.PhoneNumber,
		CountryCode: req.CountryCode,
		UserID:      user.ID,
	}

	phoneNumber, err = s.phoneNumberRepo.Create(ctx, phoneNumber)
	if err != nil {
		return dto.PhoneNumberResponse{}, err
	}

	return dto.PhoneNumberResponse{
		ID:          phoneNumber.ID,
		PhoneNumber: phoneNumber.PhoneNumber,
		CountryCode: phoneNumber.CountryCode,
		UserID:      phoneNumber.UserID,
	}, nil
}

func (s *phoneNumberService) GetAllPhoneNumber(ctx context.Context) ([]dto.PhoneNumberResponse, error) {
	phoneNumbers, err := s.phoneNumberRepo.FindAll(ctx)
	if err != nil {
		return []dto.PhoneNumberResponse{}, err
	}

	var res []dto.PhoneNumberResponse
	for _, phoneNumber := range phoneNumbers {
		res = append(res, dto.PhoneNumberResponse{
			ID:          phoneNumber.ID,
			PhoneNumber: phoneNumber.PhoneNumber,
			CountryCode: phoneNumber.CountryCode,
			UserID:      phoneNumber.UserID,
		})
	}

	return res, nil
}

func (s *phoneNumberService) GetPhoneNumberById(ctx context.Context, id uint64) (dto.PhoneNumberResponse, error) {
	phoneNumber, err := s.phoneNumberRepo.FindByID(ctx, id)
	if err != nil {
		return dto.PhoneNumberResponse{}, err
	}

	return dto.PhoneNumberResponse{
		ID:          phoneNumber.ID,
		PhoneNumber: phoneNumber.PhoneNumber,
		CountryCode: phoneNumber.CountryCode,
		UserID:      phoneNumber.UserID,
	}, nil
}