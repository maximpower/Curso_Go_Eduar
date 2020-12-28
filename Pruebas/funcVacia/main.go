package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("La función es anónima")
	}()

}

func foo() {

	fmt.Println("foo se ejecutó")

}
