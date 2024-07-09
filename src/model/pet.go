package model

type Pet struct {
	Id             string `json:"id" gorm:"primary_key"`
	OwnerId        string `json:"ownerid" gorm:"foreign_key"`
	Name           string `json:"name"`
	Breed          string `json:"breed"`
	Age            string `json:"age"`
	MedicalHistory string `json:"medicalhistory"`
}
