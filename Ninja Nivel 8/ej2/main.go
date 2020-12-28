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

	s := `[{"nombre":"Maxim","edad":26},{"nombre":"Natalia","edad":28},{"nombre":"Nazar","edad":31}]`
	v := []Usuario{}

	err := json.Unmarshal([]byte(s), &v)
	if err != nil {
		fmt.Println("No se ha podido decodificar el Json")
	}

	fmt.Printf("El tipo es: %T, El valor: %+v", v, v)
}
