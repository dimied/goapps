package main;

import (
	"fmt"
	"flag"
)

func main() {
	fmt.Println("Command line arguments")
	arg1 := flag.String("a", "10", "Parameter a")
	arg2 := flag.String("b", "", "Parameter b")

	flag.Parse()


	if *arg1 == "" {
		fmt.Println("Parameter a empty")
	}

	if *arg2 == "" {
		fmt.Println("Parameter b empty")
	}
}