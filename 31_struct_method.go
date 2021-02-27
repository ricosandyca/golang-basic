package main

import "fmt"

// Person is using to define person data
type Person struct {
	name string
	age  int
}

func (person Person) sayHello() {
	fmt.Println("Hello", person.name)
}

func main() {
	me := Person{
		name: "Rico Sandyca",
		age:  22,
	}

	fmt.Println(me)
	me.sayHello()
}
