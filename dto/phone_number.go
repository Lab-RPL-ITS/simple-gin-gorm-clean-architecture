package dto

type PhoneNumberRequest struct {
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
	UserID      uint64 `json:"user_id"`
}

type PhoneNumberResponse struct {
	ID          uint64 `json:"id"`
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
	UserID      uint64 `json:"user_id"`
}
