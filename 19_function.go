package main

import "fmt"

func sayHelloWorld() {
	fmt.Println("Hello World")
}

func main() {
	// say hello 3 times
	for i := 1; i <= 3; i++ {
		sayHelloWorld()
	}
}
