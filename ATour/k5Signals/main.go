package main

import (
	utl "autilities"
)

var header = utl.Header{}

/*
Sometimes we’d like our Go programs to intelligently handle Unix signals. For example, we might want a server to gracefully shutdown
when it receives a SIGTERM, or a command-line tool to stop processing input if it receives a SIGINT. Here’s how to handle signals
in Go with channels.
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing Signals")

}
