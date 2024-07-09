package model

type Review struct {
	Id            string `json:"id" gorm:"primary_key" validate:"required,min=1"`
	AppointmentId string `json:"appointmentid" gorm:"foreign_key" validate:"required,min=1"`
	Rating        int    `json:"rating" validate:"required,min=1"`
	ReviewText    string `json:"reviewtext" validate:"required,min=1"`
}
