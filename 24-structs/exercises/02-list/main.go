// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus


// ---------------------------------------------------------
// EXERCISE: List
//
//  Now, it's time to add an interface to your program using
//  the bufio.Scanner. So the users can list the games, or
//  search for the games by id.
//
//  1. Scan for the input in a loop (use bufio.Scanner)
//
//  2. Print the available commands.
//
//  3. Implement the quit command: Quits from the loop.
//
//  4. Implement the list command: Lists all the games.
//
//
// EXPECTED OUTPUT
//  Please run the solution and try the program with list and
//  quit commands.
// ---------------------------------------------------------

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
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

	fmt.Printf("Andy's game store has %d games.\n", len(games))
	// fmt.Printf("id  name          genre        price\n")


	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("\n> l   : lists all the games")
		fmt.Print("\n> q   : quits\n\n")
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := strings.ToLower(scanner.Text())
		if text == "l" {
			for _, g := range games {
				fmt.Printf("#%-2d %-13s %-12s $%-8d\n", g.item.id, g.item.name, g.genre, g.item.price,)
			}
		} else if text == "q" {
			fmt.Println("bye!")
			return
		}
	}
}
