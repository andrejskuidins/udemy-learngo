// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus


// ---------------------------------------------------------
// EXERCISE: Decode
//
//  At the beginning of the file:
//
//  1. Load the initial data to the game store from json.
//     (see the data constant below)
//
//  2. Load the decoded values into the usual `game` values (to the games slice as well).
//
//     So the rest of the program can work intact.
//
//
// HINT
//
//  Move the jsonGame type to the top and reuse it both when
//  loading the initial data, and in the "save" command.
//
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"encoding/json"
)

const data = `
[
        {
                "id": 4,
                "name": "HOMM3",
                "genre": "turn-based strategy",
                "price": 50
        },
        {
                "id": 5,
                "name": "csgo",
                "genre": "shooter",
                "price": 40
        },
        {
                "id": 6,
                "name": "startcraft",
                "genre": "strategy",
                "price": 20
        }
]`

type item struct {
	Id int
	Name string
	Price int
}

type game struct {
	Item item
	Genre string
}

type jsonGame struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Price int    `json:"price"`
}

func main() {
	var encodable []jsonGame

	games := []game{
		{
			Item:  item{1, "god of war", 59},
			Genre: "Action",
		},
		{
			Item:  item{2, "x-com 2", 49},
			Genre: "Adventure",
		},
		{
			Item:  item{3, "minecraft", 69},
			Genre: "RPG",
		},
	}

	fmt.Printf("Andy's game store has %d games.\n", len(games))
	// fmt.Printf("id  name          genre        price\n")

	// index the games by id
	byID := make(map[int]game)
	for _, g := range games {
		byID[g.Item.Id] = g
	}

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("\n> l    : lists all the games")
		fmt.Print("\n> id N : queries a game by id")
		fmt.Print("\n> s    : exports the data to json and quits")
		fmt.Print("\n> ld   : loads the data into slice and quits")
		fmt.Print("\n> q    : quits\n\n")
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := strings.Fields(scanner.Text())
		switch text[0] {
		case "l":
			for _, g := range games {
				fmt.Printf("#%-2d %-13s %-12s $%-8d\n", g.Item.Id, g.Item.Name, g.Genre, g.Item.Price)
				continue
			}
		case "q":
			fmt.Println("bye!")
			return
		case "id":
			if len(text) < 2 {
				fmt.Println("wrong id")
				continue
			}
			ind, err := strconv.Atoi(text[1])
			if err != nil {
				fmt.Println("wrong id")
			} else if ind > len(games) {
				fmt.Println("sorry. I don't have the game")
			} else {
				g := byID[ind]
				fmt.Printf("#%-2d %-13s %-12s $%-8d\n", g.Item.Id, g.Item.Name, g.Genre, g.Item.Price)
			}
		case "s":
			for _, g := range games {
				encodable = append(encodable,
					jsonGame{g.Item.Id, g.Item.Name, g.Genre, g.Item.Price})
			}

			json_out, err := json.MarshalIndent(encodable, "",  "\t")
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(string(json_out))
		case "ld":

		}
	}
}