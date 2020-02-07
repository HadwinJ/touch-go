package main

import (
	"errors"
	"fmt"
)

func main() {
	var port2 int
	fmt.Println("Hello, playground")
	port := 3000
	port2, err := startWebServer(port, 2)
	fmt.Println(err)
	fmt.Println(port2)
}

func startWebServer(port, numberOfRetry int) (int, error) {
	fmt.Println("Starting Server")
	// do something
	fmt.Println(port)
	//
	fmt.Println("Server Started on port", port)
	fmt.Println("Number of retries", numberOfRetry)
	return port, errors.New("Something web wrong")

}
