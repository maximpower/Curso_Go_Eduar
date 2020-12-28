package main

import "fmt"

var x int
var y string
var z bool

func main() {

	fmt.Printf("Valor de x: %v \n", x)
	fmt.Printf("Valor de y: %v \n", y)
	fmt.Printf("Valor de z: %v \n", z)

}

/*
	Los valores no inicializados el compilador les asigna lo que se
	conoce como el "Valor cero"
*/
