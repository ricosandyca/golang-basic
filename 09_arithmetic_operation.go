package main

import "fmt"

func main() {
	var a = 20
	var b = 3
	var c float32 = float32(a) / float32(b)

	fmt.Println("1 + 2 =", 1+2)
	fmt.Println("6 - 2 =", 6-2)
	fmt.Println("6 * 6 =", 6*6)
	fmt.Println("20 / 6 =", c)
	fmt.Println("20 % 6 =", a%b)

	var i = 0
	i += 2
	i++
	i *= 2
	i -= 2
	fmt.Println("I =", i)

	var status = false
	fmt.Println("Status", !status)
}
