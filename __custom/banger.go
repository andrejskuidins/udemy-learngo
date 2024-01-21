package main

import (
  "fmt"
  "os"
  "strings"
)

func addExclam(s string) string {
  var exclam string
  for range s {
    exclam += "!"
  }
  return exclam
}

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Please provide a command-line argument.")
    return
  }

  argWithProg := os.Args[1]
  argWithProgUpper := strings.ToUpper(argWithProg)

  exclamString := addExclam(argWithProg)
  fmt.Println(exclamString + argWithProgUpper + exclamString)
}
