package main

import "fmt"

func main() {

	year := 1994
	for {
		fmt.Println(year)
		if year < 2020 {
			year++
		} else {
			break
		}
	}

}
