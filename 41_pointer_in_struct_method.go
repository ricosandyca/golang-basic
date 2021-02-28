package main

import "fmt"

// Person is defining person dara
type Person struct {
	name string
	age  int
}

/**
Struct method with pass by value
So if you try to modify the parson value of the method, it won't effect to the main data
*/
func (person Person) decrementAge() {
	person.age--
}

/**
Struct method with pass by reference
So if you try to modify the parson value of the method, it will effect to the main data
*/
func (person *Person) incrementAge() {
	person.age++
}

func main() {
	me := Person{
		name: "Rico",
		age:  22,
	}

	fmt.Println(me) // 22
	me.decrementAge()
	fmt.Println(me) // 22 (no effect)
	me.incrementAge()
	fmt.Println(me) // 23 (has effect)
}
