package main

import "fmt"

/**
 * Function that returns single value
 */
func getFullName(firstName string, lastName string) string {
	return firstName + " " + lastName
}

/**
 * Function that returns multiple values
 * It returns fullname, firstname and lastname
 */
func getDetailName(firstName, lastName string) (string, string, string) {
	return firstName + " " + lastName, firstName, lastName
}

/**
 * Function that returns named multiple value
 * Named return values also work like variable declaration, so you don't need to declare it again
 */
func getMyFullName() (firstName, lastName, fullName string) {
	firstName = "Rico"
	fullName = firstName + " " + lastName
	return
	// or you can return like this
	// return firstName, lastName, fullName
	// or you can directly return like this
	// return "Rico", "Sandyca", "Novenza"
}

func main() {
	var firstName, lastName string
	firstName = "Rico"
	lastName = "Sandyca"

	fullName := getFullName(firstName, lastName)
	fmt.Println("Hello", fullName)

	// skip second and third return value
	detailFullName, _, _ := getDetailName(firstName, lastName)
	fmt.Println(detailFullName)

	_, _, myFullName := getMyFullName()
	fmt.Println(myFullName)
}
