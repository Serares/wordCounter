package main

import (
	"bytes"
	"fmt"
	"testing"
)

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := fmt.Sprintf("Counted this many %v : %d", "words", 4)
	res := count(b, false, false)
	if res != exp {
		t.Errorf("\nExpected %v \nAnd got %v", exp, res)
	}
}

// TestCountLines tests the count function set to count lines
func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")
	exp := fmt.Sprintf("Counted this many %v : %d", "lines", 3)
	res := count(b, true, false)
	if res != exp {
		t.Errorf("\nExpected: %v \nAnd %v", exp, res)
	}
}

func TestCountBytes(t *testing.T) {
	stringCounted := "how many bytes should be in here?"
	b := bytes.NewBufferString(stringCounted)
	exp := fmt.Sprintf("Counted this many %v : %d", "bytes", len(stringCounted))
	res := count(b, false, true)

	if res != exp {
		t.Errorf("\nExpected: %v \nAnd %v", exp, res)
	}
}
