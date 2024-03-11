package main

import (
	utl "autilities"
	"fmt"
	"os"
	"time"
)

var header = utl.Header{}

/*
Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup.
defer is often used where e.g. ensure and finally would be used in other languages.
*/
func main() {
	// A defer statement defers the execution of a function until the surrounding function returns.
	header.DisplayHeader("Showing Defer")

	// The deferred call's arguments are evaluated immediately,
	// but the function call is not executed until the surrounding function returns.
	defer utl.PLine("world")

	utl.PLine("hello")

	showDefer()

	defer showDeferV2()

	showFileDefer()
}

func showDefer() {
	defer utl.PLine("showDefer() world !!")

	utl.PLine("showDefer() hello ")
}

func showDeferV2() {
	utl.PLine("showDeferV2() world !!")

	utl.PLine("showDeferV2() hello ")
}

// Suppose we wanted to create a file, write to it, and then close when we’re done. Here’s how we could do that with defer.
func showFileDefer() {
	utl.PLine("showFileDefer()")

	f := createFile("./data/defer.txt")

	/*
		Immediately after getting a file object with createFile, we defer the closing of that file with closeFile.
		This will be executed at the end of the enclosing function (showFileDefer), after writeFile has finished.
	*/
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	utl.PLine("creating")

	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	utl.PLine("writing")
	fmt.Fprintln(f, time.Now(), " :: Writing to file using defer.")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	// It’s important to check for errors when closing a file, even in a deferred function.
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
