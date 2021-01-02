package main

import "fmt"

func main() {

	ch := make(chan []int, 1)

	go func() {
		ch <- []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	}()

	fmt.Println("Channel try: ", <-ch)

}
