package service

import (
	"fmt"
	"ginpet/src/db"
	"ginpet/src/model"
)

var serviceprovider_connection = db.Pool()

func CREATE_SERVICE_PROVIDER(sp model.ServiceProvider) (model.ServiceProvider, error) {
	sqlStatement := `INSERT INTO serviceprovider (id,name,servicetype,contactinfo,location) VALUES ($1,$2,$3,$4,$5)`

	if sp.Id == "" || sp.Name == "" || sp.ServiceType == "" || sp.ContactInfo == "" || sp.Location == "" {
		return model.ServiceProvider{}, fmt.Errorf("all fields are required")
	}

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
