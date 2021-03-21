package main

import (
	"runtime/pprof"
	"fmt"
	"time"
	"runtime"
	"os"
)

func addText(input string) string{
	return input+" Hallo"
}

//
// TODO: Take a look at sync.Pool for reducing allocating memory
//
func main() {
	
	s := "Hallo"

	startedAt := time.Now()
	numRuns := 0
	for {
		if time.Since(startedAt) > time.Second {
			var ms runtime.MemStats
			runtime.ReadMemStats(&ms)
			fmt.Println("Heap size = ",ms.HeapAlloc, ms.NextGC)
			fmt.Println("Stack: ", ms.StackSys, ms.StackInuse)
			// Number of garbage collection
			fmt.Println("#GC = ",ms.NumGC)
			// Can be useful to understand how much is used for GC
			fmt.Println("#GC time = ", ms.GCCPUFraction)
			fmt.Println("Len: ", len(s))
			startedAt = time.Now()

			if numRuns == 5 {
				f, ok := os.Create("myprofile.pprof")

				if ok != nil {
					fmt.Println("Failed to create profile file")
				} else {
					// Use
					// go tool pprof -http localhost:3000 myprofile.pprof
					err := pprof.WriteHeapProfile(f)
					if err != nil {
						fmt.Println("Failed to write to profile file", err)
					}
					
					f.Close()
				}
			}
			numRuns++
		}
		s = addText(s)
	}
}