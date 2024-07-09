package helper

import (
	"database/sql"
	"ginpet/src/model"
)

func PetScanner(rows *sql.Rows) []model.Pet {
	pets := []model.Pet{}
	var err error

	for rows.Next() {
		pet := model.Pet{}
		err = rows.Scan(
			&pet.Id,
			&pet.OwnerId,
			&pet.Name,
			&pet.Breed,
			&pet.Age,
			&pet.MedicalHistory,
		)

		if err != nil {
			panic("failed to scan row : " + err.Error())
		}

		pets = append(pets, pet)
	}

	err = rows.Err()
	if err != nil {
		panic("error occured during row iteration: " + err.Error())
	}
	return pets
}
