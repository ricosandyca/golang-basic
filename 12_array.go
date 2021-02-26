/**
 * Array data type declaration format
 * - var variable_name [array_length]data_type => var nums [3]int
 * - var variable_name = [array_length]data_type{value, value, ...,}
 * - var variable_name = [...]data_type{value,value,....,}
 * - variable_name := {value, value...,}
 *
 * Array function
 * - len(array) => count the array length (not the array element in array)
 * - cap(array) => count the array length from the first index to the last index
 */

package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Rico"
	names[1] = "Sandyca"
	names[2] = "Novenza"

	fmt.Println("Names", names)

	var badges = [3]string{
		"Back-End Developer",
		"Front-End Developer",
	}

	fmt.Println("Badges", badges)
	fmt.Println("Count of badges", len(badges)) // 3
	fmt.Println("First Badges", badges[0])
	fmt.Println("Third Badges", badges[2])

	favoriteColors := [3]string{
		"Yellow",
		"Pink",
		"Purple",
	}
	fmt.Println("Favorite Colors", favoriteColors)

	var friends = [...]string{
		"John",
		"Doe",
	}

	fmt.Println("Friends", friends)
	fmt.Println("Count of friends", len(friends))

}
