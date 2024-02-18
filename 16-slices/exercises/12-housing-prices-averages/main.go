// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------

package main

import (
  "fmt"
  "strings"
  "strconv"
)

const (
  header = "Location,Size,Beds,Baths,Price"
  data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

  separator = ","
)

func main() {
  h := strings.Split(header, separator)
  input := strings.Split(data, "\n")

  for _, item := range h {
    fmt.Printf("%-15s", item)
  }
  fmt.Println("\n" + strings.Repeat("=", 75))

  var totalsize, totalbed, totalbath, totalprice float64

  for s := range(input) {
    input := strings.Split(input[s], separator)
    size, _ := strconv.ParseFloat(input[1], 64)
    totalsize += size
    bed, _ := strconv.ParseFloat(input[2], 64)
    totalbed += bed
    bath, _ := strconv.ParseFloat(input[3], 64)
    totalbath += bath
    price, _ := strconv.ParseFloat(input[4], 64)
    totalprice += price
    fmt.Printf("%-15s%-15s%-15s%-15s%-15s\n", input[0], input[1], input[2], input[3], input[4])
  }

  size_a := totalsize / float64(len(input))
  bed_a := totalbed / float64(len(input))
  bath_a := totalbath / float64(len(input))
  price_a := totalprice / float64(len(input))
  fmt.Println("\n" + strings.Repeat("=", 75))
  fmt.Printf("%-15.2s%-15.2f%-15.2f%-15.2f%-15.2f\n", "", size_a, bed_a, bath_a, price_a)
}
