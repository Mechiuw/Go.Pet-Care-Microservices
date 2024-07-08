package service

import (
	"fmt"
	"ginpet/src/db"
	"ginpet/src/helper"
	"ginpet/src/model"
	"strings"
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

	defer connection.Close()

	fmt.Println("successfully added data [201]")
	return client, nil
}

func UPDATE_CLIENT(id string, updateClient map[string]string) (model.Client, error) {
	//check if the updateClient map null or 0
	if len(updateClient) == 0 {
		return model.Client{}, fmt.Errorf("no updates provided")
	}

	//extract updateClient k and v map
	columns := make([]string, 0, len(updateClient))
	values := make([]interface{}, 0, len(updateClient))
	for col, val := range updateClient {
		columns = append(columns, fmt.Sprintf("%s = $%d", col, len(values)+1))
		values = append(values, val)
	}
	values = append(values, id)

	//UPDATE SQL Statement here
	sqlStatement := fmt.Sprintf("UPDATE client SET %s WHERE id = '$%d';", strings.Join(columns, ", "), len(values))
	_, err := connection.Exec(sqlStatement, values...)
	if err != nil {
		return model.Client{}, fmt.Errorf("failed to update client: %w", err)
	}

	//GET SQL Statement here for response
	getSqlStatement := `SELECT * FROM client where id =$1`
	newPatchClient := model.Client{}
	err = connection.QueryRow(getSqlStatement, id).Scan(
		&newPatchClient.Id, &newPatchClient.Name, &newPatchClient.ProfileNumber, &newPatchClient.Address, &newPatchClient.PhoneNumber, &newPatchClient.Email,
	)
	if err != nil {
		return newPatchClient, fmt.Errorf("failed to fetch client response: %w", err)
	}

	//Response
	fmt.Println("successfully updated data [200]")
	return newPatchClient, nil
}

func DELETE_CLIENT(id string) (model.Client, error) {
	sqlStatement := `DELETE FROM client where id=$1`

	_, err := connection.Exec(sqlStatement, id)
	if err != nil {
		return model.Client{}, fmt.Errorf("failed to delete client: %w", err)
	}

	fmt.Println("successfully deleted data")
	return model.Client{}, nil

}

func GET_CLIENT(id string) ([]model.Client, error) {
	sqlStatement := `SELECT * FROM client`

	rows, err := connection.Query(sqlStatement)
	if err != nil {
		return []model.Client{}, fmt.Errorf("failed to fetch client: %w", err)
	}

	defer rows.Close()

	clients := helper.ScanClient(rows)

	fmt.Println("successfully fetch client data ")
	return clients, nil
}
