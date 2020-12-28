package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	s := fmt.Sprintf("Valor de x, y, z respectivamente son: %v, %v, %v", x, y, z)

	fmt.Println(s)

}
