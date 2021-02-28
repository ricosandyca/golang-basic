package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(4.4))  // 4
	fmt.Println(math.Round(4.55)) // 5
	fmt.Println(math.Floor(4.6))  // 4
	fmt.Println(math.Ceil(4.2))   // 5

	fmt.Println(math.Max(10, 20)) // 20
	fmt.Println(math.Min(5, 3))   // 3
}
