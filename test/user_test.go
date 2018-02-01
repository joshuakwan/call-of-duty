package test

import (
	"testing"

	"../utils"
)

func TestCreateUser(t *testing.T) {
	users := createTestUsers()
	number := len(users)
	dup := false
	for i := 0; i < number; i++ {
		for j := i + 1; j < number; j++ {
			if users[i].ID == users[j].ID {
				dup = true
				break
			}
		}
	}

	utils.AssertNotEqual(t, dup, true, "UUID not unique")
}
