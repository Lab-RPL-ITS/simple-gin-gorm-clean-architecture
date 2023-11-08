package entity

type PhoneNumber struct {
	ID          uint64 `gorm:"primaryKey:autoIncrement" json:"id"`
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
	UserID      uint64 `json:"user_id"`
	User        *User  `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user,omitempty"`
}
