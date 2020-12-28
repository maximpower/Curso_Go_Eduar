package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	cont := 0
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			incr := cont
			runtime.Gosched()
			incr++
			cont = incr
			fmt.Println(incr)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("El incremento al final vale :", cont)
}
