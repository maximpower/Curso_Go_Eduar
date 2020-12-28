package main

import "fmt"

func bar() func() int {
	return func() int {
		return 1492
	}
}

func main() {
	fmt.Println(bar()())
}
