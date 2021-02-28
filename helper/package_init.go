package helper

import db "golang-basic/database"

// CallDBPackage is a function to try invoking package initializdation in database package
func CallDBPackage() {
	db.Connect()
}
