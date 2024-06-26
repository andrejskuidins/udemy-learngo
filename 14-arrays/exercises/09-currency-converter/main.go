// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus


// ---------------------------------------------------------
// EXERCISE: Currency Converter
//
//   In this exercise, you're going to display currency exchange ratios
//   against USD.
//
//   1. Declare a few constants with iota. They're going to be the keys
//      of the array.
//
//   2. Create an array that contains the conversion ratios.
//
//      You should use keyed elements and the contants you've declared before.
//
//   3. Get the USD amount to be converted from the command line.
//
//   4. Handle the error cases for missing or invalid input.
//
//   5. Print the exchange ratios.
//
// EXPECTED OUTPUT
//   go run main.go
//     Please provide the amount to be converted.
//
//   go run main.go invalid
//     Invalid amount. It should be a number.
//
//   go run main.go 10.5
//     10.50 USD is 9.24 EUR
//     10.50 USD is 8.19 GBP
//     10.50 USD is 1186.71 JPY
//
//   go run main.go 1
//     1.00 USD is 0.88 EUR
//     1.00 USD is 0.78 GBP
//     1.00 USD is 113.02 JPY
// ---------------------------------------------------------

package main

import (
  "fmt"
  "os"
  "strconv"
)

const (
  usage    = "Please provide the amount to be converted."
  wrong    = "Invalid amount: %q. It should be a number.\n"
	JPY = iota
	EUR
	RUB
)

func main() {
  args := os.Args

	if len(args) != 2 {
    fmt.Println(usage)
    return
  }

	rate, err := strconv.ParseFloat(args[1], 64)

	if err != nil {
		fmt.Printf(wrong, args[1])
    return
	}

	rates := [...]float64{
		JPY: 150,
		EUR: 0.9,
		RUB: 90,
	}

	for i := 2; i < len(rates); i++ {
		fmt.Printf("%.2f USD is %.2f\n", rate, float64(rate)*rates[i])
	}
}
