// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus


// ---------------------------------------------------------
// EXERCISE: Encode
//
//  Add a new command: "save". Encode the games to json, and
//  print it, then terminate the loop.
//
//  1. Create a new struct type with exported fields: ID, Name, Genre and Price.
//
//  2. Create a new slice using the new struct type.
//
//  3. Save the games into the new slice.
//
//  4. Encode the new slice.
//
//
// RESTRICTION
//  Do not export the fields of the game struct.
//
//
// EXPECTED OUTPUT
//  Inanc's game store has 3 games.
//
//    > list   : lists all the games
//    > id N   : queries a game by id
//    > save   : exports the data to json and quits
//    > quit   : quits
//
//  save
//
//  [
//          {
//                  "id": 1,
//                  "name": "god of war",
//                  "genre": "action adventure",
//                  "price": 50
//          },
//          {
//                  "id": 2,
//                  "name": "x-com 2",
//                  "genre": "strategy",
//                  "price": 40
//          },
//          {
//                  "id": 3,
//                  "name": "minecraft",
//                  "genre": "sandbox",
//                  "price": 20
//          }
//  ]
//
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