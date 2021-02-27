package main

import "fmt"

func main() {
	var name = "Rico"
	var r = name[0]
	var rString = string(r)

	fmt.Println("Name", name)
	fmt.Println("R in Byte", r)
	fmt.Println("R in String", rString)
}
