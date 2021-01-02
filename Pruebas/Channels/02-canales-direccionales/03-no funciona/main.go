package main

import "fmt"

func main() {

	//buffered channel (canal con búffer)
	//Se añade otro parámetro en el make, el número de valores en buffer
	// <-chan es solo receive only.
	ch := make(<-chan int, 2)

	ch <- 42
	ch <- 43

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("--------------------")

	fmt.Printf("%T\n", ch)

}
