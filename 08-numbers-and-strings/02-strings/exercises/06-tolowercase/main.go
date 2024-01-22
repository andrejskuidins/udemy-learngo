// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: ToLowercase
//
//  1. Look at the documentation of strings package
//  2. Find a function that changes the letters into lowercase
//  3. Get a value from the command-line
//  4. Print the given value in lowercase letters
//
// HINT
//  Check out the strings package from Go online documentation.
//  You will find the lowercase function there.
//
// INPUT
//  "SHEPARD"
//
// EXPECTED OUTPUT
//  shepard
// ---------------------------------------------------------

package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Please provide a command-line argument.")
    return
  }

  argWithProg := os.Args[1]
  argWithProgLower := strings.ToLower(argWithProg)

  fmt.Println(argWithProgLower)
}

