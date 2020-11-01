package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Usage: dinnertotal <Total Meal Amount> <Tip Percentage>")
		fmt.Println("Example: dinnertotal 20 10")
	} else {
		if len(args) != 0 {
			fmt.Println(("You must enter 2 arguments! type /help for help"))
		} else {
			mealTotal, _ := strconv.ParseFloat(args[0], 32)
			tipAmount, _ := strconv.ParseFloat(args[1], 32)
			totalPrice := calculateTotal(float32(mealTotal), float32(tipAmount))
			fmt.Printf("Your meal total will be %.2f \n", totalPrice)
		}
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	text, _ := reader.ReadString('\n')
	fmt.Println("Hello " + text)
	fmt.Println("Our current version of Go is " + runtime.Version())
	fmt.Printf("Our current version of Go is %v! Sweet\n", runtime.Version())
	fmt.Printf("\nThis binary was built for %s in %s\n", runtime.Version(), runtime.GOOS)
}
func calculateTotal(meatlTotal float32, tipAmount float32) float32 {
	totalPrice := meatlTotal + (meatlTotal * (tipAmount / 100))
	return totalPrice
}
