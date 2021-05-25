package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID       int64          // Unique User ID
	UserName string         // Username
	FName    string         // First Name
	LName    string         // Last Name
	Email    sql.NullString // Email
	PhoneNo  sql.NullString // Phone Number
	Password string         // Password
}

func (u User) String() string {
	if u.ID != 0 {
		return fmt.Sprintf("USER %v %v, (USERNAME %v and ID %v)", u.FName, u.LName, u.UserName, u.ID)
	}

	return fmt.Sprintf("USER %v %v, (USERNAME %v)", u.FName, u.LName, u.UserName)
}
