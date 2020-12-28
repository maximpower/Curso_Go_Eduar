package main

import "fmt"

type myType int

var x myType

func main() {
	fmt.Printf("El tipo de la variable es: %T, y su valor es: %v\n", x, x)

	x = 42
	fmt.Println("X=", x)

}
