package helper

import (
	"database/sql"
	"ginpet/src/model"
)

func ScanClient(rows *sql.Rows) []model.Client {
	// clients instance, error instance
	clients := []model.Client{}
	var err error

	// scan each property of the struct
	for rows.Next() {
		client := model.Client{}
		err = rows.Scan(
			&client.Id,
			&client.Name,
			&client.ProfileNumber,
			&client.Address,
			&client.PhoneNumber,
			&client.Email,
		)

		//handle the error
		if err != nil {
			panic("failed to scan row: " + err.Error())
		}

		clients = append(clients, client)
	}

	// handle the error and the exceptions for row iterations
	err = rows.Err()
	if err != nil {
		panic("error occurred during row iteration: " + err.Error())
	}

	//return the value
	return clients
}
