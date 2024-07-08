package helper

import (
	"database/sql"
	"ginpet/src/model"
)

func ScanClient(rows *sql.Rows) []model.Client {
	clients := []model.Client{}
	var err error

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
		if err != nil {
			panic("failed to scan row: " + err.Error())
		}

		clients = append(clients, client)
	}

	err = rows.Err()
	if err != nil {
		panic("error occurred during row iteration: " + err.Error())
	}
	return clients
}
