package account

import (
	"context"
	"database/sql"
	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"github.com/Aakanksha-jais/picshot-golang-backend/services"
	"github.com/Aakanksha-jais/picshot-golang-backend/stores"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestServices_Account_Get(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAccountStore := stores.NewMockAccount(mockCtrl)
	mockBlogService := services.NewMockBlog(mockCtrl)

	s := New(mockAccountStore, mockBlogService, log.NewLogger())

	account := &models.Account{
		User:   models.User{ID: 1, UserName: "aakanksha_jais", FName: "Aakanksha", LName: "Jaiswal", Email: sql.NullString{String: "jaiswal14aakanksha@gmail.com", Valid: true}, PhoneNo: sql.NullString{String: "7807052049", Valid: true}, Password: "$2a$10$.HUjOWXbMuVBXkpRLX9fuOg623yZP0/UTF4EAGHCJu1fXNWP4M7eS"},
		Status: "ACTIVE",
	}

	mockAccountStore.EXPECT().Get(gomock.Any(), &models.Account{User: models.User{UserName: "aakanksha_jais", Password: "hello123"}}).
		Return(account, nil)

	mockAccountStore.EXPECT().Get(gomock.Any(), &models.Account{User: models.User{Email: sql.NullString{String: "jaiswal14aakanksha@gmail.com", Valid: true}, Password: "hello123"}}).
		Return(account, nil)

	mockAccountStore.EXPECT().Get(gomock.Any(), &models.Account{User: models.User{UserName: "random_user", Password: "hello123"}}).
		Return(nil, nil)

	mockAccountStore.EXPECT().Get(gomock.Any(), &models.Account{User: models.User{UserName: "aakanksha_jais", Password: "wrong123"}}).
		Return(account, nil)

	tests := []struct {
		description string
		input       *models.User
		output      *models.Account
		err         error
	}{
		{
			description: "login for a valid user, via user_name",
			input:       &models.User{UserName: "aakanksha_jais", Password: "hello123"},
			output:      account,
		},
		{
			description: "login for a valid user, via email",
			input:       &models.User{Email: sql.NullString{String: "jaiswal14aakanksha@gmail.com", Valid: true}, Password: "hello123"},
			output:      account,
		},
		{
			description: "login for a invalid user, user does not exist",
			input:       &models.User{UserName: "random_user", Password: "hello123"},
			err:         errors.EntityNotFound{Entity: "user"},
		},
		{
			description: "login for a valid user, via user_name, but incorrect password",
			input:       &models.User{UserName: "aakanksha_jais", Password: "wrong123"},
			err:         errors.AuthError{Err: bcrypt.ErrMismatchedHashAndPassword, Message: "invalid password"},
		},
		{
			description: "login for a invalid (empty) user details",
			input:       nil,
			err:         errors.MissingParam{Param: "user details"},
		},
	}

	for i := range tests {
		output, err := s.Get(context.TODO(), tests[i].input)

		assert.Equal(t, tests[i].output, output, "TEST [%v], failed.\n%s", i+1, tests[i].description)

		assert.Equal(t, tests[i].err, err, "TEST [%v], failed.\n%s", i+1, tests[i].description)
	}
}
