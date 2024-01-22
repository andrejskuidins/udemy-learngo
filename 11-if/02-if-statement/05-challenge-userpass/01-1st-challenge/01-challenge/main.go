// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

package main
import (
  "fmt"
	"os"
)

func main() {
	username, password := "jack", "1888"

	l := os.Args

	if len(l) < 3 {
    fmt.Println("Usage: [username] [password]")
  }	else if l[1] != username {
		fmt.Printf("Access denied for %q.\n", l[1])
	} else if l[2] != password {
		fmt.Printf("Invalid password for %q.\n", l[1])
	} else {
		fmt.Printf("Access granted to %q.\n", l[1])
	}
}
