package entity

type User struct {
	ID           uint64        `gorm:"primaryKey:autoIncrement" json:"id"`
	Name         string        `json:"name"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers,omitempty"`
}
