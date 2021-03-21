package main


import "fmt"
// Will be overwritten by 
// go build -ldflags="-X main.version 1.2.0" main.go
var version = "1.0.0"

func main() {
	fmt.Println("Version: ", version)
}