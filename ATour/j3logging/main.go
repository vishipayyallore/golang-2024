package main

import (
	utl "autilities"
	"log"
	"os"
)

var header = utl.Header{}

/*
The Go standard library provides straightforward tools for outputting logs from Go programs, with the log package for
free-form output and the log/slog package for structured output.
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing Logging")

	// Simply invoking functions like Println from the log package uses the standard logger, which is already pre-configured
	// for reasonable logging output to os.Stderr. Additional methods like Fatal* or Panic* will exit the program after logging.
	log.Println("standard logger")

	// Loggers can be configured with flags to set their output format. By default, the standard logger has the log.Ldate
	// and log.Ltime flags set, and these are collected in log.LstdFlags. We can change its flags to emit time with microsecond
	// accuracy, for example.
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// It also supports emitting the file name and line from which the log function is called.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// It may be useful to create a custom logger and pass it around. When creating a new logger, we can set a prefix to
	// distinguish its output from other loggers.
	custlog := log.New(os.Stdout, "my:", log.LstdFlags)
	custlog.Println("from mylog")

	// We can set the prefix on existing loggers (including the standard one) with the SetPrefix method.
	custlog.SetPrefix("ohmy:")
	custlog.Println("from mylog")
}
