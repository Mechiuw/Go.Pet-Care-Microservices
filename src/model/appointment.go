package model

type Appointment struct {
	Id                string `json:"id" gorm:"primary_key" validate:"required,min=1"`
	PetId             string `json:"petid" gorm:"foreign_key" validate:"required,min=1"`
	ServiceProviderId string `json:"serviceproviderid" gorm:"foreign_key" validate:"required,min=1"`
	ServiceType       string `json:"servicetype" validate:"required,min=1"`
	ServiceTime       string `json:"servicetime" validate:"required,min=1"`
}
