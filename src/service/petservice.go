package service

import (
	"fmt"
	"ginpet/src/db"
	"ginpet/src/model"
)

var pet_connection = db.Pool()

func CREATE_PET(pet model.Pet) (model.Pet, error) {
	sqlStatement := `INSERT INTO pet (id,ownerid,name,breed,age,medicalhistory) VALUES ($1,$2,$3,$4,$5,$6)`

	var PetResponse model.Pet

	_, err := pet_connection.Exec(sqlStatement, pet.Id, pet.OwnerId, pet.Name, pet.Breed, pet.Age, pet.MedicalHistory)

	if err != nil {
		return model.Pet{}, fmt.Errorf("failed to create pet: %w", err)
	}

	PetResponse.Id = pet.Id
	PetResponse.OwnerId = pet.OwnerId
	PetResponse.Name = pet.Name
	PetResponse.Breed = pet.Breed
	PetResponse.Age = pet.Age
	PetResponse.MedicalHistory = pet.MedicalHistory

	fmt.Println("successfully create pet")
	return PetResponse, nil
}

func UPDATE_PET(id string, pet model.Pet) (model.Pet, error) {
	sqlStatement := `UPDATE pet SET ownerId=$1 ,name=$2, breed=$3, age=$4, medicalhistory=$5 WHERE id=$6`

	_, err := pet_connection.Exec(sqlStatement, pet.OwnerId, pet.Name, pet.Breed, pet.Age, pet.MedicalHistory, id)
	if err != nil {
		return model.Pet{}, fmt.Errorf("failed to update pet: %w", err)
	}

	PetResponse := model.Pet{
		Id:             id,
		OwnerId:        pet.OwnerId,
		Name:           pet.Name,
		Breed:          pet.Breed,
		Age:            pet.Age,
		MedicalHistory: pet.MedicalHistory,
	}
	fmt.Println("successfully updated pet")
	return PetResponse, nil
}
