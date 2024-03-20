package main

import (
	utl "autilities"
	"time"
)

var header = utl.Header{}

var p = utl.PLine

/*
Go offers extensive support for times and durations; here are some examples.
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing Time")

	// We’ll start by getting the current time.
	now := time.Now()
	p("Now: ", now)

	// You can build a time struct by providing the year, month, day, etc. Times are always associated with a Location, i.e. time zone.
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p("Then: ", then)

	// You can extract the various components of the time value as expected.
	p("Year: ", then.Year())
	p("Month: ", then.Month())
	p("Day: ", then.Day())
	p("Hour: ", then.Hour())
	p("Minute: ", then.Minute())
	p("Second: ", then.Second())
	p("Nanosecond: ", then.Nanosecond())
	p("Location: ", then.Location())

	// The Monday-Sunday Weekday is also available.
	p("Weekday: ", then.Weekday())

	// These methods compare two times, testing if the first occurs before, after, or at the same time as the second, respectively.
	p("Before: ", then.Before(now))
	p("After: ", then.After(now))
	p("Equal: ", then.Equal(now))

	// The Sub methods returns a Duration representing the interval between two times.
	diff := now.Sub(then)
	p("Diff: ", diff)

	// We can compute the length of the duration in various units.
	p("Hours: ", diff.Hours())
	p("Minutes: ", diff.Minutes())
	p("Seconds: ", diff.Seconds())
	p("Nanoseconds: ", diff.Nanoseconds())

	// You can use Add to advance a time by a given duration, or with a - to move backwards by a duration.
	p("Add: ", then.Add(diff))
	p("Sub: ", then.Add(-diff))
}
