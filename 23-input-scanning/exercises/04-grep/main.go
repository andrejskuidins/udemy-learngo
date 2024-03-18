// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus



// ---------------------------------------------------------
// EXERCISE: Grep Clone
//
//  Create a grep clone. grep is a command-line utility for
//  searching plain-text data for lines that match a specific
//  pattern.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Accept a command-line argument for the pattern
//
//  3. Only print the lines that contains that pattern
//
//  4. If no pattern is provided, print all the lines
//
//
// EXPECTED OUTPUT
//
//  go run main.go come < shakespeare.txt
//
//    come night come romeo come thou day in night
//    come gentle night come loving black-browed night
//
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
	grep := os.Args[1]

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		if strings.Contains(line, grep) {
			fmt.Println(line)
		}
	}

	// linesLoop:
	// for scanner.Scan() {
	// 	line := strings.ToLower(scanner.Text())
	// 	for _, v := range strings.Split(line, " ") {
	// 		if rx.MatchString(v) {
	// 			fmt.Println(line)
	// 			continue linesLoop
	// 		}
	// 	}
	// }

	// fmt.Println(unique)

}