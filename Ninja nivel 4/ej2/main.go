package main

import "fmt"

func main() {

	s := []int{5, 12, 28, 11, 6, 12, 7, 8, 9, 10}

	for i, v := range s {
		fmt.Printf("El valor en la posición %v, es: %v\n", i, v)
	}

	fmt.Printf("El tipo del arreglo es: %T", s)

}
