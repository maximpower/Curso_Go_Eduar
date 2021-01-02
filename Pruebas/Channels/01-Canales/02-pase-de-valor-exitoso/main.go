package main

import (
	"fmt"
)

func main() {

	//unbuffered channel (canal sin búffer)
	ch := make(chan int)

	// Necesita mínimo dos Goroutines, una que envia y otra que recibe.
	//La siguiente función envía por canal el 42 y la goroutine main recibe y printa el valor.
	go func() {
		ch <- 42
	}()

	fmt.Println(<-ch)

}
