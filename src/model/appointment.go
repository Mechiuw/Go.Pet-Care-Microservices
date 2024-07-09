package model

type Appointment struct {
	Id                string `json:"id" gorm:"primary_key"`
	PetId             string `json:"petid" gorm:"foreign_key"`
	ServiceProviderId string `json:"serviceproviderid" gorm:"foreign_key"`
	ServiceType       string `json:"servicetype"`
	ServiceTime       string `json:"servicetime"`
}
