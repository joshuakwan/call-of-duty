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
	TeamIDs            []string `json:"team_ids"`
	EscalationChainIDs []string `json:"escalation_chain_ids"`
}

// CreateUser creates a new user
func CreateUser(name, email string) *User {
	return &User{ID: utils.GenerateUUID(), Name: name, Email: email}
}

// RemoveFromTeam removes a Team from a User
func (user *User) removeFromTeam(teamID string) string {
	for i, id := range user.TeamIDs {
		if teamID == id {
			user.TeamIDs = utils.RemoveElement(user.TeamIDs, i)
			break
		}
	}
	return teamID
}

// String returns the literal form of a User
func (user *User) String() string {
	return fmt.Sprintf("UUID: %s, User Name: %s, Email: %s", user.ID, user.Name, user.Email)
}
