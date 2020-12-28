package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func (p persona) presentar() {
	fmt.Println("Hola, soy", p.nombre, p.apellido)
}

func main() {
	p1 := persona{
		"Maxim",
		"Nako",
		26,
	}

	p1.presentar()
}
