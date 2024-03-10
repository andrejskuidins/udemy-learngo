package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	link := os.Args[1]
	for _, v := range strings.Split(link, " ")  {
		if v[0:6] == "http://" {
			buf := v[0:6]
			for _, r := range buf[6:] {
				buf = append(buf, "*")
			}
		}

		fmt.Printf("%q\n", v)
	}
}