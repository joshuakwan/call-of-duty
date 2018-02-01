package test

import (
	"../models"
)

var testInputUsers = []struct {
	name  string
	email string
}{
	{"abe", "abe@example.com"},
	{"bobby", "bobby@example.com"},
	{"carl", "carl@example.com"},
	{"dick", "dick@example.com"},
	{"erick", "erick@example.com"},
	{"fredric", "fredric@example.com"},
	{"gabriel", "gabriel@example.com"},
	{"hana", "hana@example.com"},
	{"ibe", "ibe@example.com"},
	{"jojo", "jojo@example.com"},
}

func createTestUsers() []*models.User {
	var users []*models.User
	for _, input := range testInputUsers {
		users = append(users, models.CreateUser(input.name, input.email))
	}

	return users
}

func createTestServices() []*models.Service {
	return nil
}

func createTestTeam() *models.Team {
	teamName := "Team A"
	teamDescription := "A Testing Team"
	return &models.Team{Name: teamName, Description: teamDescription}
}
