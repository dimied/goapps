package main

import "fmt"

func withError(x int, y int) int {
	return (x + y) / (x - y)
}

func main() {
	fmt.Println("Hello errors")

	defer func() {
		error := recover()
		fmt.Println("Error is: ", error)
		fmt.Printf("Type: %T\n", error)
	}()

	fmt.Println("Result", withError(3, 2))
	fmt.Println("Result", withError(3, 3))
}
