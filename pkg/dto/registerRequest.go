package dto

type RegisterRequestDto struct {
	LastName    string `json:"lastname"validate:"required,min=8,max=100,alpha_space"`
	FirstName   string `json:"firstname" validate:"required"`
	Email       string
	PhoneNumber string `json:"phonenumber" validate:"required"`
	Password    string `json:"password" validate:"required"`
}
