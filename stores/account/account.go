package account

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Aakanksha-jais/picshot-golang-backend/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/filters"
	"github.com/Aakanksha-jais/picshot-golang-backend/log"
	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"strings"
	"time"
)

type account struct {
	db     *sql.DB
	logger log.Logger
}

func New(db *sql.DB) account {
	return account{db: db}
}

const (
	UserName   = `user_name`
	Email      = `email`
	Password   = `password`
	FName      = `f_name`
	LName      = `l_name`
	PhoneNo    = `phone_no`
	ID         = `id`
	CreatedAt  = `created_at`
	PwdUpdate  = `pwd_update`
	DelRequest = `del_req`
	Status     = `status`
)

// GetAll retrieves all accounts that match the given filter.
func (ac account) GetAll(ctx context.Context, filter *filters.Account) ([]*models.Account, error) {
	where, qp := filter.WhereClause()
	query := fmt.Sprintf(
		"SELECT %s, %s, %s , %s, %s, %s, %s, %s, %s, %s FROM accounts WHERE %s",
		ID,
		UserName,
		Email,
		FName,
		LName,
		PhoneNo,
		CreatedAt,
		PwdUpdate,
		DelRequest,
		Status,
		where,
	)

	accounts := make([]*models.Account, 0)

	rows, err := ac.db.QueryContext(ctx, query, qp...)
	if err != nil {
		return nil, errors.DBError{Err: err}
	}

	defer rows.Close()

	for rows.Next() {
		var account models.Account
		err := rows.Scan(
			&account.ID,
			&account.UserName,
			&account.Email,
			&account.FName,
			&account.LName,
			&account.PhoneNo,
			&account.CreatedAt,
			&account.PwdUpdate,
			&account.DelRequest,
			&account.Status,
		)

		if err != nil {
			return nil, errors.DBError{Err: err}
		}

		accounts = append(accounts, &account)
	}

	return accounts, nil
}

// Get retrieves a single account that matches a given filter.
func (ac account) Get(ctx context.Context, filter *filters.Account) (*models.Account, error) {
	where, qp := filter.WhereClause()
	query := fmt.Sprintf(
		"SELECT %s, %s, %s , %s, %s, %s, %s, %s, %s, %s FROM accounts WHERE %s",
		ID,
		UserName,
		Email,
		FName,
		LName,
		PhoneNo,
		CreatedAt,
		PwdUpdate,
		DelRequest,
		Status,
		where,
	)

	rows, err := ac.db.QueryContext(ctx, query, qp...)
	if err != nil {
		return nil, errors.DBError{Err: err}
	}

	defer rows.Close()

	var account models.Account

	for rows.Next() {
		err := rows.Scan(
			&account.ID,
			&account.UserName,
			&account.Email,
			&account.FName,
			&account.LName,
			&account.PhoneNo,
			&account.CreatedAt,
			&account.PwdUpdate,
			&account.DelRequest,
			&account.Status,
		)

		if err == sql.ErrNoRows {
			return nil, errors.EntityNotFound{Entity: "account"}
		}

		if err != nil {
			return nil, errors.DBError{Err: err}
		}
	}

	return &account, nil
}

// Create creates an account.
func (ac account) Create(ctx context.Context, model *models.Account) (*models.Account, error) {
	query := fmt.Sprintf(
		"INSERT INTO accounts( %s, %s, %s, %s, %s, %s, %s, %s) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)",
		UserName,
		Password,
		Email,
		FName,
		LName,
		PhoneNo,
		PwdUpdate,
		Status,
	)

	res, err := ac.db.ExecContext(
		ctx,
		query,
		model.UserName,
		model.Password,
		model.Email,
		model.FName,
		model.LName,
		model.PhoneNo,
		time.Now(),
		model.Status,
	)
	if err != nil {
		return nil, errors.DBError{Err: err}
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, errors.DBError{Err: err}
	}

	account, err := ac.Get(ctx, &filters.Account{ID: id})
	if err != nil {
		return nil, err
	}

	return account, nil
}

// Update updates an account.
func (ac account) Update(ctx context.Context, model *models.Account) (*models.Account, error) {

	query, qp := generateSetClause(model)

	res, err := ac.db.ExecContext(ctx, query, qp...)
	if err != nil {
		return nil, errors.DBError{Err: err}
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, errors.DBError{Err: err}
	}

	if model.Password != "" {
		_, err = ac.db.ExecContext(ctx, "UPDATE accounts SET pwd_update = ? WHERE id = ?", time.Now(), id)
		if err != nil {
			return nil, errors.DBError{Err: err}
		}
	}

	return ac.Get(ctx, &filters.Account{ID: id})
}

func generateSetClause(model *models.Account) (setClause string, qp []interface{}) {
	setClause = `set`

	if model.UserName != "" {
		setClause += ` user_name = ?,`

		qp = append(qp, model.UserName)
	}

	if model.Password != "" {
		setClause += ` password = ?,`

		qp = append(qp, model.Password)
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

	setClause = strings.TrimSuffix(setClause, ",")
	setClause = strings.TrimSuffix(setClause, "set")

	return setClause, qp
}

// Delete updates a delete request for an account and sets its status to inactive.
// Account is then permanently deleted after 30 days of inactivity.
func (ac account) Delete(ctx context.Context, id int64) error {
	_, err := ac.db.ExecContext(ctx, "UPDATE accounts SET del_req = ?, status = ? WHERE id = ?", time.Now(), "INACTIVE", id)
	if err != nil {
		return errors.DBError{Err: err}
	}

	// TODO: trigger a cronjob for 30 days deletion functionality
	return nil
}
