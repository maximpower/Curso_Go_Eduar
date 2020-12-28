package main

import "fmt"

func foo() {
	fmt.Println("Esto se imprimirá al finalizar el main")
}

func main() {
	defer foo()
	fmt.Println("Esto se debería imprimir último, pero como foo() tiene defer, se imprimirá después")

}
