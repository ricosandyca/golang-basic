/**
Another way to declare memory allocation variable you can use function "new"
The function will return a new empty data
*/

package main

import "fmt"

// Person is using for person data type
type Person struct {
	firstName, lastName string
}

func main() {
	me := new(Person) // the parameter is type
	age := new(int)
	age2 := age

	fmt.Println(me) // empty
	me.firstName = "Rico"
	fmt.Println(me) // {Rico }

	fmt.Println(age, age2)
	*age = 12
	fmt.Println(age, age2)
}
