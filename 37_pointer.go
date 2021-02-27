/**
Go by default passes data by value not reference
It's like duplicating data to a new memory allocation
*/

package main

import "fmt"

// Person is using to define person type
type Person struct {
	name string
	age  int
}

func changeName(person Person) {
	// passing person by value (not reference)
	person.name = "Sandyca"
}

func changeNameByRef(person *Person) {
	person.name = "Holla..."
}

func main() {
	me := Person{
		name: "Rico",
		age:  22,
	}

	you := me

	insideChangeName := func(person Person) {
		// passing person by value (not reference)
		person.name = "Novenza"
	}

	you.name = "Anda"
	fmt.Println(me) // Rico

	changeName(me)
	fmt.Println(me) // Rico

	insideChangeName(me)
	fmt.Println(me) // Rico

	me.name = "Rico Sandyca"
	fmt.Println(me) // Rico Sandyca

	changeNameByRef(&me)
	fmt.Println(me) // Holla... (changed)
}
