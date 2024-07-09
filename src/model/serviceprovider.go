package model

type ServiceProvider struct {
	Id          string `json:"id" gorm:"primary_key" validate:"required,min=1"`
	Name        string `json:"name" validate:"required,min=1"`
	ServiceType string `json:"servicetype" validate:"required,min=1"`
	ContactInfo string `json:"contactinfo" validate:"required,min=1"`
	Location    string `json:"location" validate:"required,min=1"`
}
