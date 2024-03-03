package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Printf("File name %s \n", file.Name())
		os.Chdir(file.Name())
		files, err := os.ReadDir(".")
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
		os.Chdir("..")
	}
}