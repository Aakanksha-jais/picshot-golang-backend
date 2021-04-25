package models

import (
	"database/sql"
	"time"
)

type Account struct {
	User                    // Details of the User
	ID         int64        `json:"id"`         // Unique Account ID
	Password   string       `json:"password"`   // Password
	PwdUpdate  time.Time    `json:"pwd_update"` // Time Stamp of most recent Password Update
	Blogs      []Blog       `json:"blogs"`      // List of Blogs posted by Account
	CreatedAt  time.Time    `json:"created_at"` // Time of Creation of Account
	DelRequest sql.NullTime `json:"del_req"`    // Time Stamp of Account Delete Request
	Status     string       `json:"status"`     // Account Active or Inactive
}

type User struct {
	UserName string         `json:"user_name"`
	Email    sql.NullString `json:"email"`
	FName    string         `json:"f_name"`
	LName    string         `json:"l_name"`
	PhoneNo  sql.NullString `json:"phone_no"`
}