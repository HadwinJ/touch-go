package main

func main() {
	var i int
	for i < 5 {

		println(i)
		i++

		if i == 3 {
			break
		}
		if i == 2 {
			continue
		}
		println("continuing..")
	}

	for j := 0; j < 5; j++ {
		println("J", j)
	}

	// Infinite loop
	i = 1
	for {
		if i == 5 {
			break
		}
		println(i)
		i++
	}

	slice := []int{1, 2, 3}
	for index, value := range slice {
		println(index, value)
	}

	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for _, v := range wellKnownPorts {
		println(v)
	}

}
