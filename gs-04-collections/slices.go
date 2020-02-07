package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	slice := arr[:]

	fmt.Println(arr, slice)

	arr[1] = 42
	slice[2] = 27
	fmt.Println(arr, slice)

	newSlice := []int{1, 2, 3}
	fmt.Println(newSlice)
	newSlice = append(newSlice, 4, 42, 27)
	fmt.Println(newSlice)
	s2 := newSlice[1:]
	s3 := newSlice[:2]
	s4 := newSlice[1:2]
	fmt.Println(s2, s3, s4)
}
