/**
- public variable (function or variable) means that the variable can be accessed by other packages
	The requirement to make a public variable, you must start the name of the variable by uppercase character
	for example: SayHello, MakeData, etc (public)
- private variable (function or variable) means that the variable can only be access by its package
	The requirement to make a private variable , you must start the name of variable by lowecase character
	for example: sayHello, makeData, etc (private)
*/
package main

import (
	"fmt"
	alias "golang-basic/helper"
)

func init() {
	fmt.Println("A")
}

func main() {
	alias.HelloWorld()
}
