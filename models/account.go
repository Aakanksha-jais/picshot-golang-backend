package models

import (
	"database/sql"
	"reflect"
	"strings"
	"time"
)

type Account struct {
	User                    // Details of the User
	ID         int64        `json:"id"`         // Unique Account ID
	PwdUpdate  time.Time    `json:"pwd_update"` // Time Stamp of most recent Password Update
	Blogs      []Blog       `json:"blogs"`      // List of Blogs posted by Account
	CreatedAt  time.Time    `json:"created_at"` // Time of Creation of Account
	DelRequest sql.NullTime `json:"del_req"`    // Time Stamp of Account Delete Request
	Status     string       `json:"status"`     // Account Active or Inactive
}

type User struct {
	UserName string         `json:"user_name"` // Username
	FName    string         `json:"f_name"`    // First Name
	LName    string         `json:"l_name"`    // Last Name
	Email    sql.NullString `json:"email"`     // Email
	PhoneNo  sql.NullString `json:"phone_no"`  // Phone Number
	Password []byte         `json:"password"`  // Password
}

const (
	UserName  = `user_name`
	Email     = `email`
	FName     = `f_name`
	LName     = `l_name`
	PhoneNo   = `phone_no`
	ID        = `id`
	CreatedAt = `created_at`
	Status    = `status`
)

func (a *Account) WhereClause() (whereClause string, queryParams []interface{}) {
	columnList := make([]string, 0)
	queryParams = make([]interface{}, 0)

	if a.UserName != "" {
		columnList = append(columnList, UserName)
		queryParams = append(queryParams, a.UserName)
	}

	if a.Email.String != "" {
		columnList = append(columnList, Email)
		queryParams = append(queryParams, a.Email)
	}

	if a.FName != "" {
		columnList = append(columnList, FName)
		queryParams = append(queryParams, a.FName)
	}

	if a.LName != "" {
		columnList = append(columnList, LName)
		queryParams = append(queryParams, a.LName)
	}

	if a.PhoneNo.String != "" {
		columnList = append(columnList, PhoneNo)
		queryParams = append(queryParams, a.PhoneNo)
	}

	if a.ID != 0 {
		columnList = append(columnList, ID)
		queryParams = append(queryParams, a.ID)
	}

	if !reflect.DeepEqual(a.CreatedAt, time.Time{}) {
		columnList = append(columnList, CreatedAt)
		queryParams = append(queryParams, a.CreatedAt)
	}

	if a.Status != "" {
		columnList = append(columnList, Status)
		queryParams = append(queryParams, a.Status)
	}

	whereClause = strings.Join(columnList, `= ? and `) + `= ? `

	return whereClause, queryParams
}
