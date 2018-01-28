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
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	Email               string   `json:"email"`
	TeamIDs             string   `json:"team_ids"`
	EscalationPolicyIDs []string `json:"escalation_policy_ids"`
}

// CreateUser creates a new user
func CreateUser(name string, email string) *User {
	uuid := utils.GenerateUUID()
	newUser := &User{ID: uuid, Name: name, Email: email}
	return newUser
}

func (u *User) String() string {
	return fmt.Sprintf("UUID: %s, User Name: %s, Email: %s", u.ID, u.Name, u.Email)
}
