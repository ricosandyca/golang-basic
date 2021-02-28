package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Trim("  hello world  ", " "))
	fmt.Println(strings.Trim("---hello world--", "-"))
	fmt.Println(strings.Contains("hello world", "hello"))
	fmt.Println(strings.Split("Rico Sandyca Novenza", " "))
	fmt.Println(strings.ToLower("Hello World"))
	fmt.Println(strings.ToUpper("Hello World"))
	fmt.Println(strings.ReplaceAll("Rico Sandyca Novenza", "Rico", "Ric"))
}
