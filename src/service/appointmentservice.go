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

func UPDATE_APPOINTMENT(id string, app model.Appointment) (model.Appointment, error) {
	sqlStatement := `UPDATE appointment SET petid=$2, serviceproviderid=$3, servicetype=$4, servicetime=$5 WHERE Id=$1`

	err := validator.ValidateAppointment(app)

	if err != nil {
		return model.Appointment{}, fmt.Errorf("validation error: %w", err)
	}

	_, err = appointment_connection.Exec(sqlStatement, id, app.PetId, app.ServiceProviderId, app.ServiceType, app.ServiceTime)
	if err != nil {
		return model.Appointment{}, fmt.Errorf("failed to update appointment: %w", err)
	}

	appResponse := model.Appointment{
		Id:                id,
		PetId:             app.PetId,
		ServiceProviderId: app.ServiceProviderId,
		ServiceType:       app.ServiceType,
		ServiceTime:       app.ServiceTime,
	}

	fmt.Println("successfully update appointment")
	return appResponse, nil
}

func DELETE_APPOINTMENT(id string) (model.Appointment, error) {
	sqlStatement := `DELETE FROM appointment WHERE id= $1`

	_, err := appointment_connection.Exec(sqlStatement, id)
	if err != nil {
		return model.Appointment{}, fmt.Errorf("failed to delete appointment: %w", err)
	}

	fmt.Println("successfully delete appointment")
	return model.Appointment{}, nil
}

func GET_ALL_APPOINTMENT() ([]model.Appointment, error) {
	sqlStatement := `SELECT * FROM appointment`

	rows, err := appointment_connection.Query(sqlStatement)
	if err != nil {
		return []model.Appointment{}, fmt.Errorf("failed to fetch appointment: %w", err)
	}

	apps := helper.ScanAppointment(rows)
	defer rows.Close()

	fmt.Println("successfully fetch all appointment")
	return apps, nil
}

func GET_BY_ID_APPOINTMENT(id string) (model.Appointment, error) {
	sqlStatement := `SELECT * FROM appointment WHERE id=$1`

	fetchedAppointment := model.Appointment{}
	err := appointment_connection.QueryRow(sqlStatement, id).Scan(
		&fetchedAppointment.Id,
		&fetchedAppointment.PetId,
		&fetchedAppointment.ServiceProviderId,
		&fetchedAppointment.ServiceType,
		&fetchedAppointment.ServiceTime,
	)

	if err != nil {
		return model.Appointment{}, fmt.Errorf("failed to fetch appointment: %w", err)
	}

	fmt.Println("successfully fetch appointment")
	return fetchedAppointment, nil
}
