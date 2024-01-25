// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus


// ---------------------------------------------------------
// STORY
//
//  Your boss wants you to create a program that will check
//  whether a person can watch a particular movie or not.
//
//  Imagine that another program provides the age to your
//  program. Depending on what you return, the other program
//  will issue the tickets to the person automatically.
//
// EXERCISE: Movie Ratings
//
//  1. Get the age from the command-line.
//
//  2. Return one of the following messages if the age is:
//     -> Above 17         : "R-Rated"
//     -> Between 13 and 17: "PG-13"
//     -> Below 13         : "PG-Rated"
//
// RESTRICTIONS:
//  1. If age data is wrong or absent let the user know.
//  2. Do not accept negative age.
//
// BONUS:
//  1. BONUS: Use if statements only twice throughout your program.
//  2. BONUS: Use an if statement only once.
//
// EXPECTED OUTPUT
//  go run main.go 18
//    R-Rated
//
//  go run main.go 17
//    PG-13
//
//  go run main.go 12
//    PG-Rated
//
//  go run main.go
//    Requires age
//
//  go run main.go -5
//    Wrong age: "-5"
// ---------------------------------------------------------

package main

import (
  "fmt"
  "os"
  "strconv"
)

const (
  usage    = "Requires age"
  wrong    = "Wrong age: %q\n"
  pg       = "PG-Rated"
  pg13     = "PG-13"
  rr       = "R-Rated"
)

func main() {
  args := os.Args

  if len(args) != 2 {
    fmt.Println(usage)
    return
  }

  rating, err := strconv.Atoi(args[1])

  if err != nil || rating < 0{
    fmt.Printf(wrong, args[1])
  } else if rating < 13 {
    fmt.Println(pg)
  } else if rating >= 13 && rating <= 17 {
    fmt.Println(pg13)
  } else if rating > 17 {
    fmt.Println(rr)
  }
}
