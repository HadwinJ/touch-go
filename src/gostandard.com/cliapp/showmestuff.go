package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {
	args := os.Args
	fmt.Println(args)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	text, _ := reader.ReadString('\n')
	fmt.Println("Hello " + text)
	fmt.Println("Our current version of Go is " + runtime.Version())
	fmt.Printf("Our current version of Go is %v! Sweet\n", runtime.Version())
	fmt.Printf("\nThis binary was built for %s in %s\n", runtime.Version(), runtime.GOOS)
}
