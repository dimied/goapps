package main

import (
	"fmt"
	"math"
)

/*
 Asks the user for some input and 
 calculates the area of requested geometric shape
*/
func main() {

	fmt.Print("Select type (t=triangle, c=circle, s=square): ")
	var selectedType string
	fmt.Scanf("%s", &selectedType)
	fmt.Println("You selected", selectedType)
	if selectedType == "t" {
		var height float32
		var width float32
		fmt.Print("Input height: ")
		fmt.Scanf("%f", &height)
		fmt.Print("Input width: ")
		fmt.Scanf("%f", &width)
		area := height * width * 0.5
		fmt.Println("Area = ", area)
	} else if selectedType == "c" {
		var radius float32
		fmt.Print("Input radius: ")
		fmt.Scanf("%f", &radius)
		area := radius * radius * math.Pi
		fmt.Println("Area = ", area)
	} else if selectedType == "s" {
		var length float32
		fmt.Print("Input edge length: ")
		fmt.Scanf("%f", &length)
		area := length * length
		fmt.Println("Area = ", area)
	} else {
		fmt.Println("Nothing valid selected")
	}
}
