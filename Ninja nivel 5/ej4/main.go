package main

import "fmt"

func main() {
	m := struct {
		nombre   string
		apellido string
		edad     int
	}{
		"Maxim",
		"Nako",
		26,
	}

	fmt.Println(m)
}
