package main
import (
    "fmt"
    // "time"
		"strconv"
		"os"
)

func main()  {
	args := os.Args

	if len(os.Args) != 3 {
		fmt.Println("Give me size for the board")
		return
	}

	rows, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("No rows provided.")
		return
	}
	cols, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("No cols provided.")
		return
	}

	slice := make([][]bool, rows, cols)


	fmt.Print("\033[H\033[2J")

	// Initializing each row with false values
	for i := range slice {
			slice[i] = make([]bool, cols)
	}

	// time.Sleep(time.Second)
	fmt.Println(slice)
	fmt.Println(slice2)

}