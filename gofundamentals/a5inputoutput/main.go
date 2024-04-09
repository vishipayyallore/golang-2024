package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter you name: ")
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	fmt.Println("Welcome", s, ".")
}
