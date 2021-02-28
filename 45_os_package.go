package main

// run this file by passing arguments
// example: go run [file_name].go hello world

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args[0]) //
	fmt.Println(args[1]) // hello
	fmt.Println(args[2]) // world
}
