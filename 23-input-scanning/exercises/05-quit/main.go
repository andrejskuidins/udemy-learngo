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
	// rx := regexp.MustCompile(os.Args[1])
	unique := make(map[string]bool)
	grep := os.Args[1]

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		if strings.Contains(line, grep) {
			fmt.Println(line)
		}
	}
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	words := 0
	unique := make(map[string]bool)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words++
		unique[strings.ToLower(scanner.Text())] = true
	}

	// fmt.Println(unique)
	fmt.Printf("There are %d words, %d of them are unique.\n", words, len(unique)) // Println will add back the final '\n'
}