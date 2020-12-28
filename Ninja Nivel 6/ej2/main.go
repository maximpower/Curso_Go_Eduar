package main

import "fmt"

func foo(nums ...int) int {
	total := 0

	for _, v := range nums {
		total += v
	}
	return total

}

func bar(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	r1 := foo(nums...)
	fmt.Println("El valor es ", r1)
	r2 := bar(nums)
	fmt.Println("El valor es ", r2)

}
