package model

type Pet struct {
	Id             string `json:"id" gorm:"primary_key" validate:"required,min=1"`
	OwnerId        string `json:"ownerid" gorm:"foreign_key" validate:"required,min=1"`
	Name           string `json:"name" validate:"required,min=1"`
	Breed          string `json:"breed" validate:"required,min=1"`
	Age            string `json:"age" validate:"required,min=1"`
	MedicalHistory string `json:"medicalhistory" validate:"required,min=1"`
}
