package main

import "fmt"

func main() {

	var s [5]int
	s[0] = 5
	s[1] = 12
	s[2] = 28
	s[3] = 11
	s[4] = 6

	for i, v := range s {
		fmt.Printf("El valor en la posición %v, es: %v\n", i, v)
	}

	fmt.Printf("El tipo del arreglo es: %T", s)

}
