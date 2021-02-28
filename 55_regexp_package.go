package main

import (
	"fmt"
	"regexp"
)

func isValidEmail(email string) bool {
	regex := regexp.MustCompile(`(?i)^[a-z]+\@[a-z]+\.[a-z]+$`)
	return regex.MatchString(email)
}

func main() {
	// simpel email validation
	email := "ricosandyca@gmail.com"
	fmt.Println(isValidEmail(email))

	// replace string
	name := "Rico Sandyca Rico"
	regex := regexp.MustCompile(`(?i)rico`)
	myName := regex.ReplaceAllString(name, "me")
	fmt.Println(name)
	fmt.Println(myName)

	// split
	str := "Rico!Sandyca*Novenza me"
	regex = regexp.MustCompile(`(?i)[^a-z]`)
	// to split all matching pattern, use limit -1 in the second param
	// n is limit array result of the splitting
	strArray := regex.Split(str, -1)
	strArray2 := regex.Split(str, 2)
	fmt.Println(strArray, len(strArray))
	fmt.Println((strArray2), len(strArray2))
}
