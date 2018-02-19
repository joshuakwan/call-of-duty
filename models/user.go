package models

import (
	"errors"
	"fmt"

	"github.com/joshuakwan/call-of-duty/db"
	"github.com/joshuakwan/call-of-duty/utils"
	"gopkg.in/mgo.v2/bson"
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

// QueryUserByEmail queries user by an email address
func QueryUserByEmail(email string) (User, error) {
	var result []User
	err := db.FindDBObjects("models.User", bson.M{"email": email}, &result)

	if len(result) > 1 {
		err = errors.New("user " + email + " has duplicated entries")
	}

	return result[0], err
}

// CreateUser creates a new user
func CreateUser(name, email string) (*User, error) {
	user, err := QueryUserByEmail(email)

	if err != nil {
		return nil, err
	}

	if &user != nil {
		return nil, errors.New("user " + email + " exists")
	}

	user = User{ID: utils.GenerateUUID(), Name: name, Email: email}
	err = db.CreateDBObject(user)
	return &user, err
}

// Delete removes a user itself from DB
func (user *User) Delete() error {
	err := db.RemoveDBObject("models.User", bson.M{"email": user.Email})
	return err
}

// Update updates a user itself
func (user *User) Update() error {
	err := db.UpdateDBObject("models.User", bson.M{"email": user.Email}, user)
	return err
}

// RemoveFromTeam removes a Team from a User
func (user *User) RemoveFromTeam(teamID string) string {
	for i, id := range user.TeamIDs {
		if teamID == id {
			user.TeamIDs = utils.RemoveElement(user.TeamIDs, i)
			break
		}
	}
	user.Update()
	return teamID
}

// String returns the literal form of a User
func (user *User) String() string {
	return fmt.Sprintf("UUID: %s, User Name: %s, Email: %s", user.ID, user.Name, user.Email)
}
