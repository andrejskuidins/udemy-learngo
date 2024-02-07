// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Set an Alarm
//
//  Goal is printing " ALARM! " every 10 seconds.
//
// EXPECTED OUTPUT
//
//     ██   ███       ███  ██        ███  ███
//      █     █   ░     █   █    ░     █  █ █
//      █   ███       ███   █        ███  ███
//      █   █     ░     █   █    ░     █    █
//     ███  ███       ███  ███       ███  ███
//
//          ███  █    ███  ██   █ █   █
//          █ █  █    █ █  █ █  ███   █
//          ███  █    ███  ██   █ █   █
//          █ █  █    █ █  █ █  █ █
//          █ █  ███  █ █  █ █  █ █   █
//
//     ██   ███       ███  ██        █ █  ██
//      █     █   ░     █   █    ░   █ █   █
//      █   ███       ███   █        ███   █
//      █   █     ░     █   █    ░     █   █
//     ███  ███       ███  ███         █  ███
//
// HINTS
//
//  <<< BEWARE THE SPOILER! >>>


// ---------------------------------------------------------

package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	screen.Clear()

	for {
		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		clock_a := [...]placeholder{
			digits[a], digits[l], digits[a], digits[r], digits[m], digits[exc],
		}

		for line := range clock[0] {
			for index, digit := range clock {
				// colon blink
				next := clock[index][line]
				if sec%10 == 0 {
					fmt.Print(clock_a)
				} else if digit == colon && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}

		time.Sleep(time.Second)
	}
}
