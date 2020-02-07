package main

import "fmt"

func main() {
	var firstName *string = new(string)
	fmt.Println(firstName)
	*firstName = "Arthur"
	fmt.Println(*firstName)

	lastName := "Zhou"
	fmt.Println(lastName)
	ptr := &lastName
	fmt.Println(ptr, *ptr)

	lastName = "Wang"
	fmt.Println(ptr, *ptr)
}
