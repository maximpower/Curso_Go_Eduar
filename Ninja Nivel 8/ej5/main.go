package main

import (
	"fmt"
	"sort"
)

type Usuario struct {
	Nombre   string
	Apellido string
	Edad     int
	Dichos   []string
}

//PorApellido es un slice de usuarios que se ordenará por apellido
type PorApellido []Usuario

//PorEdad es un slice de usuarios que se ordenará por edad
type PorEdad []Usuario

//Len .
func (a PorEdad) Len() int { return len(a) }

// Swap .
func (a PorEdad) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less .
func (a PorEdad) Less(i, j int) bool { return a[i].Edad < a[j].Edad }

//Len .
func (a PorApellido) Len() int { return len(a) }

// Swap .
func (a PorApellido) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less .
func (a PorApellido) Less(i, j int) bool { return a[i].Apellido < a[j].Apellido }

func printUsers(usuarios []Usuario) {

	for _, v := range usuarios {
		// Primero ordenamos los "Dichos" alfabéticamente porque lo pide el Ej.
		sort.Strings(v.Dichos)

		fmt.Printf("Usuario: %v %v %v", v.Nombre, v.Apellido, v.Edad)
		fmt.Println("\nDichos:")
		for i, v2 := range v.Dichos {
			i++
			fmt.Printf("\t%v: %v\n", i, v2)
		}
	}
}

func main() {
	u1 := Usuario{
		Nombre:   "Eduar",
		Apellido: "Tua",
		Edad:     32,
		Dichos: []string{
			"Cachicamo diciéndole a morrocoy conchudo",
			"La mona, aunque se vista de seda, mona se queda",
			"Poco a poco se anda lejos",
		},
	}

	u2 := Usuario{
		Nombre:   "Carlos",
		Apellido: "Pérez",
		Edad:     27,
		Dichos: []string{
			"Camarón que se duerme se lo lleva la corriente",
			"A ponerse las alpargatas que lo que viene es joropo",
			"No gastes pólvora en zamuro",
		},
	}

	u3 := Usuario{
		Nombre:   "Che",
		Apellido: "López",
		Edad:     54,
		Dichos: []string{
			"Ni lava ni presta la batea",
			"Hijo de gato, caza ratón",
			"Más vale pájaro en mano que cien volando",
		},
	}

	usuarios := []Usuario{u1, u2, u3}
	usuarios2 := []Usuario{u1, u2, u3}

	fmt.Println("----------------------------------------")
	fmt.Println("Usuarios iniciales: ")
	fmt.Println("----------------------------------------")
	printUsers(usuarios)

	fmt.Println("----------------------------------------")

	sort.Sort(PorEdad(usuarios))
	fmt.Println("Usuarios ordenados por edad: ")
	fmt.Println("----------------------------------------")
	printUsers(usuarios)

	fmt.Println("----------------------------------------")
	printUsers(usuarios2)
	fmt.Println("----------------------------------------")
	fmt.Println("Usuarios depués del sort por apellido")
	sort.Sort(PorApellido(usuarios2))
	fmt.Println("----------------------------------------")
	printUsers(usuarios2)

}
