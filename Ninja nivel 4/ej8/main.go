package main

import "fmt"

func main() {

	dicc := map[string][]string{
		"maxim_nako": []string{"computadoras", "montaña", "libro"},
		"naty_shkv":  []string{"Ropita", "montaña", "vinito"},
		"nazik_nako": []string{"jefe", "guapo", "gym"},
	}
	for i, v := range dicc {
		fmt.Printf("A %v le gusta: \t %v \n", i, v)
	}
}
