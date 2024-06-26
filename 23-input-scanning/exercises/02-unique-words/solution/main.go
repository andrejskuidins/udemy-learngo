// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	total, words := 0, make(map[string]int)
	for in.Scan() {
		total++
		words[in.Text()]++
	}
	fmt.Printf("There are %d words, %d of them are unique.\n",
		total, len(words))

	fmt.Println(words)
}
