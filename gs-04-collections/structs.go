package main

import "fmt"

func main() {
	type user struct {
		ID int
		FirstName string
		LastName string
	}
	var u user 
	u.ID = 1
	u.FirstName = "Authur"
	u.LastName = "Dent"
	fmt.Println(u)
	fmt.Println(u.FirstName)

	u2 := user { ID: 2, 
		FirstName: "Hadwin", 
		LastName: "Jiang",
	}

	fmt.Println(u2)
	
}