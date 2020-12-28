package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

func main() {

	p1 := persona{
		"James",
		"Bond",
		32,
	}
	p2 := persona{
		"Miss",
		"Moneypenny",
		27,
	}

	personas := []persona{p1, p2}
	fmt.Println(personas)

	bs, err := json.Marshal(personas)
	if err != nil {
		fmt.Println("Ha ocurrido un error:", err)
	}
	fmt.Printf("%s", bs)

}
