package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func example(i int) {
	if i%5 == 0 {
		fmt.Println("Skipped ", i)
		runtime.Goexit()
		return
	}
	fmt.Println("GO routine ", i)
}

func example2(i int) {
	if calledAsGoRoutine() {
		fmt.Println("Called as goroutine ", i)
	} else {
		fmt.Println("Called as function ", i)
	}
}

var myMutex sync.Mutex

func task1() {
	myMutex.Lock()
	defer myMutex.Unlock()
}

func task2() {
	myMutex.Lock()
	defer myMutex.Unlock()
}

func channelFunc1(i int, c chan int) {
	fmt.Println("Channel func ", i)
	c <- i
}

func simpleGoRoutineTest() {
	for i := 0; i < 10; i++ {
		go example(i)
	}
	time.Sleep(time.Second)

	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			go example2(i)
		} else {
			example2(i)
		}
	}
	time.Sleep(time.Second)
}

func main() {
	fmt.Println("Concurrency example")
	fmt.Println("You have #cpus = ", runtime.NumCPU())
	// define maximal number of proce
	runtime.GOMAXPROCS(4)

	fmt.Print("\nSimple test\n\n")
	simpleGoRoutineTest()

	fmt.Print("\nUnbuffered\n\n")
	unbufferedChannel()
	time.Sleep(time.Microsecond * 500)

	fmt.Print("\nBuffered\n\n")
	bufferedChannel(4, 10, time.Microsecond*2)

	fmt.Print("\nIn-Out\n\n")
	inAndOutChannels()

	fmt.Print("\nFan-In pattern\n\n")
	fanIn()

	fmt.Print("\nFan-Out pattern\n\n")
	fanOut()
}

func unbufferedChannel() {
	//Unbuffered channel
	c := make(chan int)
	numRoutines := 5
	// Using channels
	for i := 0; i < numRoutines; i++ {
		go channelFunc1(i, c)
	}
	for i := 0; i < numRoutines; i++ {
		fmt.Println("Receiving value from ", <-c)
	}
}

func bufferedChannel(size int, numRoutines int, sleepTime time.Duration) {
	//Unbuffered channel
	c := make(chan int, size)

	// Using channels
	for i := 20; i < 20+numRoutines; i++ {
		go channelFunc1(i, c)
	}

	for i := 0; i < numRoutines; i++ {
		fmt.Println("Buffered: Receiving value from ", <-c)
		time.Sleep(sleepTime)
	}
}

func inAndOutChannels() {
	inChannel := make(chan int, 3)
	outChannel := make(chan bool)
	go inAndOutChannel(inChannel, outChannel)

	for i := 10; i < 20; i++ {
		fmt.Println("Send ", i)
		inChannel <- i
	}
	close(inChannel)
	fmt.Println("Wait to finish")
	<-outChannel
	fmt.Println("Finished")
}

// Using one-way channels
func inAndOutChannel(in <-chan int, out chan<- bool) {
	// Infinite loop
	for {
		value, ok := <-in
		if !ok {
			fmt.Println("Not OK from input channel")
			break
		}
		fmt.Println("Receiving ", value)
	}
	//Send completion signal
	out <- true
}

func fanIn() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	out := fanInFunc(ch1, ch2)
	var i int

	for {
		//Send
		if i%2 == 0 {
			ch1 <- i
		} else {
			ch2 <- i
		}
		i++

		//Receive
		select {
		case val := <-out:
			fmt.Println("Value: ", val)
		case <-time.After(time.Millisecond * 20):
			fmt.Println("Timeout ...")
			return
		}
	}
}

func fanInFunc(inChannel1, inChannel2 <-chan int) chan int {
	out := make(chan int)
	go func() {
		for {
			var num int

			// Deferred sending to output
			select {
			case num = <-inChannel1:
				time.Sleep(time.Millisecond * time.Duration(num))
				out <- num * 2
			case num = <-inChannel2:
				time.Sleep(time.Millisecond * time.Duration(num))
				out <- num * 3
			}

		}
	}()

	return out
}

func fanOut() {

}

// Checks if we called the function as goroutine
func calledAsGoRoutine() bool {
	// skip:
	// 1: for calling function
	// 2: runtime.Callers
	// 3: calledAsGoRoutine()
	// 4: runtime.goexit()
	count := runtime.Callers(4, make([]uintptr, 1))

	return count == 0
}
