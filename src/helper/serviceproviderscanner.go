package helper

import (
	"database/sql"
	"ginpet/src/model"
)

func ScanSP(rows *sql.Rows) []model.ServiceProvider {
	sps := []model.ServiceProvider{}
	var err error

	for rows.Next() {
		sp := model.ServiceProvider{}
		err = rows.Scan(
			&sp.Id,
			&sp.Name,
			&sp.ServiceType,
			&sp.ContactInfo,
			&sp.Location,
		)

		if err != nil {
			panic("failed to scan row : " + err.Error())
		}

		sps = append(sps, sp)
	}

	err = rows.Err()
	if err != nil {
		panic("error occured during row iteration: " + err.Error())
	}
	return sps
}
