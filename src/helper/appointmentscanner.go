package helper

import (
	"database/sql"
	"ginpet/src/model"
)

func ScanAppointment(rows *sql.Rows) []model.Appointment {
	apps := []model.Appointment{}
	var err error

	for rows.Next() {
		app := model.Appointment{}
		err = rows.Scan(
			&app.Id,
			&app.PetId,
			&app.ServiceProviderId,
			&app.ServiceType,
			&app.ServiceTime,
		)
		if err != nil {
			panic("failed to scan row: " + err.Error())
		}

		apps = append(apps, app)
	}

	err = rows.Err()
	if err != nil {
		panic("error occurred when while row iterations: " + err.Error())
	}

	return apps

}
