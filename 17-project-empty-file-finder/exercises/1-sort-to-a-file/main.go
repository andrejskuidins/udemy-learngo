// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus


// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file
//
//  1. Get arguments from command-line
//
//  2. Sort them
//
//  3. Write the sorted slice to a file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Send me some items and I will sort them
//
//   go run main.go orange banana apple
//
//   cat sorted.txt
//     apple
//     banana
//     orange
//
//
// HINTS
//
//   + REMEMBER: os.Args is a []string
//
//   + String slices are sortable using `sort.Strings`
//
//   + Use ioutil.WriteFile to write to a file.
//
//   + But you need to convert []string to []byte to be able to
//     write it to a file using the ioutil.WriteFile.
//
//   + To do that, create a new []byte and append the elements of your
//     []string.
//
// ---------------------------------------------------------



package main

import (
	"fmt"
	"sort"
	"os"
)

var pl = fmt.Println

const (
	wrong = "Send me some items and I will sort them"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		pl(wrong)
	}
	sort.Strings(args[:])
	var file []byte
	for _, v := range args {
		file = append(file, v...)
		file = append(file, '\n')
	}

	err := os.WriteFile("sorted.txt", file, 0666)
	if err != nil {
		pl("cannot write a file")
	}

	data, err := os.ReadFile("sorted.txt")
	if err != nil {
		pl("cannot read a file")
	}
	os.Stdout.Write(data)
}
