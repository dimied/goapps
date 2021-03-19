package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	value int
}

type CounterWithMutex struct {
	sync.Mutex
	Counter	
}

// To see data races, run
// go run -race mutexes.go
func main() {
	fmt.Println("Mutexes")
	numCounts := 1000
	
	counter:= Counter{}
	counter2:= CounterWithMutex{}

	for i:=0; i < numCounts; i++ {
		go func() {
			counter.value++
		}()
	}

	for i:=0; i < numCounts; i++ {
		go func() {
			counter2.Lock()
			counter2.value++
			counter2.Unlock()
		}()
	}

	time.Sleep(time.Second)
	// Will most likely show wrong result
	fmt.Println("Counter (naive):", counter.value)
	fmt.Println("Counter (mutex):", counter2.value)
}