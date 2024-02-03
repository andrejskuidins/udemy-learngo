package main
import (
    "fmt"
    "math/rand"
    "time"
		"os"
		"strconv"
)

const (
	usage = "[your name] [positive|negative]"
	mood = "%s feels %s\n"
)

func main()  {
	if len(os.Args) != 2 {
		fmt.Println("Give me millis")
		return
	}

	d, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("No millis provided. x10")
		return
	}

	millis := time.Duration(d)
	s := "\\/|-"

	for {
		fmt.Printf("\r%s Please Wait. Processing....\n", string(s[rand.Intn(4)]))
		time.Sleep(millis * time.Millisecond)
	}
}