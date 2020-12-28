package main

import "fmt"

func main() {

	s1 := []string{"Batman", "Jefe", "Vestido de negro"}
	s2 := []string{"Robin", "ayúdame", "Vestido de colores"}

	ss := [][]string{s1, s2}

	for i, v := range ss {
		fmt.Println("Registro: ", i)
		for j, v2 := range v {
			fmt.Printf("\t índice de posición %v, \t %v \n", j, v2)
		}
	}

}
