package main

import (
	"fmt"
	"sort"
)

// Persona .
type Persona struct {
	Nombre string
	Edad   int
}

//PorEdad .
type PorEdad []Persona

// Len .
func (a PorEdad) Len() int { return len(a) }

// Swap .
func (a PorEdad) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

//Less .
func (a PorEdad) Less(i, j int) bool { return a[i].Edad < a[j].Edad }

//PorNombre .
type PorNombre []Persona

// Len .
func (a PorNombre) Len() int { return len(a) }

// Swap .
func (a PorNombre) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

//Less .
func (a PorNombre) Less(i, j int) bool { return a[i].Nombre < a[j].Nombre }

func main() {
	p1 := Persona{"Natalia", 28}
	p2 := Persona{"Maxim", 26}
	p3 := Persona{"Stepan", 22}
	p4 := Persona{"Nazar", 31}

	personas := []Persona{p1, p2, p3, p4}

	fmt.Println(personas)

	sort.Sort(PorEdad(personas))
	fmt.Println(personas)
	sort.Sort(PorNombre(personas))
	fmt.Println(personas)

}
