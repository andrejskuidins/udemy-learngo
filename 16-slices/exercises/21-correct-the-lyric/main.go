// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus


// ---------------------------------------------------------
// EXERCISE: Correct the lyric
//
//  You have a slice that contains the words of Beatles'
//  legendary song: Yesterday. However, the order of the
//  words are incorrect.
//
// CURRENT OUTPUT
//
//  [all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay]
//
// EXPECTED OUTPUT
//
//  [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday]
//
//
// STEPS
//
//  INITIAL SLICE:
//    [all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay]
//
//
//  1. Prepend "yesterday" to the `lyric` slice.
//
//     RESULT SHOULD BE:
//     [yesterday all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay]
//
//
//  2. Put the words to the correct positions in the `lyric` slice.
//
//     RESULT SHOULD BE:
//     [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday]
//
//
//  3. Print the `lyric` slice.
//
//
// BONUS
//
//   + Think about when does the append allocate a new backing array.
//
//   + Check whether your conclusions are correct.
//
//
// HINTS
//
//   If you get stuck, check out the hints.md file.
//
// ---------------------------------------------------------

package main

import (
	"fmt"
	"strings"
)

func findIndex (arr []string, targetString string) int {
	for i, v := range(arr) {
		if v == targetString {
			return i
		}
	}
	return -1
}

func main() {
	// DON'T TOUCH THIS:
	lyric := strings.Fields(`all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay`)

	// ADD YOUR CODE BELOW:
	l := append([]string{"yesterday"}, lyric...)
	start := findIndex(lyric, "now")
	stop := findIndex(lyric, "stay")
	insertAt := findIndex(lyric, "away")
	fmt.Printf("%s\n", lyric)
	final := append(l[0:insertAt+2], append(l[start+1:stop+2],l[insertAt+2:start+1]...)...)
	fmt.Printf("%s\n", final)
}
