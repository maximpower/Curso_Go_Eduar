package main

import (
	"encoding/json"
	"fmt"
)

//Usuario estructura que permite trabajar con Usuarios
type Usuario struct {
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad"`
}

func main() {

	u1 := Usuario{
		"Maxim",
		26,
	}
	u2 := Usuario{
		"Natalia",
		28,
	}
	u3 := Usuario{
		"Nazar",
		31,
	}

	usuarios := []Usuario{u1, u2, u3}

	fmt.Println(usuarios)

	bs, err := json.Marshal(usuarios)

	if err != nil {
		fmt.Printf("No se ha podido convertir %v a json\n", usuarios)
	}

	fmt.Printf("%s", bs)

}
