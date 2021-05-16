package account

import (
	"context"
	"database/sql"
	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"github.com/Aakanksha-jais/picshot-golang-backend/stores"
	"reflect"
	"strings"
	"time"
)

type account struct {
	db     *sql.DB
	logger log.Logger
}

func New(db *sql.DB, logger log.Logger) stores.Account {
	return account{db: db, logger: logger}
}

const (
	getAll = "SELECT id, user_name, email, f_name, l_name, phone_no, created_at, pwd_update, del_req, status FROM accounts WHERE "
	get    = "SELECT id, user_name, email, password, f_name, l_name, phone_no, created_at, pwd_update, del_req, status FROM accounts WHERE "
	insert = "INSERT INTO accounts( user_name, password, email, f_name, l_name, phone_no, pwd_update, status) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
)

// GetAll retrieves all accounts that match the given filter.
func (a account) GetAll(ctx context.Context, filter *models.Account) ([]*models.Account, error) {
	where, qp := filter.WhereClause()
	query := getAll + where

	accounts := make([]*models.Account, 0)

	rows, err := a.db.QueryContext(ctx, query, qp...)
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

	a.logger.Debugf("successful execution of 'GetAll' accounts in storage layer")
	return accounts, nil
}

// Get retrieves a single account that matches a given filter.
func (a account) Get(ctx context.Context, filter *models.Account) (*models.Account, error) {
	where, qp := filter.WhereClause()
	query := get + where

	rows, err := a.db.QueryContext(ctx, query, qp...)
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
func (a account) Create(ctx context.Context, model *models.Account) (*models.Account, error) {
	res, err := a.db.ExecContext(ctx, insert, model.UserName, model.Password, model.Email, model.FName, model.LName, model.PhoneNo, time.Now(), model.Status)
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
func (a account) Update(ctx context.Context, model *models.Account) (*models.Account, error) {

	query, qp := generateSetClause(model)

	res, err := a.db.ExecContext(ctx, query, qp...)
	if err != nil {
		return nil, errors.DBError{Err: err}
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, errors.DBError{Err: err}
	}

	return a.Get(ctx, &models.Account{User: models.User{ID: id}})
}

func generateSetClause(model *models.Account) (setClause string, qp []interface{}) {
	setClause = `UPDATE accounts SET`

	if model.UserName != "" {
		setClause += ` user_name = ?,`

		qp = append(qp, model.UserName)
	}

	if !reflect.DeepEqual(model.Password, []byte{}) {
		setClause += ` password = ?, pwd_update = ?,`

		qp = append(qp, model.Password, time.Now())
	}

	if model.Email.String != "" {
		setClause += ` email = ?,`

		qp = append(qp, model.Email.String)
	}

	if model.FName != "" {
		setClause += ` f_name = ?,`

		qp = append(qp, model.FName)
	}

	if model.LName != "" {
		setClause += ` l_name = ?,`

		qp = append(qp, model.LName)
	}

	if model.PhoneNo.String != "" {
		setClause += ` phone_no = ?,`

		qp = append(qp, model.PhoneNo.String)
	}

	if model.Status != "ACTIVE" {
		setClause += ` status = ?,`

		qp = append(qp, "ACTIVE")
	}

	// todo update del req
	if !model.DelRequest.Valid {
		setClause += ` del_req = ?,`

		qp = append(qp, "NULL") //set to null
	}

	setClause = strings.TrimSuffix(setClause, ",")

	return setClause, qp
}

// Delete updates a delete request for an account and sets its status to inactive.
// Account is then permanently deleted after 30 days of inactivity.
func (a account) Delete(ctx context.Context, id int64) error {
	_, err := a.db.ExecContext(ctx, "UPDATE accounts SET del_req = ?, status = ? WHERE id = ?", time.Now(), "INACTIVE", id)
	if err != nil {
		return errors.DBError{Err: err}
	}

	// TODO: trigger a cronjob for 30 days deletion functionality
	return nil
}
