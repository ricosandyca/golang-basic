/**
Interface is like a type contract
you can only use an interface as a type not a variable
Interface must have "method" defainitions only
*/

package main

import "fmt"

// IPerson is using for person interface
type IPerson interface {
	// name string (can't do like this)
	getFullNameAndAge() (string, int)
}

// Person is using for person type
type Person struct {
	firstName, lastName string
	age                 int
}

func (person Person) getFullNameAndAge() (string, int) {
	return person.firstName + " " + person.lastName, person.age
}

func sayHello(person IPerson) {
	name, _ := person.getFullNameAndAge()
	fmt.Println("Hello", name)
}

func main() {
	me := Person{
		firstName: "Rico",
		lastName:  "Sandyca",
		age:       22,
	}

	fmt.Println(me)
	fmt.Println(me.getFullNameAndAge())

	sayHello(me)
}
