package service

import (
	"fmt"
	"ginpet/src/db"
	"ginpet/src/helper"
	"ginpet/src/model"
)

var appointment_connection = db.Pool()
var validator = helper.NewValidator()

func CREATE_APPOINTMENT(app model.Appointment) (model.Appointment, error) {
	sqlStatement := `INSERT INTO appointment (Id,petid,serviceproviderid,servicetype,servicetime) VALUES ($1,$2,$3,$4,$5)`

	err := validator.ValidateAppointment(app)

	if err != nil {
		return model.Appointment{}, fmt.Errorf("validation error: %w", err)
	}
	_, err = appointment_connection.Exec(sqlStatement, app.Id, app.PetId, app.ServiceProviderId, app.ServiceType, app.ServiceTime)
	if err != nil {
		return model.Appointment{}, fmt.Errorf("failed to create appointment: %w", err)
	}

	appResponse := model.Appointment{
		Id:                app.Id,
		PetId:             app.PetId,
		ServiceProviderId: app.ServiceProviderId,
		ServiceType:       app.ServiceType,
		ServiceTime:       app.ServiceTime,
	}

	fmt.Println("successfully create appointment")
	return appResponse, nil
}
