package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("An error ocurred creating log.txt")
	}
	defer f.Close()

	log.SetOutput(f)

	f2, err := os.Open("sin-archivo.txt")
	if err != nil {
		log.Println(err)
	}
	defer f2.Close()

	fmt.Println("An error has ocurred. Look at the log.txt file")

}
