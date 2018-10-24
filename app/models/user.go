package models

type User struct {
	ID        int64     `json:"id"`
	PhoneNumber   string    `json:"phone_number" validate:"required"`
	Password string `json:"-"`
	Name string `json:"name"`
}