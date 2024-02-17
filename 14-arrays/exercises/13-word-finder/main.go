// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus


// ---------------------------------------------------------
// EXERCISE: Word Finder
//
//   Your goal is to search for the words inside the corpus.
//
//   Note: This exercise is similar to the previous word finder program:
//   https://github.com/inancgumus/learngo/tree/master/13-loops/10-word-finder-labeled-switch
//
//   1. Get the search query from the command-line (it can be multiple words)
//
//   2. Filter these words, do not search for them:
//      and, or, was, the, since, very
//
//      To do this, use an array for the filtered words.
//
//   3. Print the words found.
//
// RESTRICTION
//   + The search and the filtering should be case insensitive
//
// HINT
//   + strings.Fields function converts a given string to a slice.
//
//     You can find its example in the word finder program that I've mentioned
//     above.
//
// EXPECTED OUTPUT
//   go run main.go
//     Please give me a word to search.
//
//   go run main.go and was
//
//   go run main.go AND WAS
//
//   go run main.go cat beginning
//     #2 : "cat"
//     #11: "beginning"
//
//   go run main.go Cat Beginning
//     #2 : "cat"
//     #11: "beginning"
// ---------------------------------------------------------



package main

import (
  "fmt"
  "os"
  "strings"
)

const (
  usage    = "Please give me a word to search."
  havebook    = "We have the book: #%d : %q\n"
	corpus   = "lazy cat jumps again and again and again since the beginning this was very important"
)

func main() {
  args := os.Args

	if len(args) < 3 {
    fmt.Println(usage)
    return
  }

	filtered := [...]string{
		"and", "or", "was", "the", "since", "very",
	}

	sentence := strings.Fields(corpus)

mark:
	for i := 0; i < len(args[1:]); i++ {
		for _, v := range filtered {
			if strings.Contains(strings.ToLower(v), strings.ToLower(args[i+1])) {
				continue mark
			}
		}

		for ind2, w := range sentence {
			if strings.Contains(strings.ToLower(w), strings.ToLower(args[i+1])) {
				fmt.Printf(havebook, ind2+1, w)
			}
		}
	}
}