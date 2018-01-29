package test

import (
	"testing"
	"time"

	"../models"
	"../utils"
)

func TestCreateUser(t *testing.T) {
	abe := models.CreateUser("abe", "abe@example.com")
	bobby := models.CreateUser("bobby", "bobby@example.com")
	utils.AssertEqual(t, abe.Name, "abe", "User name incorrect")
	utils.AssertEqual(t, abe.Email, "abe@example.com", "User email incorrect")
	utils.AssertNotEqual(t, abe.ID, bobby.ID, "UUID not unique")
}

func TestTeamFunction(t *testing.T) {
	teamName := "Team A"
	teamDescription := "A Testing Team"

	teamA := models.CreateTeam(teamName, teamDescription)

	utils.AssertEqual(t, teamA.Name, teamName, "Team name incorrect")
	utils.AssertEqual(t, teamA.Description, teamDescription, "Team description incorrect")

	abe := models.CreateUser("abe", "abe@example.com")
	bobby := models.CreateUser("bobby", "bobby@example.com")

	teamA.AddNewUser(abe)
	teamA.AddNewUser(bobby)

	utils.AssertEqual(t, len(teamA.UserIDs), 2, "Incorrect number of users")
	foundAbe := false
	foundBobby := false
	for _, userID := range teamA.UserIDs {
		if userID == abe.ID {
			foundAbe = true
		}
		if userID == bobby.ID {
			foundBobby = true
		}
	}
	utils.AssertEqual(t, foundAbe, true, "Abe not added")
	utils.AssertEqual(t, foundBobby, true, "Bobby not added")

	teamA.RemoveUser(abe)

	foundAbe = false
	for _, userID := range teamA.UserIDs {
		if userID == abe.ID {
			foundAbe = true
		}
	}
	utils.AssertEqual(t, foundAbe, false, "Abe not removed")

	teamA.RemoveUser(bobby)
	utils.AssertEqual(t, len(teamA.UserIDs), 0, "Not all users removed")
}

func TestCreateCalendarEntry(t *testing.T) {
	dateStart := time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)
	dateDuration := 40
	dateEnd := dateStart.AddDate(0, 0, dateDuration)
	timeSlot := "08:00-23:00"
	newCalEntry := models.CreateCalendarEntry(dateStart, dateEnd, "", timeSlot)
	timeSlots := newCalEntry.TimeSlots

	utils.AssertEqual(t, len(timeSlots), dateDuration+1, "Incorrect time slot length")

	for _, entry := range timeSlots {
		utils.AssertEqual(t, entry.Start.Hour(), 8, "Incorrect start hour")
		utils.AssertEqual(t, entry.Start.Minute(), 0, "Incorrect start min")
		utils.AssertEqual(t, entry.End.Hour(), 23, "Incorrect end hour")
		utils.AssertEqual(t, entry.End.Minute(), 0, "Incorrect end min")
	}
}
