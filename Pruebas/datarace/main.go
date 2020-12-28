package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("Número de CPU:", runtime.NumCPU())
	fmt.Println("Número de Goroutines:", runtime.NumGoroutine())
	var contador int64

	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("Contador:", atomic.LoadInt64(&contador))
			wg.Done()
		}()
		fmt.Println("Número de Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Contador:", contador)
}
