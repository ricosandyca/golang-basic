/**
 * Variable declaration format
 * - var [variable_name] [variable_type] = [value]
 * - var [variable_name] [variable_type]
 * - var [variable_name] = [value]
 * - [variable_name] := [value]
 *
 * Multi variable declaration format
 * var (
 *   [variable_name_1] = [value]
 *   [variable_name_2] = [value]
 * )
 */

package main

import "fmt"

func main() {
	var rico string = "Rico"
	fmt.Println("First Name", rico)

	var sandyca = "Sandyca"
	fmt.Println("Middle Name " + sandyca)

	var novenza string
	novenza = "Novenza"
	fmt.Println("Last Name", novenza)

	age := 22
	fmt.Println("Age", age, "years old")

	var (
		hello = "Hello"
		world = "World"
	)
	fmt.Println(hello, world)
}
