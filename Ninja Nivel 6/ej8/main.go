package main

import "fmt"

func funcionRetornaFuncion() func() {

	fmt.Println("La primera función ejecutada, se retorna la función...")
	return func() {
		fmt.Println("Esta es la función retornada")
	}

}

func main() {

	f := funcionRetornaFuncion()

	fmt.Println("Pasaremos a imprimir la función en: \n5")
	fmt.Println(4)
	fmt.Println(3)
	fmt.Println(2)
	fmt.Println(1)
	f()
}
