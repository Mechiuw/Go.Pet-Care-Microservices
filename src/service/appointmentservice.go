package service

import (
	"fmt"
	"ginpet/src/db"
	"ginpet/src/helper"
	"ginpet/src/model"
)

var appointment_connection = db.Pool()
var validator = helper.NewValidator()

func CREATE_APPOINTMENT(app model.Appointment) (model.Appointment,error){
	sqlStatement := `INSERT INTO appointment (Id,petid,serviceproviderid,servicetype,servicetime) VALUES ($1,$2,$3,$4,$5)`

	validated := 

	_,err := appointment_connection.Exec(sqlStatement,)

	appResponse := model.Appointment{
		Id: app.Id,
		PetId: app.PetId,
		ServiceProviderId: app.ServiceProviderId,
		ServiceType: app.ServiceType,
		ServiceTime: app.ServiceTime,
	}

	fmt.Println("successfully create appointment")
	return appResponse,nil
}