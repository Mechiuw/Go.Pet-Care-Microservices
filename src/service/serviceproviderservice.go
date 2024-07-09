package service

import (
	"fmt"
	"ginpet/src/db"
	"ginpet/src/model"
)

var serviceprovider_connection = db.Pool()

func CREATE_SERVICE_PROVIDER(sp model.ServiceProvider) (model.ServiceProvider, error) {
	sqlStatement := `INSERT INTO serviceprovider (id,name,servicetype,contactinfo,location) VALUES ($1,$2,$3,$4,$5)`

	_, err := serviceprovider_connection.Exec(sqlStatement, sp.Id, sp.Name, sp.ServiceType, sp.ContactInfo, sp.Location)

	if err != nil {
		return model.ServiceProvider{}, fmt.Errorf("failed to create service provider: %w", err)
	}

	SpResponse := model.ServiceProvider{
		Id:          sp.Id,
		Name:        sp.Name,
		ServiceType: sp.ServiceType,
		ContactInfo: sp.ContactInfo,
		Location:    sp.Location,
	}

	fmt.Println("successfully create service provider")
	return SpResponse, nil
}

func UPDATE_SERVICE_PROVIDER(id string, sp model.ServiceProvider) (model.ServiceProvider, error) {
	sqlStatement := `UPDATE serviceprovider SET name= $1, servicetype= $2, contactinfo= $3, location= $4 WHERE id= $5`

	_, err := serviceprovider_connection.Exec(sqlStatement, sp.Name, sp.ServiceType, sp.ContactInfo, sp.Location, sp.Id)

	if err != nil {
		return model.ServiceProvider{}, fmt.Errorf("failed to update service provider: %w", err)
	}

	SpResponse := model.ServiceProvider{
		Id:          sp.Id,
		Name:        sp.Name,
		ServiceType: sp.ServiceType,
		ContactInfo: sp.ContactInfo,
		Location:    sp.Location,
	}

	fmt.Println("successfully update service provider")
	return SpResponse, nil
}
