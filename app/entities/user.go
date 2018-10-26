package entities

type User struct {
	ID          int64  ``
	PhoneNumber string `validate:"required"`
	Password    string `validate:"omitempty,required"`
	Name        string `validate:"required,min=6,max=20"`
}
