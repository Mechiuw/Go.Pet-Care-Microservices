package model

type Client struct {
	Id            string `json:"id" gorm:"primary_key" validate:"required,min=1"`
	Name          string `json:"name" validate:"required,min=1"`
	ProfileNumber string `json:"profilenumber" validate:"required,min=1"`
	Address       string `json:"address" validate:"required,min=1"`
	PhoneNumber   string `json:"phonenumber" validate:"required,min=1"`
	Email         string `json:"email" validate:"required,min=1"`
}
