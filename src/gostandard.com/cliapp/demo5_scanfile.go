package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println("File contains: " + scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}
}
