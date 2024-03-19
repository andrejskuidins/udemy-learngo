// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus


// ---------------------------------------------------------
// EXERCISE: Quit
//
//  Create a program that quits when a user types the
//  same word twice.
//
//
// RESTRICTION
//
//  The program should work case insensitive.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//
//   hey
//   HEY
//   TWICE!
//
//  go run main.go
//
//   hey
//   hi
//   HEY
//   TWICE!
// ---------------------------------------------------------

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "regexp"
)

func main() {
	// To create dynamic map
	vocab := make(map[string]bool)
	scanner := bufio.NewScanner(os.Stdin)
	for {
			fmt.Print("Enter Text: ")
			// Scans a line from Stdin(Console)
			scanner.Scan()
			// Holds the string that scanned
			text := scanner.Text()
			if !vocab[strings.ToLower(text)] {
				vocab[strings.ToLower(text)] = true
			} else {
					break
			}

	}
	// Use collected inputs
	fmt.Println(vocab)
}
