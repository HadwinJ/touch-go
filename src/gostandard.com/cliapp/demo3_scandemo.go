package main

import "fmt"

func main() {
	var name string
	fmt.Println("What is your name?")
	inputs, _ := fmt.Scanf("%s", &name)
	switch inputs {
	case 0:
		fmt.Printf("You must enter a name!\n")
	case 1:
		fmt.Printf("Hello %s! nice to meet you. \n", name)
	}
}
