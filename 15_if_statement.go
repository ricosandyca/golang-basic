package main

import "fmt"

func main() {
	age := 22

	if age <= 18 {
		fmt.Println("You're a child")
	} else if age > 18 && age <= 30 {
		fmt.Println("You're a teen")
	} else {
		fmt.Println("You're an adult")
	}

	// if short statement
	if name := "Rico"; name == "Rico" {
		fmt.Println("Hello me")
	} else {
		fmt.Println("Hello", name)
	}
}
