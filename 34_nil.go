/**
nil is using to reset a variable into its default value
nil can only be using on some data types such as "interface", "function", "map", "slice", "pointer" and "channel"
*/

package main

import "fmt"

func main() {
	me := map[interface{}]interface{}{
		"name": "Rico",
		"age":  22,
	}
	nums := []int{1, 2, 3}
	friends := []string{}

	fmt.Println("me", me)
	fmt.Println("nums", nums)
	fmt.Println("friends", friends)
	fmt.Println("friends == nil", friends == nil)
	fmt.Println("are friends empty", len(friends) <= 0)

	fmt.Println("Resetting data...")
	me = nil
	nums = nil
	friends = nil

	fmt.Println("me", me)
	fmt.Println("nums", nums)
	fmt.Println("friends", friends)
	fmt.Println("friends == nil", friends == nil)
	fmt.Println("are friends empty", len(friends) <= 0)

	// the best way to check emptiness of an array is using len(array) <= 0
}
