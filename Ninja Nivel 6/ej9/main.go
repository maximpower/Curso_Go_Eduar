package main

import "fmt"

func foo(sum func(z []int) int, enteros []int) int {
	x := sum(enteros)
	x++
	return x
}

func sum(enteros []int) int {
	if len(enteros) == 0 {
		return 0
	} else if len(enteros) == 1 {
		return enteros[0]
	}

	val1 := enteros[0]
	val2 := enteros[len(enteros)-1]

	return val1 + val2

}

func main() {

	enteros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	res := foo(sum, enteros)

	fmt.Println("El resultado final es:", res)

}
