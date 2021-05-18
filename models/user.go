package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID       int64          `json:"-"`                  // Unique User ID
	UserName string         `json:"user_name"`          // Username
	FName    string         `json:"f_name"`             // First Name
	LName    string         `json:"l_name"`             // Last Name
	Email    sql.NullString `json:"email,omitempty"`    // Email
	PhoneNo  sql.NullString `json:"phone_no,omitempty"` // Phone Number
	Password string         `json:"password,omitempty"` // Password
}

func (u User) String() string {
	if u.ID != 0 {
		return fmt.Sprintf("User: %v %v, (Username: %v and ID: %v)", u.FName, u.LName, u.UserName, u.ID)
	}

	return fmt.Sprintf("User: %v %v, (Username: %v)", u.FName, u.LName, u.UserName)
}
