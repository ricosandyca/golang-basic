package main

import (
	"fmt"
	db "golang-basic/database"
	helper "golang-basic/helper"
)

// if you only want to invoke the package initialization without using any data from its package
// you can set alias of the package to _
// import _ "golang-basic/database"

func main() {
	// it only invokes the database package initialization once even you call it twice
	fmt.Println(db.Connect())
	helper.CallDBPackage()
}
