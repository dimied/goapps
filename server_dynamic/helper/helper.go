package helper

import (
	"errors"
	"fmt"
	"net"
	"strconv"
)

func FindAvailablePort(port uint16, verbose bool) (uint16, error) {

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



