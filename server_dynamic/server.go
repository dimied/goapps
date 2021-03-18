package main

import (
	"errors"
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os"
	"strconv"
)

// Title-case, i.e. exported
var MyBlockTemplate = template.Must(template.ParseGlob("./templates/*"))

type BlockParams struct {
	Title   string
	Content string
}

type PageValues struct {
	Title   string
	Content string
	Blocks  []*BlockParams
}

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

	http.HandleFunc("/", ServeIndexPage)
	http.ListenAndServe(port_str, nil)
}

func ServeIndexPage(writer http.ResponseWriter, r *http.Request) {

	name := "Unknown visitor"

	block1 := BlockParams{
		"Hello",
		"Template content",
	}

	block2 := BlockParams{
		"Another",
		"Another template content",
	}

	switch r.Method {
	case "POST":
		name = r.FormValue("name")
		r.ParseForm()
		// bad case
		//http.Error(writer, "Method not supported: "+r.Method,
		//http.StatusMethodNotAllowed)
	}

	page := PageValues{
		"Page",
		"Hello " + name,
		[]*BlockParams{&block1, &block2},
	}

	err := MyBlockTemplate.ExecuteTemplate(writer, "index", &page)
	if err != nil {
		fmt.Println("Failed to fill template: ", err)
		os.Exit(2)
	}
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
