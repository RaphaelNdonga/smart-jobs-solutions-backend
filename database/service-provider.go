package database

import (
	"database/sql"
	"smartjobsolutions/types"
)

func AddServiceProvider(db *sql.DB, serviceProvider types.ServiceProvider) error {
	pingErr := db.Ping()
	if pingErr != nil {
		return pingErr
	}
	query := `
		INSERT INTO serviceProvider VALUES (
			$1,
			$2,
			$3
		)	
	`
	_, err := db.Exec(query, serviceProvider.Id, serviceProvider.Service, serviceProvider.Description)
	if err != nil {
		return err
	}
	return nil
}

func GetServiceProviders(db *sql.DB) ([]types.ServiceProviderResponse, error) {
	pingErr := db.Ping()

	if pingErr != nil {
		return []types.ServiceProviderResponse{}, pingErr
	}

	query := `
	SELECT userdetails.username, serviceprovider.service, serviceprovider.description FROM serviceprovider INNER JOIN userdetails on userdetails.id = serviceprovider.id;
	`
	rows, err := db.Query(query)
	if err != nil {
		return []types.ServiceProviderResponse{}, err
	}
	var serviceProviderResponseList []types.ServiceProviderResponse
	for rows.Next() {
		var serviceProviderResponse types.ServiceProviderResponse
		if err := rows.Scan(&serviceProviderResponse.Username, &serviceProviderResponse.Service, &serviceProviderResponse.Description); err != nil {
			return []types.ServiceProviderResponse{}, err
		}
		serviceProviderResponseList = append(serviceProviderResponseList, serviceProviderResponse)
	}
	return serviceProviderResponseList, nil
}
