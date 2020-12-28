package main

import "fmt"

func main() {
	n1 := factorial(4)
	fmt.Println(n1)
	n2 := factorialCiclos(4)
	fmt.Println(n2)
}

// recursividad
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//Factorial con cilos
func factorialCiclos(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}

// 4 * factorial(3)
// 3! = 3* factorial(2)
// 2! = 2* factorial(1)
// 1! = 1* factorial(0)
// 0! = 1* factorial()
