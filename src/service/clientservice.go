package service

import (
	"fmt"
	"ginpet/src/db"
	"ginpet/src/model"
)

var connection = db.Pool()

func CREATE_CLIENT(client model.Client) (model.Client, error) {
	sqlStatement := `INSERT INTO client (id,name,profilenumber,address,phonenumber,email) VALUES ($1,$2,$3,$4,$5,$6)`

	_, err :=
		connection.Exec(
			sqlStatement,
			client.Id,
			client.ProfileNumber,
			client.Address,
			client.PhoneNumber,
			client.Email,
		)
	if err != nil {
		return model.Client{}, fmt.Errorf("failed to create client: %w", err)
	}

	fmt.Println("successfully added data!")
	return client, nil
}
