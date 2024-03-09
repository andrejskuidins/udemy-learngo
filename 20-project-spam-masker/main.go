package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	link := os.Args[1]
	for i, v := range strings.Split(link,  {
		if v == rune('h') {
			buf := link[i:i+7]
			if buf == "http://"
		}

		fmt.Printf("%q\n", v)
	}
}