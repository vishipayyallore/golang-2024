package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	fmt.Println("Log Level Filter")
	fmt.Println("Usage: go run . -level=DEBUG")

	// Parse command line flags.
	level := flag.String("level", "CRITICAL", "log level to filter for")
	flag.Parse()

	f, err := os.Open("./logs/log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bufReader := bufio.NewReader(f)

	for line, err := bufReader.ReadString('\n'); err == nil; line, err = bufReader.ReadString('\n') {
		if strings.Contains(line, *level) {
			fmt.Print(line)
		}
	}
}
