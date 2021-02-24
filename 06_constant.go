/**
 * Constant variable declaration format
 * - const [variable_name] [variable_type] = [value]
 * - const [variable_name] = [value]
 *
 * Multi constant variable declaration format
 * const (
 *   [variable_name_1] = [value]
 *   [variable_name_2] = [value]
 * )
 */

package main

import "fmt"

func main() {
	const pi float32 = 3.14
	fmt.Println("PI", pi)

	const name = "Rico"
	fmt.Println("Name", name)

	const (
		myName = "Rico"
		age    = 22
	)
	fmt.Println(myName, age)
}
