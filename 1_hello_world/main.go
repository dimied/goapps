package main

import (
	"fmt"
)

func deferIsLiFo() {
	for i:= 0; i < 5; i++ {
		defer fmt.Println("", i)
	}
}
func main() {

	fmt.Println("Hello World")

	deferIsLiFo()
}

