package main

import "fmt"

func foo() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}

func main() {

	a := foo()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

}
