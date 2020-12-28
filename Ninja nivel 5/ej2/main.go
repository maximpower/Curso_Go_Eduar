package main

import "fmt"

type persona struct {
	nombre, apellido string
	saboresFav       []string
}

func main() {

	p1 := persona{
		"Maxim",
		"Nako",
		[]string{
			"pistacho",
			"coco",
			"fresa",
		},
	}
	p2 := persona{
		"Naty",
		"Shkv",
		[]string{
			"Turrón",
			"Vinito",
			"Maxik",
		},
	}

	fmt.Printf("A %v %v le gustan los helados de sabores: \n", p1.nombre, p1.apellido)
	for _, v := range p1.saboresFav {
		fmt.Println("\t", v)
	}
	fmt.Printf("A %v %v le gustan los helados de sabores: \n", p2.nombre, p2.apellido)
	for _, v := range p2.saboresFav {
		fmt.Println("\t", v)
	}

	m := make(map[string]persona)

	m[p1.apellido] = p1
	m[p2.apellido] = p2

	fmt.Printf("%+v\n", m)

	for _, v := range m {
		fmt.Printf("Los sabores favoritos de %v, son:\n", v.nombre)
		for _, v2 := range v.saboresFav {
			fmt.Printf("%v\n", v2)
		}
	}

}
