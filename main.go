package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Defining a boolean flag -l to count iines instead of words
	lines := flag.Bool("l", false, "Count lines")
	isCountBytes := flag.Bool("b", false, "Count bytes")
	fileName := flag.String("file", "", "Path to the file to count\n You can provide multiple filenames delimited by space to be parsed")
	// Parsing the flags provided by the user
	flag.Parse()
	var fileNames []string
	if *fileName != "" {
		fileNames = strings.Split(*fileName, " ")
	}
	// TODO this is a mess
	// try to clean this
	if len(fileNames) > 0 {
		for _, fName := range fileNames {
			file, err := os.Open(fName)
			if err != nil {
				fmt.Printf("error reading file %s", *fileName)
				os.Exit(1)
			}
			fmt.Println(count(file, *lines, *isCountBytes, fName))
			if err := file.Close(); err != nil {
				fmt.Println("Error trying to close the file", file.Name())
			}
		}
		os.Exit(1)
	} else {
		stdin, err := os.Stdin.Stat()
		if err != nil {
			fmt.Println("No file or input provided to be read")
			os.Exit(1)
		}
		if stdin.Size() > 0 {
			fmt.Println(count(os.Stdin, *lines, *isCountBytes, "stdin"))
			os.Exit(1)
		}
	}

	flag.Usage()
}

func count(r io.Reader, countLines bool, countBytes bool, fileName string) string {
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
	return fmt.Sprintf("Counted for file %v this many %v : %d", fileName, format, wc)
}
