package main

import "fmt"

func main() {
	salir := make(chan int)
	c := gen(salir)

	recibir(c, salir)

	fmt.Println("A punto de finalizar.")
}

func gen(salir chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		salir <- 1
		close(c)
	}()

	return c
}

func recibir(ch, salir <-chan int) {
	for {
		select {
		case v := <-ch:
			fmt.Println("ch:", v)
		case v := <-salir:
			fmt.Println("salir:", v)
			return
		}
	}

}
