// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Slicing the Housing Prices
//
//  We have received housing prices. Your task is to print only the requested
//  columns of data (see the expected output).
//
//  1. Separate the data by the newline ("\n") -> rows
//
//  2. Separate each row of the data by the separator (",") -> columns
//
//  3. Find the from and to positions inside the columns depending
//     on the command-line arguments.
//
//  4. Print only the requested column headers and their data
//
//
// RESTRICTIONS
//
//  + You should use slicing when printing the columns.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Location       Size           Beds           Baths          Price
//
//    New York       100            2              1              100000
//    New York       150            3              2              200000
//    Paris          200            4              3              400000
//    Istanbul       500            10             5              1000000
//
//
//  go run main.go Location
//    Location       Size           Beds           Baths          Price
//
//    New York       100            2              1              100000
//    New York       150            3              2              200000
//    Paris          200            4              3              400000
//    Istanbul       500            10             5              1000000
//
//
//  NOTE : Finds the position of the Size column and slices the columns
//         appropriately.
//
//  go run main.go Size
//    Size           Beds           Baths          Price
//
//    100            2              1              100000
//    150            3              2              200000
//    200            4              3              400000
//    500            10             5              1000000
//
//
//  NOTE : Finds the positions of the Size and Baths columns and
//         slices the columns appropriately.
//
//  go run main.go Size Baths
//    Size           Beds           Baths
//
//    100            2              1
//    150            3              2
//    200            4              3
//    500            10             5
//
//
//  go run main.go Beds Price
//    Beds           Baths          Price
//
//    2              1              100000
//    3              2              200000
//    4              3              400000
//    10             5              1000000
//
//
//  Note : It works even if the given column name does not exist.
//
//  go run main.go Beds NotExist
//    Beds           Baths          Price
//
//    2              1              100000
//    3              2              200000
//    4              3              400000
//    10             5              1000000
//
//
//  go run main.go NotExist NotExist
//    Location       Size           Beds           Baths          Price
//
//    New York       100            2              1              100000
//    New York       150            3              2              200000
//    Paris          200            4              3              400000
//    Istanbul       500            10             5              1000000
//
//
// Note : It works even if the Price's index > Size's index
//
//        In that case, it resets the invalid starting position to 0.
//
//        But it still uses the Size column.
//
//  go run main.go Price Size
//    Location       Size
//
//    New York       100
//    New York       150
//    Paris          200
//    Istanbul       500
//
//
// HINTS
//
//  + strings.Split function can separate a string into a []string
//
// ---------------------------------------------------------

import (
  "fmt"
  "strings"
	"os"
  // "strconv"
)

const (
  header = "Location,Size,Beds,Baths,Price"
  data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

  separator = ","
)

func findIndex (arr []string, targetString string) int {
	for i, v := range(arr) {
		if v == targetString {
			return i
		}
	}
	return len(arr)-1
}

func main() {
	var locations [][5]string

	var col1 int
	var col2 int

	args := os.Args
  h := strings.Split(header, separator)
  input := strings.Split(data, "\n")

	switch len(args) {
	case 2:
		col1 = findIndex(h, args[1])
		col2 = len(h)
	case 3:
		col1 = findIndex(h, args[1])
		col2 = findIndex(h, args[2])+1
	default:
		col1 = 0
		col2 = 5
	}

	if col1 > col2 {
		col1 = 0
	}

  for _, item := range h[col1:col2] {
    fmt.Printf("%-15s", item)
  }
  fmt.Println("\n")

  for s := range(input) {
    input := strings.Split(input[s], separator)
		locationData := [5]string{input[0], input[1], input[2], input[3], input[4]}
		locations = append(locations, locationData)

    // fmt.Printf("%-15s%-15s%-15s%-15s%-15s\n", input[0], input[1], input[2], input[3], input[4])
  }
	for _, v := range locations {
		fmt.Println(v[col1:col2])
	}
}