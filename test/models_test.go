package test

import (
	"fmt"
	"testing"

	"../models"
)

func TestCreateUser(t *testing.T) {
	abe := models.CreateUser("abe", "abe@example.com")
	bobby := models.CreateUser("bobby", "bobby@example.com")
	fmt.Println(abe)
	fmt.Println(bobby)
	if abe.Name != "abe" || abe.Email != "abe@example.com" {
		t.Error("User name incorrect")
	}
	if abe.ID == bobby.ID {
		t.Error("UUID not unique")
	}
}
