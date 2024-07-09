package model

type ServiceProvider struct {
	Id          string `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	ServiceType string `json:"servicetype"`
	ContactInfo string `json:"contactinfo"`
	Location    string `json:"location"`
}
