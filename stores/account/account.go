package account

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/stores"
)

type account struct{}

func New() stores.Account {
	return account{}
}

const (
	getAll = "SELECT id, user_name, email, f_name, l_name, phone_no, created_at, pwd_update, del_req, status FROM accounts WHERE "
	get    = "SELECT id, user_name, email, password, f_name, l_name, phone_no, created_at, pwd_update, del_req, status FROM accounts WHERE "
	insert = "INSERT INTO accounts( user_name, password, email, f_name, l_name, phone_no, status) VALUES(?, ?, ?, ?, ?, ?, ?)"
)

// GetAll retrieves all accounts that match the given filter.
func (a account) GetAll(ctx *app.Context, filter *models.Account) ([]*models.Account, error) {
	where, qp := filter.WhereClause()
	query := getAll + where

	accounts := make([]*models.Account, 0)

	rows, err := ctx.SQL.QueryContext(ctx, query, qp...)
	if err != nil {
		return nil, errors.DBError{Err: err}
	}

	defer rows.Close()

	for rows.Next() {
		var account models.Account
		err := rows.Scan(&account.ID, &account.UserName, &account.Email, &account.FName, &account.LName, &account.PhoneNo, &account.CreatedAt, &account.PwdUpdate, &account.DelRequest, &account.Status)

		if err != nil {
			return nil, errors.DBError{Err: err}
		}

		accounts = append(accounts, &account)
	}

	ctx.Debugf("successful execution of 'GetAll' accounts in storage layer")

	return accounts, nil
}

// Get retrieves a single account that matches a given filter.
func (a account) Get(ctx *app.Context, filter *models.Account) (*models.Account, error) {
	db := ctx.SQL

	where, qp := filter.WhereClause()
	query := get + where

	rows, err := db.QueryContext(ctx, query, qp...)
	if err != nil {
		return nil, errors.DBError{Err: err}
	}

	defer rows.Close()

	var account models.Account

	for rows.Next() {
		err := rows.Scan(&account.ID, &account.UserName, &account.Email, &account.Password, &account.FName, &account.LName, &account.PhoneNo, &account.CreatedAt, &account.PwdUpdate, &account.DelRequest, &account.Status)
		if err != nil {
			return nil, errors.DBError{Err: err}
		}
	}

	if reflect.DeepEqual(account, models.Account{}) {
		return nil, errors.EntityNotFound{Entity: "account"}
	}

	return &account, nil
}

// Create creates an account.
func (a account) Create(ctx *app.Context, model *models.Account) (*models.Account, error) {
	db := ctx.SQL

	res, err := db.ExecContext(ctx, insert, model.UserName, model.Password, model.Email, model.FName, model.LName, model.PhoneNo, model.Status)
	if err != nil {
		return nil, errors.DBError{Err: err}
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, errors.DBError{Err: err}
	}

	account, err := a.Get(ctx, &models.Account{User: models.User{ID: id}})
	if err != nil {
		return nil, err
	}

	return account, nil
}

// Update updates an account.
func (a account) Update(ctx *app.Context, model *models.Account) (*models.Account, error) {
	db := ctx.SQL

	query, qp := generateSetClause(model)

	query = fmt.Sprintf("%s WHERE id = ?;", query)

	qp = append(qp, model.ID)

	_, err := db.ExecContext(ctx, query, qp...)
	if err != nil {
		return nil, errors.DBError{Err: err}
	}

	return a.Get(ctx, &models.Account{User: models.User{ID: model.ID}})
}

func generateSetClause(model *models.Account) (setClause string, qp []interface{}) {
	setClause = "UPDATE accounts SET"

	if model.UserName != "" {
		setClause += " user_name = ?,"

		qp = append(qp, model.UserName)
	}

	if model.Password != "" {
		setClause += " password = ?, pwd_update = ?,"

		qp = append(qp, model.Password, model.PwdUpdate)
	}

	if model.Email.Valid {
		setClause += " email = ?,"

		qp = append(qp, model.Email)
	}

	if model.FName != "" {
		setClause += " f_name = ?,"

		qp = append(qp, model.FName)
	}

	if model.LName != "" {
		setClause += " l_name = ?,"

		qp = append(qp, model.LName)
	}

	if model.PhoneNo.Valid {
		setClause += " phone_no = ?,"

		qp = append(qp, model.PhoneNo)
	}

	if model.Status != "" {
		setClause += " status = ?,"

		qp = append(qp, model.Status)
	}

	setClause += " del_req = ?,"

	qp = append(qp, model.DelRequest)

	setClause = strings.TrimSuffix(setClause, ",")

	return setClause, qp
}

// Delete updates a delete request for an account and sets its status to inactive.
// Account is then permanently deleted after 30 days of inactivity.
func (a account) Delete(ctx *app.Context, id int64) error {
	_, err := ctx.SQL.ExecContext(ctx, "UPDATE accounts SET del_req = ?, status = ? WHERE id = ?", time.Now(), "INACTIVE", id)
	if err != nil {
		return errors.DBError{Err: err}
	}

	return nil
}
