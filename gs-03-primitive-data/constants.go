package main

import "fmt"

func main() {
	const pi = 3.1415
	fmt.Println(pi)
	// pi = 3.2, can not change

	const c = 3
	fmt.Println(c + 3)
	fmt.Println(c + 1.2)
}
