package main

import (
	"errors"
	"fmt"
	"time"
)

type SomeResponse struct {
	Error    error
	Response string
}

type SomeResponseCollection struct {
	ErrorExists bool
	Responses   []SomeResponse
}

func barrier(
	requestFunc func(chan SomeResponse, string),
	resources ...string) SomeResponseCollection {

	requestNumber := len(resources)

	in := make(chan SomeResponse, requestNumber)
	defer close(in)

	result := SomeResponseCollection{}

	result.Responses = make([]SomeResponse, requestNumber)

	for _, resource := range resources {
		// We assume that only one value will be written to the channel
		go requestFunc(in, resource)
	}

	for i := 0; i < requestNumber; i++ {
		// Wait for data
		resp := <-in
		if resp.Error != nil {
			result.ErrorExists = true
		}

		result.Responses[i] = resp
	}

	return result
}

func main() {
	// Create a barrier that must be reached by all participants
	fmt.Println("Barrier design pattern")

	resources := []string{
		"hallo",
		"Hi",
		"Failure",
		"Successful",
		"...",
		"Done",
	}

	// Request and wait for completion
	result := barrier(requestFuncfunc, resources...)

	fmt.Printf("Completed with errors %v\n", result.ErrorExists)

	for _, value := range result.Responses {
		if value.Error != nil {
			fmt.Println("Result: ", value.Error)
		} else {
			fmt.Println("Result: ", value.Response)
		}

	}
}

func requestFuncfunc(resultChannel chan SomeResponse, resource string) {
	result := SomeResponse{}

	if len(resource)%2 == 0 {
		time.Sleep(time.Millisecond * 1501)
		result.Response = "OK for " + resource
	} else {
		time.Sleep(time.Millisecond * 1500)
		result.Error = errors.New("Failed for " + resource)
	}
	resultChannel <- result
}
