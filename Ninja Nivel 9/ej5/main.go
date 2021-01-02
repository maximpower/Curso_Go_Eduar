package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var at int64
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&at, 1)

			fmt.Println(atomic.LoadInt64(&at))

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("El incremento al final vale :", at)
}
