/**
 * Closure is an inner function that's accessing a variable from its parent
 * You can call the inner function is closure
 */

package main

import "fmt"

func main() {
	name := "Rico"
	sayHello := func() {
		fmt.Println("Hello", name)
		// name is accessed from parent variable
		// this sayHello function is closure
	}

	sayHello()
}
