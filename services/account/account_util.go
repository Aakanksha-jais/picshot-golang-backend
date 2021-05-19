package account

import (
	"context"
	"database/sql"
	"regexp"
	"strings"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
)

func (a account) checkUsernameAvailability(ctx context.Context, username string) error {
	if err := validateUsername(username); err != nil {
		return err
	}

	acc, _ := a.accountStore.Get(ctx, &models.Account{User: models.User{UserName: username}})
	if acc != nil {
		return errors.EntityAlreadyExists{Entity: "user", ValueType: "username", Value: username}
	}

	a.logger.Debugf("username %s available", username)
	return nil
}

func (a account) checkEmailAvailability(ctx context.Context, email string) error {
	if err := validateEmail(email); err != nil {
		return err
	}

	acc, _ := a.accountStore.Get(ctx, &models.Account{User: models.User{Email: sql.NullString{String: email, Valid: true}}})
	if acc != nil {
		return errors.EntityAlreadyExists{Entity: "user", ValueType: "email", Value: email}
	}

	a.logger.Debugf("email %s available", email)
	return nil
}

func (a account) checkPhoneAvailability(ctx context.Context, phone string) error {
	if err := validatePhone(phone); err != nil {
		return err
	}

	acc, _ := a.accountStore.Get(ctx, &models.Account{User: models.User{PhoneNo: sql.NullString{String: phone, Valid: true}}})
	if acc != nil {
		return errors.EntityAlreadyExists{Entity: "user", ValueType: "phone_no", Value: phone}
	}

	a.logger.Debugf("phone number %s available", phone)
	return nil
}

func (a account) checkUserExists(ctx context.Context, user *models.User) error {
	acc, _ := a.accountStore.Get(ctx, &models.Account{User: models.User{UserName: user.UserName}})
	if acc != nil {
		return errors.EntityAlreadyExists{Entity: "user", ValueType: "username", Value: user.UserName}
	}

	acc, _ = a.accountStore.Get(ctx, &models.Account{User: models.User{Email: user.Email}})
	if acc != nil {
		return errors.EntityAlreadyExists{Entity: "user", ValueType: "email", Value: user.Email.String}
	}

	acc, _ = a.accountStore.Get(ctx, &models.Account{User: models.User{PhoneNo: user.PhoneNo}})
	if acc != nil {
		return errors.EntityAlreadyExists{Entity: "user", ValueType: "phone number", Value: user.PhoneNo.String}
	}

	return nil
}

func validateDetails(user *models.User) error {
	if user.UserName == "" {
		return errors.MissingParam{Param: "user_name"}
	}

	if err := validateUsername(user.UserName); err != nil {
		return err
	}

	if user.Email.String == "" && user.PhoneNo.String == "" {
		return errors.MissingParam{Param: "email"}
	}

	if user.Email.String != "" {
		if err := validateEmail(user.Email.String); err != nil {
			return err
		}
	}

	if user.PhoneNo.String != "" {
		if err := validatePhone(user.PhoneNo.String); err != nil {
			return err
		}
	}

	if err := validatePassword(user.Password); err != nil {
		return err
	}

	if user.FName == "" && user.LName == "" {
		return errors.MissingParam{Param: "name"}
	}

	if err := validateName(user.FName); err != nil {
		return err
	}

	err := validateName(user.LName)

	return err
}

func validateName(name string) error {
	// username should be aplha-numeric and should have at least 8 characters
	res, err := regexp.MatchString("^[a-zA-Z]+$", name)
	if err == nil && res {
		return nil
	}

	return errors.InvalidParam{Param: "name"}
}

func validateUsername(username string) error {
	// username should be aplha-numeric and should have at least 6 characters
	res, err := regexp.MatchString(`^[0-9A-Za-z_]{6,}$`, username)
	if err == nil && res {
		return nil
	}

	return errors.InvalidParam{Param: "username"}
}

func validateEmail(email string) error {
	res, err := regexp.MatchString(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+[.][a-zA-Z0-9-.]+$`, email)
	if err == nil && res {
		return nil
	}

	return errors.InvalidParam{Param: "email"}
}

func validatePhone(phone string) error {
	res, err := regexp.MatchString(`^[0-9]+$`, phone)
	if err == nil && res {
		return nil
	}

	return errors.InvalidParam{Param: "phone_no"}
}

func validatePassword(password string) error {
	// password between 8 to 20 characters
	// alphanumeric and !@#$%^&* symbols allowed
	res, err := regexp.MatchString(`^[a-zA-Z0-9!@#$%^&*]{8,20}$`, password)
	if err == nil && res {
		return nil
	}

	return errors.InvalidParam{Param: "password"}
}

func empty(user *models.User) bool {
	return strings.TrimSpace(user.UserName) == "" && strings.TrimSpace(user.Email.String) == "" && strings.TrimSpace(user.PhoneNo.String) == ""
}
