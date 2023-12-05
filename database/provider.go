package database

import (
	"database/sql"
	"smartjobsolutions/types"
)

func AddProvider(db *sql.DB, provider types.Provider) error {
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
	_, err := db.Exec(query, provider.Id, provider.Service, provider.Description)
	if err != nil {
		return err
	}
	return nil
}

func GetProviders(db *sql.DB) ([]types.ProviderResponse, error) {
	pingErr := db.Ping()

	if pingErr != nil {
		return []types.ProviderResponse{}, pingErr
	}

	query := `
	SELECT userdetails.username, serviceprovider.service, serviceprovider.description FROM serviceprovider INNER JOIN userdetails on userdetails.id = serviceprovider.id;
	`
	rows, err := db.Query(query)
	if err != nil {
		return []types.ProviderResponse{}, err
	}
	var providerResponseList []types.ProviderResponse
	for rows.Next() {
		var providerResponse types.ProviderResponse
		if err := rows.Scan(&providerResponse.Username, &providerResponse.Service, &providerResponse.Description); err != nil {
			return []types.ProviderResponse{}, err
		}
		providerResponseList = append(providerResponseList, providerResponse)
	}
	return providerResponseList, nil
}

func ProviderPost(db *sql.DB, ProviderPostJSON types.ProviderPostJSON) error {
	if err := db.Ping(); err != nil {
		return err
	}
	query := `
		INSERT INTO providerposts (id, post) VALUES (
			$1,
			$2
		)
	`
	_, err := db.Exec(query, ProviderPostJSON.Id, ProviderPostJSON.Post)
	if err != nil {
		return err
	}
	return nil
}