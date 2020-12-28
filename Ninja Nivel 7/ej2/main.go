package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func (p *persona) cambiame(nombre, apellido string, edad int) {
	p.nombre = nombre
	p.apellido = apellido
	p.edad = edad
	// (*p).nombre = nombre
	// (*p).apellido = apellido
	// (*p).edad = edad
}

func cambiame(p *persona) {
	p.nombre = "Anselmo"
	//(*p).nombre = "Anselmo"
}

func main() {

	p := persona{
		"Maxim",
		"Nako",
		26,
	}

	fmt.Println(p)

	cambiame(&p)

	fmt.Println(p)

	p.cambiame("Maxim", "Ya no anselmo", 27)
	fmt.Println(p)
}
