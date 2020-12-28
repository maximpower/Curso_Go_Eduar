package main

import (
	"fmt"
)

// Human interfaz humano.
type Human interface {
	presentar()
}

type persona struct {
	nombre   string
	apellido string
}

type agenteSecreto struct {
	persona
	lpm bool
}

func (as agenteSecreto) presentar() {
	fmt.Println("Hola, soy", as.nombre, as.apellido)
}
func (p persona) presentar() {
	fmt.Println("Hola, soy la persona", p.nombre, p.apellido)
}

func bar(h Human) {
	switch h.(type) {
	case persona:
		fmt.Println("El nombre de la persona es", h.(persona).nombre)

	case agenteSecreto:
		fmt.Println("El nombre del agente secreto es", h.(agenteSecreto).nombre)
	default:
		fmt.Println("Otro humano no identidicado")
	}
}

func main() {

	as := agenteSecreto{
		persona{
			"Maxim",
			"Nako",
		},
		true,
	}
	p := persona{
		"Maxim",
		"Nako",
	}

	// p.presentar()
	// as.presentar()
	bar(p)
	bar(as)

}
