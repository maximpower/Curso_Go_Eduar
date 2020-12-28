package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Printf("Número de CPUs, %v\n", runtime.NumCPU())
	fmt.Printf("Número de Goroutines, %v\n", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Goroutine 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("Goroutine 2")
		wg.Done()
	}()

	wg.Wait()

	fmt.Printf("Número de CPUs, %v\n", runtime.NumCPU())
	fmt.Printf("Número de Goroutines, %v\n", runtime.NumGoroutine())
	fmt.Println("A punto de finalizar..")
}
