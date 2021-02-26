package main

import "fmt"

func fibo(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}

// fibonacci
// 1 1 2 3 5 8 13

func main() {
	fmt.Println(fibo(5))
	fmt.Println(fibo(7))
}
