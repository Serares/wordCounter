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
	isCountBytes := flag.Bool("b", false, "Count bytes")
	fileName := flag.String("file", "", "Path to the file to count")
	// Parsing the flags provided by the user
	flag.Parse()
	var readerContent io.Reader

	if *fileName != "" {
		file, err := os.Open(*fileName)
		if err != nil {
			fmt.Printf("error reading file %s", *fileName)
			os.Exit(1)
		}
		readerContent = file

	} else {
		readerContent = os.Stdin
	}
	fmt.Println(count(readerContent, *lines, *isCountBytes))
}

func count(r io.Reader, countLines bool, countBytes bool) string {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)
	// Define the scanner split type to words (default is split by lines)
	// use the flag
	// the default method of the split function is to scan by lines
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	// Defining a counter
	wc := 0 // For every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	// TODO this is not ideal
	// try to refactor?
	format := "words"
	if countLines {
		format = "lines"
	}

	if countBytes {
		format = "bytes"
	}

	// Return the total
	return fmt.Sprintf("Counted this many %v : %d", format, wc)
}
