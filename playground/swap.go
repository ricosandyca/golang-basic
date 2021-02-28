package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}

	// simple swap in golang
	// swap element 2 with element 4
	nums[1], nums[3] = nums[3], nums[1]
	fmt.Println(nums)
}
