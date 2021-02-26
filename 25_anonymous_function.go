package main

import "fmt"

func main() {
	sayHello := func(name string) {
		fmt.Println("Hello", name)
	}

	(func() {
		fmt.Println("It's running...")
	})()

	func() {
		fmt.Println("So am i XD...")
	}()

	sayHello("Rico")
}
