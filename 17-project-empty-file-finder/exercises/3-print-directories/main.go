// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus


// ---------------------------------------------------------
// EXERCISE: Find and write the names of subdirectories to a file
//
//  Create a program that can get multiple directory paths from
//  the command-line, and prints only their subdirectories into a
//  file named: dirs.txt
//
//
//  1. Get the directory paths from command-line
//
//  2. Append the names of subdirectories inside each directory
//     to a byte slice
//
//  3. Write that byte slice to dirs.txt file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Please provide directory paths
//
//   go run main.go dir/ dir2/
//
//   cat dirs.txt
//
//     dir/
//             subdir1/
//             subdir2/
//
//     dir2/
//             subdir1/
//             subdir2/
//             subdir3/
//

package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	items := os.Args[1:]
	if len(items) == 0 {
		fmt.Println("Please provide directory paths")
		return
	}

	sort.Strings(items)

	var data []byte
	for _, s := range items {
		data = append(data, s...)
		data = append(data, '\n')
		subdirs, err := os.ReadDir(s)
		if err != nil {
			fmt.Println(err)
		}
		for _, file := range subdirs {
			data = append(data, ' ')
			data = append(data, file.Name()...)
			data = append(data, '\n')
		}

	}

	err := os.WriteFile("dirs.txt", data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
