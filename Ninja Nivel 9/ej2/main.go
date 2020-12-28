package main

import "fmt"

type persona struct {
	nombre string
	edad   int
}

type humano interface {
	hablar()
}

func (p *persona) hablar() {
	fmt.Println("Hola, soy", p.nombre)
}

func diAlgo(h humano) {
	h.hablar()
}

func main() {
	p1 := persona{
		"Anselmo",
		25,
	}

	p1.hablar()

	diAlgo(&p1)

}
