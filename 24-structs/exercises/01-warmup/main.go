// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus



// ---------------------------------------------------------
// EXERCISE: Warm Up
//
//  Starting with this exercise, you'll build a command-line
//  game store.
//
//  1. Declare the following structs:
//
//     + item: id (int), name (string), price (int)
//
//     + game: embed the item, genre (string)
//
//
//  2. Create a game slice using the following data:
//
//     id  name          price    genre
//
//     1   god of war    50       action adventure
//     2   x-com 2       30       strategy
//     3   minecraft     20       sandbox
//
//
//  3. Print all the games.
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------
package main

import (
	"fmt"
)

type item struct {
	id int
	name string
	price int
}

type game struct {
	item item
	genre string
}

func main() {
	games := []game{
		{
			item:  item{1, "god of war", 59},
			genre: "Action",
		},
		{
			item:  item{2, "x-com 2", 49},
			genre: "Adventure",
		},
		{
			item:  item{3, "minecraft", 69},
			genre: "RPG",
		},
	}

	fmt.Printf("Andy's game store has 3 games.\n")
	fmt.Printf("id  name          genre        price\n")
	for _, g := range games {
		fmt.Printf("#%-2d %-13s %-12s $%-8d\n", g.item.id, g.item.name, g.genre, g.item.price,)
	}
}
