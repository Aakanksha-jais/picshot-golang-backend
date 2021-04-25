package filters

import (
	"reflect"
	"strings"
	"time"
)

type Account struct {
	User
	ID        int64
	CreatedAt time.Time
	Status    string
}

type User struct {
	UserName string
	Email    string
	FName    string
	LName    string
	PhoneNo  string
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

func (a *Account) WhereClause() (whereClause string, queryParams []interface{}) {
	columnList := make([]string, 0)
	queryParams = make([]interface{}, 0)

	if a.UserName != "" {
		columnList = append(columnList, UserName)
		queryParams = append(queryParams, a.UserName)
	}

	if a.Email != "" {
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

	if a.PhoneNo != "" {
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
