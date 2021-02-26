package main

import "fmt"

func main() {
	// there's only one looping statement in Golang that is for

	length := 0
	for length < 3 {
		fmt.Println("Count", length)
		length++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Index", i)
	}

	// for range (for array looping)
	names := []string{
		"Rico",
		"Sandyca",
		"Novenza",
	} // slice

	for index, value := range names {
		fmt.Println("Index:", index, "Value:", value)
	}

	for _, value := range names {
		fmt.Println("Value", value)
	}
}
