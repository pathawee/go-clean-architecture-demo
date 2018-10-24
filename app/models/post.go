package models

type Post struct {
	ID        int64     `json:"id"`
	UserID        int64     `json:"user_id"`
	Message   string    `json:"phone_number" validate:"required"`
	CreatedAt string `json:"created_at"`
}