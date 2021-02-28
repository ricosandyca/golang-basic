package main

import "fmt"

// Person is using to define person data
type Person struct {
	name string
	age  int
}

func increment(num *int) {
	*num++
}

func changeName(person *Person) {
	person.name = "Sandyca"
}

func main() {
	num := 5
	fmt.Println(num) // 6
	increment(&num)
	increment(&num)
	fmt.Println(num) // 7

	me := Person{
		name: "Rico",
		age:  22,
	}
	fmt.Println(me) // Rico
	changeName(&me)
	fmt.Println(me) // Sandyca
}
