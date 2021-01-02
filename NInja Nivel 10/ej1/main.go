package main

import "fmt"

func main() {
	c := make(chan int)
	c2 := make(chan int, 1)
	//Primera parte del ej.
	go func() {
		c <- 42
	}()
	c2 <- 24

	fmt.Println("Valor del canal 1 (Sin buffer, con goroutine)", <-c)
	fmt.Println("Valor del canal 2 (Con búffer):", <-c2)
}
