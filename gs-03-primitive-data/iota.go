package main

import "fmt"

const pi = 3.1415

const (
	first = iota
	second
	third
)

const (
	f0 = iota
	f1
	f2
)

func main() {
	fmt.Println(pi)
	fmt.Println(first, second, third)
	fmt.Println(f0, f1, f2)
}
