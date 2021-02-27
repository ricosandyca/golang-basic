package main

import "fmt"

// Person is using to define person data
type Person struct {
	name, address string
	age           int
	isMarried     bool
}

func sayHello(person Person) {
	fmt.Println("Hello", person.name)
}

func main() {
	me := Person{
		name:      "Rico",
		age:       22,
		isMarried: false,
	}

	// parameters must have same position with the struct type
	you := Person{"Sandyca", "Jaksa Agung", 22, true}

	fmt.Println(me)
	fmt.Println(you)

	sayHello(me)
}
