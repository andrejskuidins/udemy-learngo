
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

		// clock_a := [...]placeholder{
		// 	digits[a], digits[l], digits[a], digits[r], digits[m], digits[exc],
		// }
		for line := range clock[0] {
			fmt.Println(line)
		}
		fmt.Println(hour)
		fmt.Println(hour/10)
		fmt.Println(min)
		fmt.Println(min/10)
		fmt.Println(sec)
		fmt.Println(sec/10)
		time.Sleep(time.Second)
	}
}
