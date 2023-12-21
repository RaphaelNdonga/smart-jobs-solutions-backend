package database

import (
	"database/sql"
	"errors"
	"smartjobsolutions/types"
)

func AddProvider(db *sql.DB, provider types.Provider) error {
	pingErr := db.Ping()
	if pingErr != nil {
		return pingErr
	}
	query := `
		INSERT INTO provider VALUES (
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

func GetProviders(db *sql.DB, service string) ([]types.ProviderResponse, error) {
	pingErr := db.Ping()

	if pingErr != nil {
		return []types.ProviderResponse{}, pingErr
	}

	query := `
	SELECT userdetails.username, provider.service, provider.description FROM provider INNER JOIN userdetails on userdetails.id = provider.id AND provider.service=$1;
	`
	rows, err := db.Query(query, service)
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

func GetProvider(db *sql.DB, providerId string) (types.Provider, error) {
	var provider types.Provider
	if err := db.Ping(); err != nil {
		return types.Provider{}, err
	}
	query := `
		SELECT * FROM provider WHERE id = $1	
	`
	err := db.QueryRow(query, providerId).Scan(&provider.Id, &provider.Service, &provider.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return types.Provider{
				Id:          "",
				Service:     "",
				Description: "",
			}, nil
		}
		return types.Provider{}, err
	}
	return provider, nil
}

func ProviderPost(db *sql.DB, ProviderPostJSON types.PostJSON) error {
	if err := db.Ping(); err != nil {
		return err
	}
	query := `
		INSERT INTO providerposts (id, post, service) VALUES (
			$1,
			$2,
			$3
		)
	`
	_, err := db.Exec(query, ProviderPostJSON.Id, ProviderPostJSON.Post, ProviderPostJSON.Service)
	if err != nil {
		return err
	}
	return nil
}

func GetProviderPosts(db *sql.DB, service string) ([]types.PostResponse, error) {
	if err := db.Ping(); err != nil {
		return []types.PostResponse{}, err
	}
	query := `
	SELECT userdetails.username, providerposts.post, providerposts.created_at, userdetails.location, providerposts.service FROM providerposts INNER JOIN userdetails ON userdetails.id = providerposts.id WHERE service = $1;
	`
	rows, err := db.Query(query, service)
	if err != nil {
		return []types.PostResponse{}, err
	}
	var providerPostResponses []types.PostResponse
	for rows.Next() {
		var providerPostResponse types.PostResponse
		rows.Scan(&providerPostResponse.Username, &providerPostResponse.Post, &providerPostResponse.CreatedAt, &providerPostResponse.Location, &providerPostResponse.Service)
		providerPostResponses = append(providerPostResponses, providerPostResponse)
	}
	return providerPostResponses, nil
}
