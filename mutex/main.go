package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	var sharedMemory int

	fmt.Printf("num. CPU = %d, num Goroutines = %d\n", runtime.NumCPU(), runtime.NumGoroutine())

	wg.Add(100)
	for i := 1; i <= 100; i++ {
		go func() {
			m.Lock()
			v := sharedMemory
			runtime.Gosched()
			v++
			sharedMemory = v
			m.Unlock()
			fmt.Printf("num. CPU = %d, num Goroutines = %d\n", runtime.NumCPU(), runtime.NumGoroutine())
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Printf("Shared Memory = %d\n", sharedMemory)
	fmt.Printf("num. CPU = %d, num Goroutines = %d\n", runtime.NumCPU(), runtime.NumGoroutine())
}
