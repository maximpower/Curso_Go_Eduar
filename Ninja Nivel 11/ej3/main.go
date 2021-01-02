package main

import "fmt"

type errorPer struct {
	info string
}

func (e errorPer) Error() string {
	return fmt.Sprintf("aquí está el error: %v", e.info)
}

func foo(e error) {
	fmt.Println("foo ocurrió -", e, "\n", e.(errorPer).info)
}

func main() {
	e := errorPer{
		info: "Necesito dormir más",
	}

	foo(e)
}
