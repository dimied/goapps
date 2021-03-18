package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Starting server ...")

	directory, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current directory path")
		os.Exit(1)
	}
	fmt.Println("Current dir: ", directory)

	port, portError := findAvailablePort(2000, false)
	if portError != nil {
		fmt.Println("Error: ", portError)
		os.Exit(1)
	}

	port_str := ":" + strconv.FormatUint(uint64(port), 10)

	fmt.Println("Starting server on http://localhost" + port_str)
	http.ListenAndServe(port_str, http.FileServer(http.Dir(directory)))
}

func findAvailablePort(port uint16, verbose bool) (uint16, error) {

	inputPort := port

	for port < 4000 {
		port_str := strconv.FormatUint(uint64(port), 10)
		if verbose {
			fmt.Println("Check port : ", port)
		}

		ln, err := net.Listen("tcp", ":"+port_str)

		if err != nil {
			port++
			continue
		}
		err = ln.Close()

		if err != nil {
			port++
			continue

		}
		// Found !
		return port, nil
	}
	return inputPort, errors.New("Failed to find any port")
}
