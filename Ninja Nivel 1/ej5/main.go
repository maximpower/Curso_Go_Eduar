package main

import "fmt"

type myType int

var x myType
var y int

func main() {
	fmt.Printf("El tipo de la variable es: %T, y su valor es: %v\n", x, x)

	x = 42
	fmt.Println("X=", x)

	y = int(x)
	fmt.Printf("El tipo de la variable es: %T, y su valor es: %v\n", y, y)

}
