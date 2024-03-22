package main

import (
	utl "autilities"
	"time"
)

var header = utl.Header{}

/*
A common requirement in programs is getting the number of seconds, milliseconds, or nanoseconds since the Unix epoch.
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing Epoch")

	// Use time.Now with Unix, UnixMilli or UnixNano to get elapsed time since the Unix epoch in seconds, milliseconds or nanoseconds, respectively.
	now := time.Now()
	utl.PLine("Now: ", now)
	utl.PLine("Seconds: ", now.Unix())
	utl.PLine("Milliseconds: ", now.UnixMilli())
	utl.PLine("Nanoseconds: ", now.UnixNano())

	// You can also convert integer seconds or nanoseconds since the epoch into the corresponding time.
	utl.PLine("Time: ", time.Unix(now.Unix(), 0))
	utl.PLine("Time: ", time.Unix(0, now.UnixNano()))

}
