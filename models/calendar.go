package models

// TimeSlot defines a slot of start_time ~ end_time
type TimeSlot struct {
	Start int64 // Start time
	End   int64 // End time
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
		  * Start:2018/01/01 | End:2018/01/03 | Days:<Day Expression> | Time:0800-2200

		  Day Expression:
			* daily
			* weekdays
			* weekends
			* every mon
			* every tue,wed,fri
	*/
	PlainText string
	TimeSlots []TimeSlot
}
