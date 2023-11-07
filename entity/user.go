package entity

type User struct {
	ID           uint64        `gorm:"primary_key:auto_increment" json:"id"`
	Name         string        `json:"name"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers,omitempty"`
}
