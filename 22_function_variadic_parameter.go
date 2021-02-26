package main

import (
	"fmt"
	"strings"
)

func sum(name string, nums ...int) (total int) {
	if strings.TrimSpace(name) != "" {
		fmt.Println("Hello", name)
	}

	total = 0
	for _, value := range nums {
		total += value
	}
	return total
}

func main() {
	nums1 := []int{1, 4, 1, 3, 2}
	fmt.Println("Sum of", nums1, sum("Rico", nums1...))

	fmt.Println("Sum of 1,3,6", sum(" ", 1, 3, 6))
}
