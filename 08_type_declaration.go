package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpRico NoKTP = "123456789"
	var isMarried Married = false
	fmt.Println("No KTP", noKtpRico)
	fmt.Println("Marriage Status", isMarried)
}
