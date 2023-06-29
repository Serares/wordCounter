package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Defining a boolean flag -l to count iines instead of words
	lines := flag.Bool("l", false, "Count lines")
	// Parsing the flags provided by the user
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines))
}

func count(r io.Reader, countLines bool) string {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)
	// Define the scanner split type to words (default is split by lines)
	// use the flag
	// the default method of the split function is to scan by lines
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	// Defining a counter
	wc := 0 // For every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	format := "words"
	if countLines {
		format = "lines"
	}

	// Return the total
	return fmt.Sprintf("Counted this many %v : %d", format, wc)
}
