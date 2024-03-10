package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	link := os.Args[1]
	sent := strings.Split(link, " ")
	buf := make([]string, 0)
	final := make([]string, 0)
	var place int


	for i, v := range sent {
		if len(v) > 6 {
			fmt.Printf(v[0:6])
			if v[0:6] == "http://" {
				buf := append(buf, v[0:6])
				for i := 0; i < len(v[6:]); i++ {
					buf = append(buf, "*")
				}
				place = i
			}
		}

		final = append(sent[:0], sent[0:place]...)
		final = append(final, buf...)
		final = append(final, sent[place+1:]...)
	}
	fmt.Printf("%q\n", final)
}