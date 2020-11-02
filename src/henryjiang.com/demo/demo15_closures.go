package main

import "fmt"

func main() {
	sayHello("Hello")

	// anonymous function decard
	func (msg string)  {
		fmt.Println(("Anonymous function: " + msg))
	}("Henry")

	printFunc := returnFunc()
	printFunc("Wayne")

	nextInt := intSequence()
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt2 := intSequence()
	fmt.Println(nextInt2())
	fmt.Println(nextInt2())
}

// sayHello - print a message
func sayHello(msg string)  {
	fmt.Println(msg)
}

func returnFunc() func(string) {
	return func(msg string)  {
		fmt.Println("Returned function: " + msg)
	}
}

func intSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}