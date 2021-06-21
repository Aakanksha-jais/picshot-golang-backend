package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID       int64          `json:"id"`        // Unique User ID
	UserName string         `json:"user_name"` // Username
	FName    string         `json:"f_name"`    // First Name
	LName    string         `json:"l_name"`    // Last Name
	Email    sql.NullString `json:"email"`     // Email
	PhoneNo  sql.NullString `json:"phone_no"`  // Phone Number
	Password string         `json:"password"`  // Password
}

type VerificationResponse struct {
	URL string `json:"url"`
	SID string `json:"sid"`
}

func (u User) String() string {
	if u.ID != 0 {
		return fmt.Sprintf("USER %v %v, (USERNAME %v and ID %v)", u.FName, u.LName, u.UserName, u.ID)
	}

	return fmt.Sprintf("USER %v %v, (USERNAME %v)", u.FName, u.LName, u.UserName)
}
