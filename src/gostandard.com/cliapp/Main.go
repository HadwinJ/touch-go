package main

import (
	"fmt"
	"runtime"
)

func main ()  {
	fmt.Printf("\nThis binary was built for %s in %s\n", runtime.Version(), runtime.GOOS)
}