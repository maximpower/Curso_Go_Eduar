package main

import "fmt"

func main() {

	//buffered channel (canal con búffer)
	//Se añade otro parámetro en el make, el número de valores en buffer
	ch := make(chan int, 1)

	ch <- 42
	// Al ser canal en búffer, no se necesita goroutine para recibir.
	fmt.Println(<-ch)

}
