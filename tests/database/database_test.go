package database_test

import (
	"log"
	"smartjobsolutions/database"
	"smartjobsolutions/types"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
		UserType:       "Employee",
	}
	query := `
		INSERT INTO userDetails VALUES (
			$1,
			$2,
			$3,
			$4	
		)
	`
	mock.ExpectExec(query).
		WithArgs(
			userDetails.Username,
			userDetails.Email,
			userDetails.HashedPassword,
			userDetails.UserType,
		).WillReturnResult(sqlmock.NewResult(1, 1))

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
