package main

import (
	utl "autilities"
	"strings"
)

var header = utl.Header{}

/*
 */
func main() {
	/*
	 */
	header.DisplayHeader("Showing Exercise: Maps")

	inputString := "Hello, world! Hello, Go! Go Go!"
	wordCounts := WordCount(inputString)

	// Print the word frequencies.
	for word, count := range wordCounts {
		utl.PFmted("Word: %s | Count: %d\n", word, count)
	}
}

// WordCount returns a map of word frequencies in the input string.
func WordCount(s string) map[string]int {
	// Initialize an empty map to store word counts.
	wordMap := make(map[string]int)

	// Split the input string into words using strings.Fields.
	words := strings.Fields(s)

	// Count the occurrences of each word.
	for _, word := range words {
		wordMap[word]++
	}

	return wordMap
}
