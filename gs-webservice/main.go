package main

import (
	"fmt"
	"github/hadwinjiang/webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}

	fmt.Println(u)
	fmt.Println("Hello from a module, Gophers!")
}
