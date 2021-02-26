package main

import "fmt"

func main() {
	name := "Rico"

	switch name {
	case "Rico":
		fmt.Println("Welcome back Rico")
	default:
		fmt.Println(name + ", You need to signup first :)")
	}

	// switch short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("The name is too long")
	case false:
		fmt.Println("The name is too short")
	default:
		fmt.Println(":)")
	}

	// switch with no condition (similar like if statement)
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("The name is too long")
	case length <= 5:
		fmt.Println("The name is too short")
	}
}
