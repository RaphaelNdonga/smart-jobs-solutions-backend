package database_test

import (
	"log"
	"smartjobsolutions/database"
	"smartjobsolutions/types"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

func Test_AddUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userDetails := types.UserDetailsDB{
		Username:       "Raphael",
		Email:          "raphael@gmail.com",
		HashedPassword: "hashedPassword",
		Location:       "komarock",
		UserType:       "Employee",
	}
	query := `
		INSERT INTO userDetails VALUES (
			$1,
			$2,
			$3,
			$4,
			$5
		)
	`
	columns := []string{"id"}
	mock.ExpectQuery(query).
		WithArgs(
			userDetails.Username,
			userDetails.Email,
			userDetails.HashedPassword,
			userDetails.Location,
			userDetails.UserType,
		).WillReturnRows(sqlmock.NewRows(columns).AddRow(1, 1, 1, 1, 1))

	_, err = database.AddUser(db, userDetails)
	if err != nil {
		t.Errorf("unexpected error addUser: %s", err)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations were not met: %s", err)
	}
}

func Test_GetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("unexpected error getUserByEmail: %s", err)
	}
	email := "raphael@gmail.com"
	query := `
		SELECT * FROM userdetails WHERE email = $1	
	`
	columns := []string{"username", "email", "hashedpassword", "usertype"}
	mock.ExpectQuery(query).
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, 1, 1, 1))

	_, err = database.GetUserByEmail(db, email)

	if err != nil {
		t.Errorf("unexpected error getUserByEmail: %s", err)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations wer not met: %s", err)
	}
}

func Test_AddProvider(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("Test_AddProvider Error mocking db: %s", err)
	}
	provider := types.Provider{
		Id:          uuid.New().String(),
		Service:     "photography",
		Description: "I take 3d photographs",
	}
	query := `
		INSERT INTO provider VALUES (
			$1,
			$2,
			$3
		)	
	`
	columns := []string{"id", "service", "description"}
	mock.ExpectQuery(query).WithArgs(
		provider.Id,
		provider.Service,
		provider.Description,
	).WillReturnRows(sqlmock.NewRows(columns).AddRow(1, 1, 1, 1))

	err = database.AddProvider(db, provider)
	if err != nil {
		t.Errorf("Test_AddProvider Error database adding provider: %s", err)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test_AddProvider Error expectations not met: %s", err)
	}
}
