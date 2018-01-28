package test

import (
	"testing"
	"time"

	"../models"
)

func TestCreateUser(t *testing.T) {
	abe := models.CreateUser("abe", "abe@example.com")
	bobby := models.CreateUser("bobby", "bobby@example.com")
	if abe.Name != "abe" || abe.Email != "abe@example.com" {
		t.Error("User name incorrect")
	}
	if abe.ID == bobby.ID {
		t.Error("UUID not unique")
	}
}

func TestCreateCalendarEntry(t *testing.T) {
	dateStart := time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)
	dateDuration := 40
	dateEnd := dateStart.AddDate(0, 0, dateDuration)
	timeSlot := "08:00-23:00"
	newCalEntry := models.CreateCalendarEntry(dateStart, dateEnd, "", timeSlot)
	timeSlots := newCalEntry.TimeSlots

	if len(timeSlots) != dateDuration+1 {
		t.Error("Incorrect time slot length")
	}

	for _, entry := range timeSlots {
		if entry.Start.Hour() != 8 {
			t.Error("Incorrect start hour")
		}
		if entry.Start.Minute() != 0 {
			t.Error("Incorrect start min")
		}
		if entry.End.Hour() != 23 {
			t.Error("Incorrect end hour")
		}
		if entry.End.Minute() != 0 {
			t.Error("Incorrect end min")
		}
	}
}
