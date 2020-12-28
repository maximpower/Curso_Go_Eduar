package main

import "fmt"

type vehiculo struct {
	puertas int
	color   string
}

type camion struct {
	vehiculo     vehiculo
	cuatroRuedas bool
}

type sedan struct {
	vehiculo vehiculo
	lujoso   bool
}

func main() {

	c := camion{
		vehiculo{
			2,
			"amarillo",
		},
		false,
	}

	s := sedan{
		vehiculo{
			5,
			"negro",
		},
		true,
	}

	fmt.Println(c, s)

	fmt.Printf("el camión es de cuatro ruedas? %v.\n", c.cuatroRuedas)
	fmt.Printf("Es el sedán lujoso? %v.\n", s.lujoso)
}
