package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("Hello")
	fmt.Println(err)
	fmt.Println(err.Error())

	err = nil
	if err == nil {
		fmt.Println("the nil is working")
	}
}
