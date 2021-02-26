/**
 * Golang doesn't support default value of function parameters
 */

package main

import "fmt"

func sayHello(name string, emoji string) {
	fmt.Println("Hello", name, emoji)
}

func main() {
	sayHello("Rico", ":)")
	sayHello("Sandyca", ":v")
	sayHello("Novenza", ":D")
}
