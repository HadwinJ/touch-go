package main

import (
	"fmt"
	"time"
)

func main() {
	first := time.Now()
	fmt.Printf("It is currently %v\n", first)
	time.Sleep(2 * time.Second)
	second := time.Now()
	fmt.Printf("It is now %v\n", second)
}
