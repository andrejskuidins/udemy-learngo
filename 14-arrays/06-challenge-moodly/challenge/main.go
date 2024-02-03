// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Moodly
//
//   1. Get username from command-line
//
//   2. Display the usage if the username is missing
//
//   3. Create an array
//     1. Add three positive mood messages
//     2. Add three negative mood messages
//
//   4. Randomly select and print one of the mood messages
//
// EXPECTED OUTPUT
//
//   go run main.go
//     [your name]
//
//   go run main.go Socrates
//     Socrates feels good 👍
//
//   go run main.go Socrates
//     Socrates feels bad 👎
//
//   go run main.go Socrates
//     Socrates feels sad 😞
//
//   go run main.go Socrates
//     Socrates feels happy 😀
//
//   go run main.go Socrates
//     Socrates feels awesome 😎
//
//   go run main.go Socrates
//     Socrates feels terrible 😩
// ---------------------------------------------------------

package main
import (
    "fmt"
    "math/rand"
		"os"
)

const (
	usage = "[your name]"
	mood = "%s feels %s %s\n"
)

func main()  {
	args := os.Args
	if len(args) != 2 {
		fmt.Println(usage)
		return
	}

	moods := [6]string{"good", "happy", "awesome", "bad", "sad", "terrible"}
	smileyfaces := [6]string{"👍", "😀", "😎", "👎", "😞", "😩"}
	smileys := [6]string{"\U0001F44D", "\U0001F601", "\U0001F60E", "\U0001F44E", "\U0001F61E", "\U0001F616"}

	randomMood := rand.Intn(4)
	fmt.Printf(mood, args[1], moods[randomMood], smileys[randomMood])
	fmt.Printf(mood, args[1], moods[randomMood], smileyfaces[randomMood])
}