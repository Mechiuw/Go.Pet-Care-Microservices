package service

import (
	"fmt"
	"ginpet/src/db"
	"ginpet/src/model"
)

var connection = db.Pool()

func CREATE_CLIENT(client model.Client) (model.Client, error) {
	sqlStatement := `INSERT INTO client (id,name,profilenumber,address,phonenumber,email) VALUES ($1,$2,$3,$4,$5,$6);`

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

	fmt.Println("successfully added data [201]")
	return client, nil
}

func UPDATE_CLIENT(id string, updateClient map[string]string) (model.Client, error) {
	sqlStatement := `UPDATE client SET $1=$2 WHERE id = $3;`
	newPatchClient := model.Client{}
	var col string
	var val string
	for k, v := range updateClient {
		col = k
		val = v
	}

	_, err := connection.Exec(sqlStatement, col, val, id)
	if err != nil {
		return model.Client{}, fmt.Errorf("failed to update client: %w", err)
	}

	getSqlStatement := `SELECT * FROM client where id =$1`
	err = connection.QueryRow(getSqlStatement, id).Scan(
		&newPatchClient.Id, &newPatchClient.Name, &newPatchClient.ProfileNumber, &newPatchClient.Address, &newPatchClient.PhoneNumber, &newPatchClient.Email,
	)
	if err != nil {
		return newPatchClient, fmt.Errorf("failed to fetch client response: %w", err)
	}

	fmt.Println("successfully updated data [200]")
	return newPatchClient, nil
}
