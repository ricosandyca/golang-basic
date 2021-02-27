/**
Empty interface has similar definition with "any" in typescript
*/

package main

import "fmt"

func callMe(v interface{}) interface{} {
	return v
}

func main() {
	fmt.Println(callMe([]int{1, 2, 3}))
	fmt.Println(callMe("Rico"))
}
