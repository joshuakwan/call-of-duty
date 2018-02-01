package models

import (
	"testing"

	"github.com/joshuakwan/call-of-duty/utils"
)

func TestAddTeamServices(t *testing.T) {
	teamA := createTestTeam()
	services := createTestServices()

	for _, service := range services {
		teamA.AddNewService(service)
	}

	utils.AssertEqual(t, len(teamA.ServiceIDs), len(services), "Incorrect number of services")

	for _, service := range services {
		utils.AssertEqual(t, service.TeamID, teamA.ID, "TeamID not assigned to the service")
	}
}

func TestRemoveTeamServices(t *testing.T) {
	teamA := createTestTeam()
	services := createTestServices()

	for _, service := range services {
		teamA.AddNewService(service)
	}

	for _, service := range services {
		serviceID := teamA.RemoveService(service)
		utils.AssertEqual(t, serviceID, service.ID, "ServiceID not equal")
		utils.AssertEqual(t, service.TeamID, "", "TeamID not removed from the service")
	}

	utils.AssertEqual(t, len(teamA.ServiceIDs), 0, "Services should be empty")
}

func TestAddTeamUsers(t *testing.T) {
	teamA := createTestTeam()
	users := createTestUsers()

	for _, user := range users {
		teamA.AddNewUser(user)
	}

	// Verify number of users
	utils.AssertEqual(t, len(teamA.UserIDs), len(users), "Incorrect number of users")

	// Verify if TeamID has been added
	for _, user := range users {
		utils.AssertNotEqual(t, utils.Contains(teamA.ID, user.TeamIDs), -1, "TeamID not added to the user")
	}
}

func TestRemoveTeamUsers(t *testing.T) {
	teamA := createTestTeam()
	users := createTestUsers()

	for _, user := range users {
		teamA.AddNewUser(user)
	}

	for _, user := range users {
		userID := teamA.RemoveUser(user)
		utils.AssertEqual(t, userID, user.ID, "UserID not equal")
		utils.AssertEqual(t, utils.Contains(teamA.ID, user.TeamIDs), -1, "TeamID not removed from the user")
	}

	utils.AssertEqual(t, len(teamA.UserIDs), 0, "Users should be empty")
}
