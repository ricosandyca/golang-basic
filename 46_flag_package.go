package main

import (
	"flag"
	"fmt"
)

// run this file by passing arguments
// example: go run [file_name].go -name=Rico -age=22
// to see the description of each flag
// example: go run [file_name].go --help // -h // -help

func main() {
	name := flag.String("name", "default", "Put your name")
	age := flag.Int("age", 0, "Put your age")

	flag.Parse()

	fmt.Println(*name, *age)
}
