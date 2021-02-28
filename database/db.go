package database

import "fmt"

var connection string

// Package initialization
// The function that invoked when the package called
func init() {
	fmt.Println("Package initialization invoked")
	connection = "MongoDB"
}

func Connect() string {
	return connection
}
