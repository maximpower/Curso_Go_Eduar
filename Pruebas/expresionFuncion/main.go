package main

import "fmt"

func main() {

	f := func() {
		fmt.Println("Mi primera expresión función")
	}

	f()

	f2 := func(v int) {
		fmt.Printf("Mi %v expresión función", v)
	}

	f2(2)

}
