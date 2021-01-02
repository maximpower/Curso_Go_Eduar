package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	cont := 0
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			incr := cont
			incr++
			cont = incr
			fmt.Println(cont)
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("El incremento al final vale :", cont)
}
