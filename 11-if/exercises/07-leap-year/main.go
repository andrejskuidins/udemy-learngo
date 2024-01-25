// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Leap Year
//
//  Find out whether a given year is a leap year or not.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a year number
//
//  go run main.go eighties
//    "eighties" is not a valid year.
//
//  go run main.go 2018
//    2018 is not a leap year.
//
//  go run main.go 2100
//    2100 is not a leap year.
//
//  go run main.go 2019
//    2019 is not a leap year.
//
//  go run main.go 2020
//    2020 is a leap year.
//
//  go run main.go 2024
//    2024 is a leap year.
// ---------------------------------------------------------

package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	usage = "Give me a year number"
	wrong = "%q is not a valid year.\n"
	nleap = "%d is not a leap year.\n"
	leap  = "%d is a leap year.\n"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println(usage)
		return
	}

	isleap, err := strconv.Atoi(args[1])

	if err != nil {
		fmt.Printf(wrong, args[1])
		return
	} else if isleap%400 == 0 {
		fmt.Printf(leap, isleap)
	} else if isleap%4 == 0 && isleap%100 != 0 {
		fmt.Printf(leap, isleap)
	} else {
		fmt.Printf(nleap, isleap)
	}
}
