package main

import "fmt"

func main() {

	//buffered channel (canal con búffer)
	//Se añade otro parámetro en el make, el número de valores en buffer
	ch := make(chan int, 1)

	ch <- 42
	//El tamaño del búffer es 1, pero ponemos 2 valores, deadlock!
	ch <- 43

	// Deadlock
	fmt.Println(<-ch)

}
