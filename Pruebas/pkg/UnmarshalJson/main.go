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

type personas []persona

func main() {

	s := `[{"nombre":"James","apellido":"Bond","edad":32},{"nombre":"Miss","apellido":"Moneypenny","edad":27}]`

	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)
	ps := personas{}
	err := json.Unmarshal(bs, &ps)
	if err != nil {
		fmt.Println("Ocurrió un error en el Unmarshal")
	}

	for i, v := range ps {
		fmt.Println("\nPersona número: ", i+1)
		fmt.Printf("\n \tMi nombre es %v, %v, %v y tengo %v años", v.Apellido, v.Nombre, v.Apellido, v.Edad)
	}

	fmt.Printf("\nTodas las personas: \n\t %+v \n", ps)
}
