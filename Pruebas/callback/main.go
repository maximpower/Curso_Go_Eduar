package main

import "fmt"

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// s := sum(ii...)
	// fmt.Println(s)

	sumaPares := sumEven(sum, ii...)
	fmt.Println(sumaPares)

	sumaImpares := sumOdd(sum, ii...)
	fmt.Println(sumaImpares)

}

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func sumEven(f func(x ...int) int, y ...int) int {
	pares := []int{}

	for _, v := range y {
		if v%2 == 0 {
			pares = append(pares, v)
		}
	}

	return f(pares...)
}

func sumOdd(f func(x ...int) int, y ...int) int {
	impares := []int{}

	for _, v := range y {
		if v%2 != 0 {
			impares = append(impares, v)
		}
	}
	return f(impares...)

}
