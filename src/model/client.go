package model

type Client struct {
	Id            string `json:"id" gorm:"primary_key"`
	Name          string `json:"name"`
	ProfileNumber string `json:"profilenumber"`
	Address       string `json:"address"`
	PhoneNumber   string `json:"phonenumber"`
	Email         string `json:"email"`
}
