package models

import (
	"strconv"
	"strings"
	"time"

	"../utils"
)

// TimeSlot defines a slot of start_time ~ end_time
type TimeSlot struct {
	Start time.Time // Start time
	End   time.Time // End time
}

/*
CalendarEntry defines a calendar entry made of days/hours
It should contain:
  * start
  * end
  * days
  * hours
The date is recorded in UTC, as a series of EPOCH time slots
*/
type CalendarEntry struct {
	ID string

	/*
		PlainText to describe the calendar in human readable way

		Sample:
		  * Start:2018-01-01 | End:2018-01-03 | Days:<Day Expression> | Time:0800-2200

		  Day Expression:
			* daily
			* weekdays
			* weekends
			* every mon
			* every tue,wed,fri
	*/
	TimeSlots []TimeSlot
}

// CreateTimeSlot creates a new time slot
// func CreateTimeSlot(timeStart, timeEnd time.Time) *TimeSlot {
// 	return TimeSlot(timeStart, time)
// }

/* CreateCalendarEntry creates a new CalendarEntry
timeSlot: 08:10-09:20
TODO: dayExpression not implemented yet
*/
func CreateCalendarEntry(dateStart, dateEnd time.Time, dayExpression string, timeSlot string) *CalendarEntry {
	uuid := utils.GenerateUUID()
	var timeSlots []TimeSlot

	hours := strings.Split(timeSlot, "-")
	startMinutes := strings.Split(hours[0], ":")
	startHour, _ := strconv.Atoi(startMinutes[0])
	startMinute, _ := strconv.Atoi(startMinutes[1])
	endMinutes := strings.Split(hours[1], ":")
	endHour, _ := strconv.Atoi(endMinutes[0])
	endMinute, _ := strconv.Atoi(endMinutes[1])

	daysBetween := (dateEnd.Sub(dateStart)).Hours() / 24

	for i := 0; i <= int(daysBetween); i++ {
		day := dateStart.AddDate(0, 0, i)
		timeStart := time.Date(day.Year(),
			day.Month(),
			day.Day(),
			startHour,
			startMinute,
			0, 0, time.UTC)
		timeEnd := time.Date(day.Year(),
			day.Month(),
			day.Day(),
			endHour,
			endMinute,
			0, 0, time.UTC)

		timeSlots = append(timeSlots, TimeSlot{Start: timeStart, End: timeEnd})
	}

	return &CalendarEntry{ID: uuid, TimeSlots: timeSlots}
}
