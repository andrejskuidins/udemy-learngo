package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    link := os.Args[1]
    buf := make([]byte, 0)

    for _, v := range strings.Split(link, " ") {

        if len(v) > 7 && v[:7] == "http://" {
            fmt.Println(v)
            for range v[7:] {
                buf = append(buf, '*')
            }
        }
    }

		modifiedURL := "http://" + string(buf)[:]
		fmt.Println(modifiedURL)
}

// go run main.go "Here: http://www.mylink.com Click!"