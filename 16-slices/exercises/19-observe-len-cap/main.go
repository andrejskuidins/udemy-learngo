// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Observe the length and capacity
//
//  Follow the instructions inside the code below to
//  gain more intuition about the length and capacity of a slice.
//
// ---------------------------------------------------------

func main() {
	// --- #1 ---
	// 1. create a new slice named: games
	//
	// 2. print the length and capacity of the games slice
	//
	// 3. comment out the games slice
	//    then declare it as an empty slice
	//
	// 4. print the length and capacity of the games slice
	//
	// 5. append the elements: "pacman", "mario", "tetris", "doom"
	//
	// 6. print the length and capacity of the games slice
	//
	// 7. comment out everything
	//
	// 8. declare the games slice again using a slice literal
	//    (use the same elements from step 5)
    // Creating a slice
    // games := []int{1, 2, 3, 4, 5}
		games := []string{}

		games = append(games, "pacman", "mario", "tetris", "doom")
    // Printing length and capacity
    fmt.Println("Length of slice s:", len(games))
    fmt.Println("Capacity of slice s:", cap(games))
	// --- #2 ---
	// 1. use a loop from 0 to 4 to slice the games slice, element by element.
	//
	// 2. print its length and capacity along the way (in the loop).

	fmt.Println()
	for i := 0; i < len(games); i++ {
    fmt.Printf("games[:%d]'s len: %d cap: %d\n", i, len(games), cap(games))
	}

	// --- #3 ---
	// 1. slice the games slice up to zero element
	//    (save the result to a new slice named: "zero")
	//
	// 2. print the games and the new slice's len and cap
	//
	// 3. append a new element to the new slice
	//
	// 4. print the new slice's lens and caps
	//
	// 5. repeat the last two steps 5 times (use a loop)
	//
	// 6. notice the growth of the capacity after the 5th append
	//
	// Use this slice's elements to append to the new slice:
	new := []string{"ultima", "dagger", "pong", "coldspot", "zetra"}
	fmt.Println()

	zero := games[:0]
	fmt.Printf("games's     len: %d cap: %d\n", len(games), cap(games))
	fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))

	zero = append(zero, new[0])
	fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))
	for i := 0; i < len(games); i++ {
	  zero = append(zero, new[i])
	  fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))
	}

	// --- #4 ---
	// using a range loop, slice the zero slice element by element,
	// and print its length and capacity along the way.
	//
	// observe that, the range loop only loops for the length, not the cap.
	fmt.Println()

	for i := range zero {
	  s := zero[:i]
	  fmt.Printf("zero[:%d]'s  len: %d cap: %d\n", i, len(s), cap(s))
	}

	// --- #5 ---
	// 1. do the 3rd step above again but this time, start by slicing
	//    the zero slice up to its capacity (use the cap function).
	//
	// 2. print the elements of the zero slice in the loop.
	fmt.Println()

	for i := range zero[:cap(zero)] {
		s := zero[:i]
	  fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", i, len(s), cap(s), s)
	}

	// --- #6 ---
	// 1. change the one of the elements of the zero slice
	//
	// 2. change the same element of the games slice
	//
	// 3. print the games and the zero slices
	//
	// 4. observe that they don't have the same backing array
	fmt.Println()
	zero[0] = "dream"
	games[0] = "light"
	fmt.Printf("zero  : %q\n", zero)
	fmt.Printf("games : %q\n", games)

	// --- #7 ---
	// try to slice the games slice beyond its capacity
	_ = games[:cap(games)+1]
}
