package main

import (
    "fmt"
    "os"
    "unicode/utf8"
)

func main() {
    text := os.Args[1]
    buf := make([]byte, 0, len(text))
    var runeCount int
    var prevRune rune
    for len(text) > 0 {
        r, size := utf8.DecodeRuneInString(text)
        text = text[size:]
        if runeCount > 40 && r == ' ' {
            buf = append(buf, '\n')
            runeCount = 0
        } else if runeCount > 40 && prevRune == ' ' {
            buf = append(buf, '\n')
            buf = append(buf, []byte(string(r))...)
            runeCount = 1
        } else {
            buf = append(buf, []byte(string(r))...)
            runeCount++
        }
        prevRune = r
    }
    fmt.Println(string(buf))
}