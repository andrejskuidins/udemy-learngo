// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Odd or Even
//
//  1. Get a number from the command-line.
//
//  2. Find whether the number is odd, even and divisible by 8.
//
// RESTRICTION
//  Handle the error. If the number is not a valid number,
//  or it's not provided, let the user know.
//
// EXPECTED OUTPUT
//  go run main.go 16
//    16 is an even number and it's divisible by 8
//
//  go run main.go 4
//    4 is an even number
//
//  go run main.go 3
//    3 is an odd number
//
//  go run main.go
//    Pick a number
//
//  go run main.go ABC
//    "ABC" is not a number
// ---------------------------------------------------------

package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	usage = "Pick a number"
	wrong = "%q is not a number\n"
	even  = "%d is an even number\n"
	odd   = "%d is an odd number\n"
	by8   = "%d is an even number and it's divisible by 8\n"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println(usage)
		return
	}

	oddeven, err := strconv.Atoi(args[1])

	if err != nil {
		fmt.Printf(wrong, args[1])
		return
	} else if oddeven%8 == 0 {
		fmt.Printf(by8, oddeven)
	} else if oddeven%2 == 0 {
		fmt.Printf(even, oddeven)
	} else if oddeven%2 != 0 {
		fmt.Printf(odd, oddeven)
	}
}
