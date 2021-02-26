package main

import "fmt"

// Hello is the type to declare simple sayHello function
type Hello func(string)

func sayHello(cbAfter Hello) {
	fmt.Println("Hello World")
	cbAfter("Rico")
}

func main() {
	// callback function
	sayHello(func(name string) {
		fmt.Println("Hello", name)
	})
}
