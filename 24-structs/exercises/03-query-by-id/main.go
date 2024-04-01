// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus


// ---------------------------------------------------------
// EXERCISE: Query By Id
//
//  Add a new command: "id". So the users can query the games
//  by id.
//
//  1. Before the loop, index the games by id (use a map).
//
//  2. Add the "id" command.
//     When a user types: id 2
//     It should print only the game with id: 2.
//
//  3. Handle the errors:
//
//     id
//     wrong id
//
//     id HEY
//     wrong id
//
//     id 10
//     sorry. I don't have the game
//
//     id 1
//     #1: "god of war" (action adventure) $50
//
//     id 2
//     #2: "x-com 2" (strategy) $40
//
//
// EXPECTED OUTPUT
//  Please also run the solution and try the program with
//  list, quit, and id commands to see it in action.
// ---------------------------------------------------------

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"regexp"
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

	// index the games by id
	byID := make(map[int]game)
	for _, g := range games {
		byID[g.item.id] = g
	}

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("\n> l    : lists all the games")
		fmt.Print("\n> id N : queries a game by id")
		fmt.Print("\n> q    : quits\n\n")
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := strings.ToLower(scanner.Text())
		switch text {
		case "l":
			for _, g := range games {
				fmt.Printf("#%-2d %-13s %-12s $%-8d\n", g.item.id, g.item.name, g.genre, g.item.price)
				continue
			}
		case "q":
			fmt.Println("bye!")
			return
		}
		pattern := `^id\s+(\d+)$`
		re := regexp.MustCompile(pattern)
		matches := re.MatchString(text)
		if matches {
			ind, _ := strconv.Atoi(strings.Split(text, " ")[1])
			if ind > len(games) {
				fmt.Println("sorry. I don't have the game")
			} else {
				g := byID[ind]
				fmt.Printf("#%-2d %-13s %-12s $%-8d\n", g.item.id, g.item.name, g.genre, g.item.price,)
			}
		}
	}
}
