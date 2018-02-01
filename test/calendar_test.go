package test

import (
	"testing"
	"time"

	"../models"
	"../utils"
)

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
