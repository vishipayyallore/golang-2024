package main

import (
	utl "autilities"
	"io"
	"os/exec"
)

var header = utl.Header{}

/*
Sometimes our Go programs need to spawn other, non-Go processes.
For example, we might need to run a shell command or execute a system binary. In this recipe, we will learn how to spawn processes in Go.
*/
func main() {
	/*
		Spawning Processes. EXECUTE THIS PROGRAM IN THE BASH TERMINAL
	*/
	header.DisplayHeader("Showing Spawning Processes")

	// We’ll start with a simple command that takes no arguments or input and just prints something to stdout. The exec.Command helper creates an object to represent this external process.
	dateCmd := exec.Command("date")

	// The Output method runs the command, waits for it to finish and collects its standard output. If there were no errors, dateOut will hold bytes with the date info.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	utl.PLine("> date")
	utl.PLine(string(dateOut))

	// Output and other methods of Command will return *exec.Error if there was a problem executing the command (e.g. wrong path), and *exec.ExitError if the command ran but exited with a non-zero return code.
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			utl.PLine("failed executing:", err)
		case *exec.ExitError:
			utl.PLine("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	// Next we’ll look at a slightly more involved case where we pipe data to the external process on its stdin and collect the results from its stdout.
	grepCmd := exec.Command("grep", "hello")

	// Here we explicitly grab input/output pipes, start the process, write some input to it, read the resulting output, and finally wait for the process to exit.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// We omitted error checks in the above example, but you could use the usual if err != nil pattern for all of them. We also only collect the StdoutPipe results, but you could collect the StderrPipe in exactly the same way.
	utl.PLine("> grep hello")
	utl.PLine(string(grepBytes))

	// Note that when spawning commands we need to provide an explicitly delineated command and argument array, vs. being able to just pass in one command-line string. If you want to spawn a full command with a string, you can use bash’s -c option:
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	utl.PLine("> ls -a -l -h")
	utl.PLine(string(lsOut))
}
