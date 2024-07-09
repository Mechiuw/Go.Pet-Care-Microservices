package model

type Review struct {
	Id            string `json:"id" gorm:"primary_key"`
	AppointmentId string `json:"appointmentid" gorm:"foreign_key"`
	Rating        int    `json:"rating"`
	ReviewText    string `json:"reviewtext"`
}
