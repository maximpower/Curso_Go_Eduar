package main

import (
	"fmt"

	"./perro"
)

type canino struct {
	nombre string
	edad   int
}

func main() {
	c1 := canino{
		"Dominich",
		perro.Edad(2),
	}

	fmt.Println(c1)
}
