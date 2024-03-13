package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    link := os.Args[1]
    buf := make([]byte, 0)
    splitted := strings.Split(link, " ")

    var parts []string

    for _, v := range splitted {
        if len(v) > 7 && v[:7] == "http://" {
            for range v[7:] {
                buf = append(buf, '*')
            }
            modifiedURL := "http://" + string(buf)
            parts = append(parts, modifiedURL)
        } else {
            parts = append(parts, v)
        }
    }

    fmt.Println(strings.Join(parts, " "))
}