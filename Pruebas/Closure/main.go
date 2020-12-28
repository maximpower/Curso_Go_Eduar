package main

import "fmt"

func main() {
	var x int
	fmt.Println(x)

	x++
	fmt.Println(x)
	{
		y := 42
		fmt.Println(y)
	}

	foo()
	fmt.Println(x)
}

func foo() {
	fmt.Println("Hola")

}
