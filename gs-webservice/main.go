package main

import (
	"github/hadwinjiang/webservice/controllers"
	"net/http"
)

func main() {
	// u := models.User{
	// 	ID:        2,
	// 	FirstName: "Tricia",
	// 	LastName:  "McMillan",
	// }

	// fmt.Println(u)
	// fmt.Println("Hello from a module, Gophers!")
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}
