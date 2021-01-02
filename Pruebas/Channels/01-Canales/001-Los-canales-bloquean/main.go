package main

import "fmt"

func main() {

	//unbuffered channel (canal sin búffer)
	ch := make(chan int)

	ch <- 42

	// Necesita mínimo dos Goroutines, una que envia y otra que recibe. Aquí solo hay una
	//Eso equivale a deadlock (bloqueo)
	fmt.Println(<-ch)

}
