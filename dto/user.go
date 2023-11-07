package dto

import "rpl-simple-backend/entity"

type UserRequest struct {
	Name         string `json:"name"`
	PhoneNumbers []struct {
		CountryCode string `json:"country_code"`
		PhoneNumber string `json:"phone_number"`
	} `json:"phone_numbers"`
}

type UserResponse struct {
	ID           uint64     `json:"id"`
	Name         string     `json:"name"`
	PhoneNumbers []entity.PhoneNumber `json:"phone_numbers,omitempty"`
}