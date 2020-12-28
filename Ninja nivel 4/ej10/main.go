package main

import "fmt"

func main() {

	dicc := map[string][]string{
		"maxim_nako": []string{"computadoras", "montaña", "libro"},
		"naty_shkv":  []string{"Ropita", "montaña", "vinito"},
		"nazik_nako": []string{"jefe", "guapo", "gym"},
	}
	dicc["stepan_shkv"] = []string{"dibujar", "coches", "velocidad"}
	for i, v := range dicc {
		fmt.Printf("A %v le gusta: \t %v \n", i, v)
	}

	delete(dicc, "stepan_shkv")
	fmt.Println("________________________________")
	for i, v := range dicc {
		fmt.Printf("A %v le gusta: \t %v \n", i, v)
	}

}
