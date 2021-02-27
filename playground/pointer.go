package main

import "fmt"

// Person is using to define person attributes
type Person struct {
	firstName, lastName string
}

func main() {
	x := 5
	y := &x // y is equal to address of x variable
	fmt.Println(x, y)

	x = 4
	fmt.Println(x, y)

	// read the wildcard (*) like "value of"
	*y = 2 // chenge the value of that pointer y become 2
	fmt.Println(x, y)

	me := Person{"Rico", "Sandyca"}
	you := &me

	you.firstName = "Changed"
	you = &Person{"Hello", "World"}

	fmt.Println(me)
	fmt.Println(*you)
}
