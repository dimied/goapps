package main

import (
	"fmt"
)

type pipeFunc func(in <-chan int) <-chan int

func generator(max int) <-chan int {
	outChInt := make(chan int, 100)

	go func() {
		for i := 1; i <= max; i++ {
			outChInt <- i
		}

		close(outChInt)
	}()

	return outChInt
}

func add2(in <-chan int) <-chan int {
	out := make(chan int, 100)

	go func() {
		for v := range in {
			out <- v + 2
		}
		close(out)
	}()

	return out
}

func power(in <-chan int, pow int) <-chan int {
	out := make(chan int, 100)

	go func() {
		for v := range in {
			res := v
			// Not found pow for integers
			for i := 0; i < pow; i++ {
				res *= v
			}

			out <- res
		}
		close(out)
	}()

	return out
}

func power2(in <-chan int) <-chan int {
	return power(in, 2)
}

func sum(in <-chan int) <-chan int {
	out := make(chan int, 100)

	go func() {
		var sum int

		for v := range in {
			sum += v
		}

		out <- sum
		close(out)
	}()

	return out
}

func ExecuteInPipeline(amount int) int {
	firstCh := generator(amount)
	powerCh := power(firstCh, 2)
	thirdCh := sum(powerCh)

	result := <-thirdCh

	return result
}

// More general version
func ExecuteInOrder(amount int, funcs ...pipeFunc) int {
	ch := generator(amount)

	for _, f := range funcs {
		ch = f(ch)
	}

	thirdCh := sum(ch)

	result := <-thirdCh

	return result
}

func main() {
	fmt.Println("Pipeline design pattern")

	res := ExecuteInPipeline(50)

	fmt.Println("Computed: ", res)

	res = ExecuteInOrder(50, power2)

	fmt.Println("Computed: ", res)

	res2 := ExecuteInOrder(50, power2, add2)

	//+ additions by 2
	fmt.Println("Computed: ", res2, res+50*2)
}
