package models

import (
	"fmt"

	"../utils"
)

// UserPreference defines the preference of a user
type UserPreference struct {
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	WeChat string `json:"wechat"`
}

// User defines an application user
type User struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	Email              string   `json:"email"`
	TeamIDs            string   `json:"team_ids"`
	EscalationChainIDs []string `json:"escalation_chain_ids"`
}

// CreateUser creates a new user
func CreateUser(name, email string) *User {
	return &User{ID: utils.GenerateUUID(), Name: name, Email: email}
}

func (u *User) String() string {
	return fmt.Sprintf("UUID: %s, User Name: %s, Email: %s", u.ID, u.Name, u.Email)
}
