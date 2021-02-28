// strconv => string conversion

package main

import (
	"fmt"
	"strconv"
)

func main() {
	isTrue, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(isTrue)

	num, err := strconv.ParseInt("1500000", 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)

	binary := strconv.FormatInt(5, 2)
	fmt.Println(binary)
}
