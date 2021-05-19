package account

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"github.com/Aakanksha-jais/picshot-golang-backend/stores"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func initializeTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, stores.Account) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock, New(db, log.NewLogger())
}

func TestStores_Account_GetByUserName(t *testing.T) {
	db, mock, s := initializeTest(t)
	defer db.Close()

	mock.ExpectQuery(get + "user_name= ?").WithArgs("aakanksha_jais").WillReturnRows(getRows())

	tc := struct {
		description string
		input       *models.Account //filter
		output      *models.Account
	}{
		description: "login for a valid user, via user_name",
		input:       &models.Account{User: models.User{UserName: "aakanksha_jais", Password: "hello123"}},
		output:      getOutput(),
	}

	output, err := s.Get(context.TODO(), tc.input)

	assert.Equal(t, tc.output, output, "TEST failed.\n%s", tc.description)

	assert.Equal(t, nil, err, "TEST failed.\n%s", tc.description)
}

func TestStores_Account_GetByEmail(t *testing.T) {
	db, mock, s := initializeTest(t)
	defer db.Close()

	mock.ExpectQuery(get + "email= ?").WithArgs("jaiswal14aakanksha@gmail.com").WillReturnRows(getRows())

	tc := struct {
		description string
		input       *models.Account //filter
		output      *models.Account
	}{
		description: "login for a valid user, via email",
		input:       &models.Account{User: models.User{Email: sql.NullString{String: "jaiswal14aakanksha@gmail.com", Valid: true}, Password: "hello123"}},
		output:      getOutput(),
	}

	output, err := s.Get(context.TODO(), tc.input)

	assert.Equal(t, tc.output, output, "TEST failed.\n%s", tc.description)

	assert.Equal(t, nil, err, "TEST failed.\n%s", tc.description)
}

func TestStores_Account_GetByPhone(t *testing.T) {
	db, mock, s := initializeTest(t)
	defer db.Close()

	mock.ExpectQuery(get + "phone_no= ?").WithArgs("7807052049").WillReturnRows(getRows())

	tc := struct {
		description string
		input       *models.Account //filter
		output      *models.Account
	}{
		description: "login for a valid user, via phone number",
		input:       &models.Account{User: models.User{PhoneNo: sql.NullString{String: "7807052049", Valid: true}, Password: "hello123"}},
		output:      getOutput(),
	}

	output, err := s.Get(context.TODO(), tc.input)

	assert.Equal(t, tc.output, output, "TEST failed.\n%s", tc.description)

	assert.Equal(t, nil, err, "TEST failed.\n%s", tc.description)
}

func TestStores_Account_GetScanError(t *testing.T) {
	db, mock, s := initializeTest(t)
	defer db.Close()

	mock.ExpectQuery(get + "phone_no= ?").WithArgs("7807052049").
		WillReturnRows(
			sqlmock.NewRows([]string{`id`, `user_name`}).
				AddRow(1, "aakanksha_jais"),
		)

	tc := struct {
		description string
		input       *models.Account //filter
		err         error
	}{
		description: "get account scan error",
		input:       &models.Account{User: models.User{PhoneNo: sql.NullString{String: "7807052049", Valid: true}, Password: "hello123"}},
		err:         errors.DBError{Err: fmt.Errorf("sql: expected 2 destination arguments in Scan, not 11")},
	}

	output, err := s.Get(context.TODO(), tc.input)

	assert.Equal(t, (*models.Account)(nil), output, "TEST failed.\n%s", tc.description)

	assert.Equal(t, tc.err, err, "TEST failed.\n%s", tc.description)
}

func TestStores_Account_GetErrorNotFound(t *testing.T) {
	db, mock, s := initializeTest(t)
	defer db.Close()

	mock.ExpectQuery(get + "phone_no= ?").WithArgs("abc").WillReturnRows(sqlmock.NewRows([]string{`id`, `user_name`, `email`, `password`, `f_name`, `l_name`, `phone_no`, `created_at`, `pwd_update`, `del_req`, `status`}))

	tc := struct {
		description string
		input       *models.Account //filter
		err         error
	}{
		description: "get account, account does not exist",
		input:       &models.Account{User: models.User{PhoneNo: sql.NullString{String: "abc", Valid: true}, Password: "hello123"}},
		err:         errors.EntityNotFound{Entity: "account"},
	}

	output, err := s.Get(context.TODO(), tc.input)

	assert.Equal(t, (*models.Account)(nil), output, "TEST failed.\n%s", tc.description)

	assert.Equal(t, tc.err, err, "TEST failed.\n%s", tc.description)
}

func getOutput() *models.Account {
	lt, _ := time.Parse("2006-01-02 15:04:05", "2021-05-15 19:48:12")

	return &models.Account{
		User:       models.User{ID: 1, UserName: "aakanksha_jais", FName: "Aakanksha", LName: "Jaiswal", Email: sql.NullString{String: "jaiswal14aakanksha@gmail.com", Valid: true}, PhoneNo: sql.NullString{String: "7807052049", Valid: true}, Password: "$2a$10$.HUjOWXbMuVBXkpRLX9fuOg623yZP0/UTF4EAGHCJu1fXNWP4M7eS"},
		CreatedAt:  lt,
		PwdUpdate:  &sql.NullTime{Time: lt, Valid: true},
		DelRequest: &sql.NullTime{Valid: false},
		Status:     "ACTIVE",
	}
}

func getRows() *sqlmock.Rows {
	lt, _ := time.Parse("2006-01-02 15:04:05", "2021-05-15 19:48:12")

	return sqlmock.NewRows([]string{`id`, `user_name`, `email`, `password`, `f_name`, `l_name`, `phone_no`, `created_at`, `pwd_update`, `del_req`, `status`}).
		AddRow(1, "aakanksha_jais", "jaiswal14aakanksha@gmail.com", "$2a$10$.HUjOWXbMuVBXkpRLX9fuOg623yZP0/UTF4EAGHCJu1fXNWP4M7eS", "Aakanksha", "Jaiswal", "7807052049", lt, lt, sql.NullTime{Valid: false}, "ACTIVE")
}
