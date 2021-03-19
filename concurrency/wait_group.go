package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	fmt.Println("Wait groups")

	waitGroupForOne()

	waitGroupForMany(5)

	waitGroupWithCallbacks()
}

func waitGroupForOne() {
	var wg sync.WaitGroup

	// Wait till one finishes his wors
	wg.Add(1)

	go func() {
		fmt.Println("Working ...")
		// Signal completion
		wg.Done()
	}()
	fmt.Println("Waiting ...")
	// Wait until everything is completed
	wg.Wait()
	fmt.Println("Finished")
}

func waitGroupForMany(numberOfRoutines int) {
	var wg sync.WaitGroup
	wg.Add(numberOfRoutines)
	fmt.Println("Waiting for many")
	for i := 0; i < numberOfRoutines; i++ {
		go func(someId int) {
			fmt.Println("Working ... ", someId, i)
			// without it "fatal error: all goroutines are asleep - deadlock!"
			wg.Add(-1) // == wg.Done()
		}(i)
	}
	fmt.Println("Waiting ...")
	wg.Wait()
	fmt.Println("Finished")
}

type AsyncConverter struct {
	wg sync.WaitGroup
}

func (asyncConverter *AsyncConverter) Wait() {
	asyncConverter.wg.Wait()
}

func NewAsyncConverter(size int) AsyncConverter {
	result := AsyncConverter{}
	result.wg.Add(size)
	return result
}

func (asyncConverter *AsyncConverter) addToProcess(value string, callback func(string)) {
	go func() {
		callback(strings.ToTitle(value))
		asyncConverter.wg.Done()
	}()
}

func waitGroupWithCallbacks() {
	fmt.Print("\nWith callbacks\n")
	
	values := []string{"Hallo", "Hi", "Guten Tag"}
	print := func(value string) {
		fmt.Println(value)
	}

	converter := NewAsyncConverter(len(values))

	for _, val := range values {
		converter.addToProcess(val, print)
	}

	converter.Wait()

}
