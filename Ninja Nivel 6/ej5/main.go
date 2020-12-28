package main

import (
	"fmt"
	"math"
)

type forma interface {
	calcularArea() float64
}

type cuadrado struct {
	base   float64
	altura float64
}

type circulo struct {
	radio float64
}

func (c cuadrado) calcularArea() float64 {
	area := c.base * c.altura

	//fmt.Println(area)
	return area
	// fmt.Println("El área del cuadrado es:", area)
}
func (c circulo) calcularArea() float64 {
	area := math.Pi * (c.radio * c.radio)

	// fmt.Println(area)

	return area
}

func info(f forma) {
	switch f.(type) {
	case cuadrado:
		fmt.Println("El área del cuadraro es:", f.(cuadrado).base*f.(cuadrado).altura)
	case circulo:
		fmt.Println("El área del círculo es", math.Pi*f.(circulo).radio*f.(circulo).radio)
	default:
		fmt.Println("No es ni circulo ni cuadrado")

	}
}

func main() {
	cir := circulo{
		2.50,
	}

	cuad := cuadrado{
		2.0,
		4.2,
	}

	cir.calcularArea()
	cuad.calcularArea()

	info(cir)
	info(cuad)
}
