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
	res := count(b, false)
	if res != exp {
		t.Errorf("\nExpected %v \nAnd got %v", exp, res)
	}
}

// TestCountLines tests the count function set to count lines
func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")
	exp := fmt.Sprintf("Counted this many %v : %d", "lines", 3)
	res := count(b, true)
	if res != exp {
		t.Errorf("\nExpected: %v \nAnd %v", exp, res)
	}
}
