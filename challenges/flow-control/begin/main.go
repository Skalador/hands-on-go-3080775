// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	// use os package to read the file specified as a command line argument
	args := os.Args[1]
	bs, err := os.ReadFile(args)
	if err != nil {
		panic(fmt.Errorf("failed to read file: %s", err))
	}

	// convert the bytes to a string
	stringData := string(bs)

	// initialize a map to store the counts
	var counts map[string]int

	// use the standard library utility package that can help us split the string into words
	words := strings.Fields(stringData)

	// capture the length of the words slice
	counts = map[string]int{
		"letters": 0,
		"symbols": 0,
		"numbers": 0,
		"words":   len(words),
	}

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	for _, word := range words {
		for _, character := range word {
			if unicode.IsLetter(character) {
				counts["letters"] += 1
			} else if unicode.IsNumber(character) {
				counts["numbers"] += 1
			} else {
				counts["symbols"] += 1
			}
		}
	}

	// dump the map to the console using the spew package
	spew.Dump(counts)
}
